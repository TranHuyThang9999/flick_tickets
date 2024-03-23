package usecase

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"flick_tickets/common/enums"
	"flick_tickets/common/utils"
	"flick_tickets/configs"
	"flick_tickets/core/entities"
	"strconv"
)

type AesUseCase struct {
	config *configs.Configs
}

func NewAesUseCase(cf *configs.Configs) (*AesUseCase, error) {
	return &AesUseCase{}, nil
}
func (c *AesUseCase) GeneratesTokenWithAesToQrCodeAndSendQrWithEmail(req *entities.TokenRequestSendQrCode) (*entities.TokenRespSendQrCode, error) {

	key := []byte(c.config.KeyAES128)
	plaintext := []byte(strconv.FormatInt(req.Id, 10))

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

	return &entities.TokenRespSendQrCode{
		Token: ciphertext,
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Created: utils.GenerateTimestamp(),
	}, nil
}
func (c *AesUseCase) CheckQrCode(req *entities.TokenReqCheckQrCode) (*entities.TokenResponseCheckQrCode, error) {
	return &entities.TokenResponseCheckQrCode{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Created: utils.GenerateTimestamp(),
	}, nil
}

// Encrypt sử dụng AES để mã hóa dữ liệu với khóa đã cho.
func (e *AesUseCase) EncryptAes(data []byte, key []byte) (string, error) {
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
func (e *AesUseCase) DecryptAes(ciphertextBase64 string, key []byte) ([]byte, error) {
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
