package usecase

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"flick_tickets/common/enums"
	"flick_tickets/common/log"
	"flick_tickets/common/utils"
	"flick_tickets/configs"
	"flick_tickets/core/domain"
	"flick_tickets/core/entities"
	"flick_tickets/core/events/caching/cache"
	"flick_tickets/core/mapper"
)

type UseCaseAes struct {
	config *configs.Configs
	memory cache.RepositoryCache
	order  domain.RepositoryOrder
}

func NewUseCaseAes(
	cf *configs.Configs,
	memory cache.RepositoryCache,
	order domain.RepositoryOrder,
) (*UseCaseAes, error) {
	return &UseCaseAes{
		config: cf,
		memory: memory,
		order:  order,
	}, nil
}
func (c *UseCaseAes) GeneratesTokenWithAesToQrCodeAndSendQrWithEmail(req *entities.TokenRequestSendQrCode) (*entities.TokenRespSendQrCode, error) {

	key := []byte(configs.Get().KeyAES128)
	plaintext := []byte(req.Content)

	// Mã hóa dữ liệu
	ciphertext, err := c.EncryptAes(plaintext, key)
	if err != nil {
		return &entities.TokenRespSendQrCode{
			Result: entities.Result{
				Code:    enums.AES_ENCRYPT_AES_CODE,
				Message: enums.AES_ENCRYPT_AES_MESS,
			},
		}, nil
	}
	err = utils.GeneratesQrCodeAndSendQrWithEmail(req.FromEmail, &entities.OrderSendTicketToEmail{
		ID:         req.Order.ID,
		MoviceName: req.Order.MoviceName,
		Price:      req.Order.Price,
		Seats:      req.Order.Seats,
		CinemaName: req.Order.CinemaName,
		MovieTime:  req.Order.MovieTime,
	}, req.Title, ciphertext)
	if err != nil {
		return &entities.TokenRespSendQrCode{
			Result: entities.Result{
				Code:    enums.SEND_EMAIL_ERR_CODE,
				Message: enums.SEND_EMAIL_ERR_MESS,
			},
		}, nil
	}
	return &entities.TokenRespSendQrCode{
		Token: ciphertext,
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		CreatedAt: utils.GenerateTimestamp(),
	}, nil
}
func (c *UseCaseAes) CheckQrCode(ctx context.Context, req *entities.AesContentEncryptReq) (*entities.TokenResponseCheckQrCode, error) {

	log.Infof("req : ", req)

	if req.Token == "" {
		return &entities.TokenResponseCheckQrCode{
			Result: entities.Result{
				Code:    enums.INVALID_REQUEST_CODE,
				Message: enums.INVALID_REQUEST_MESS,
			},
		}, nil
	}

	//	key := []byte(c.config.KeyAES128)

	data, err := c.DecryptAes(req.Token, []byte(configs.Get().KeyAES128))
	if err != nil {
		return &entities.TokenResponseCheckQrCode{
			Result: entities.Result{
				Code:    enums.AES_DECRYPT_AES_CODE,
				Message: enums.AES_DECRYPT_AES_MESS,
			},
		}, nil
	}
	dataOrderId := string(data)
	statusExistsObject, err := c.memory.KeyExists(ctx, dataOrderId)
	if err != nil {
		return &entities.TokenResponseCheckQrCode{
			Result: entities.Result{
				Code:    enums.CACHE_ERR_CODE,
				Message: enums.CACHE_ERR_MESS,
			},
		}, nil
	}
	if !statusExistsObject {
		detailOrder, err := c.order.GetOrderById(ctx, int64(mapper.ConvertStringToInt(dataOrderId)))
		if err != nil {
			return &entities.TokenResponseCheckQrCode{
				Result: entities.Result{
					Code:    enums.DB_ERR_CODE,
					Message: enums.DB_ERR_MESS,
				},
			}, nil
		}
		orderResp := entities.OrderHistoryEntities{
			ID:             detailOrder.ID,
			MovieName:      detailOrder.MovieName,
			CinemaName:     detailOrder.CinemaName,
			Email:          detailOrder.Email,
			ReleaseDate:    detailOrder.ReleaseDate,
			Description:    detailOrder.Description,
			Status:         detailOrder.Status,
			Price:          detailOrder.Price,
			Seats:          detailOrder.Seats,
			MovieTime:      detailOrder.MovieTime,
			AddressDetails: detailOrder.AddressDetails,
			CreatedAt:      detailOrder.CreatedAt,
		}
		err = c.memory.SetObjectById(ctx, dataOrderId, detailOrder)
		if err != nil {
			return &entities.TokenResponseCheckQrCode{
				Result: entities.Result{
					Code:    enums.CACHE_ERR_CODE,
					Message: enums.CACHE_ERR_MESS,
				},
			}, nil
		}
		return &entities.TokenResponseCheckQrCode{
			Result: entities.Result{
				Code:    enums.SUCCESS_CODE,
				Message: enums.SUCCESS_MESS,
			},
			Order:     &orderResp,
			CreatedAt: utils.GenerateTimestamp(),
		}, nil
	}
	dataStringOrder, err := c.memory.GetObjectById(ctx, dataOrderId)
	if err != nil {
		return &entities.TokenResponseCheckQrCode{
			Result: entities.Result{
				Code:    enums.CACHE_ERR_CODE,
				Message: enums.CACHE_ERR_MESS,
			},
		}, nil
	}
	var orderResp *entities.OrderHistoryEntities
	err = json.Unmarshal([]byte(dataStringOrder), &orderResp)
	if err != nil {
		return &entities.TokenResponseCheckQrCode{
			Result: entities.Result{
				Code:    enums.ERROR_CONVERT_JSON_CODE,
				Message: enums.ERROR_CONVERT_JSON_MESS,
			},
		}, nil
	}
	return &entities.TokenResponseCheckQrCode{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Order: orderResp,
	}, nil
}

// Encrypt sử dụng AES để mã hóa dữ liệu với khóa đã cho.
func (e *UseCaseAes) EncryptAes(data []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Tạo một vector khởi tạo (IV) ngẫu nhiên
	iv := make([]byte, aes.BlockSize)

	// Tạo chế độ CBC với khối mã hóa và vector khởi tạo
	mode := cipher.NewCBCEncrypter(block, iv)

	// Thêm padding vào dữ liệu
	blockSize := aes.BlockSize
	data = utils.Pkcs7Pad(data, blockSize)

	// Mã hóa dữ liệu
	ciphertext := make([]byte, len(data))
	mode.CryptBlocks(ciphertext, data)

	// Chuyển đổi mã hóa thành chuỗi base64
	ciphertextBase64 := base64.StdEncoding.EncodeToString(ciphertext)

	return ciphertextBase64, nil
}

// Decrypt sử dụng AES để giải mã dữ liệu đã được mã hóa với khóa đã cho.
func (e *UseCaseAes) DecryptAes(ciphertextBase64 string, key []byte) ([]byte, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(ciphertextBase64)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Tạo một vector khởi tạo (IV) ngẫu nhiên
	iv := make([]byte, aes.BlockSize)

	// Tạo chế độ CBC với khối mã hóa và vector khởi tạo
	mode := cipher.NewCBCDecrypter(block, iv)

	// Giải mã dữ liệu
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	// Xóa padding từ dữ liệu giải mã
	plaintext, err = utils.Pkcs7Unpad(plaintext)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
