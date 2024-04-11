package usecase

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"flick_tickets/common/enums"
	"flick_tickets/common/log"
	"flick_tickets/common/utils"
	"flick_tickets/configs"
	"flick_tickets/core/entities"
)

type UseCaseAes struct {
	config *configs.Configs
}

func NewUseCaseAes(cf *configs.Configs) (*UseCaseAes, error) {
	return &UseCaseAes{
		config: cf,
	}, nil
}
func (c *UseCaseAes) GeneratesTokenWithAesToQrCodeAndSendQrWithEmail(req *entities.TokenRequestSendQrCode) (*entities.TokenRespSendQrCode, error) {

	key := []byte(c.config.KeyAES128)
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
	log.Info(ciphertext)
	err = utils.GeneratesQrCodeAndSendQrWithEmail(req.FromEmail, req.Title, ciphertext)
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
		Created: utils.GenerateTimestamp(),
	}, nil
}
func (c *UseCaseAes) CheckQrCode(ctx context.Context, token string) (*entities.TokenResponseCheckQrCode, error) {

	if token == "" {
		return &entities.TokenResponseCheckQrCode{
			Result: entities.Result{
				Code:    enums.INVALID_REQUEST_CODE,
				Message: enums.INVALID_REQUEST_MESS,
			},
		}, nil
	}

	key := []byte(c.config.KeyAES128)

	data, err := c.DecryptAes(token, key)
	if err != nil {
		return &entities.TokenResponseCheckQrCode{
			Result: entities.Result{
				Code:    enums.AES_DECRYPT_AES_CODE,
				Message: enums.AES_DECRYPT_AES_MESS,
			},
		}, nil
	}
	return &entities.TokenResponseCheckQrCode{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Content: string(data),
		Created: utils.GenerateTimestamp(),
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
