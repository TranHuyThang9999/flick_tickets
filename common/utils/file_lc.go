package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"flick_tickets/common/enums"
	"flick_tickets/common/log"
	"flick_tickets/configs"
	"flick_tickets/core/entities"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

func SetByCurlImage(ctx context.Context, file *multipart.FileHeader) (*entities.UploadResponse, error) {

	filePath := file.Filename
	fileExt := strings.ToLower(filepath.Ext(file.Filename))
	url := strings.TrimSpace(strings.Trim(configs.Get().FileLc, " "))

	acceptedExts := []string{".jpg", ".jpeg", ".gif", ".png", ".svg", ".MP4", ".flv", ".webm", ".mp4"}
	accepted := false
	for _, ext := range acceptedExts {
		if fileExt == ext {
			accepted = true
			break
		}
	}

	if !accepted {
		return &entities.UploadResponse{
			Result: entities.Result{
				Code:    7,
				Message: "Định dạng file không hợp lệ",
			},
		}, nil
	}

	// Tạo một multipart form data
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// Thêm file vào form data
	fileWriter, err := bodyWriter.CreateFormFile("file", filePath)
	if err != nil {
		return &entities.UploadResponse{
			Result: entities.Result{
				Code:    4,
				Message: "Lỗi server",
			},
		}, nil
	}
	fh, err := file.Open()
	if err != nil {
		return &entities.UploadResponse{
			Result: entities.Result{
				Code:    4,
				Message: "Lỗi server",
			},
		}, nil
	}
	defer fh.Close()
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return &entities.UploadResponse{
			Result: entities.Result{
				Code:    4,
				Message: "Lỗi server",
			},
		}, nil
	}

	// Kết thúc form data
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	// Gửi yêu cầu POST
	resp, err := http.Post(url, contentType, bodyBuf)
	log.Errorf(err, "error")
	if err != nil {
		return &entities.UploadResponse{
			Result: entities.Result{
				Code:    404,
				Message: fmt.Sprintf("Không thể tìm thấy tài nguyên: %s", err),
			},
		}, nil
	}
	defer resp.Body.Close()

	// Đọc phản hồi
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &entities.UploadResponse{
			Result: entities.Result{
				Code:    4,
				Message: "Lỗi server",
			},
		}, nil
	}

	// Decode phản hồi thành struct
	var uploadResp entities.UploadResponse
	err = json.Unmarshal(respBody, &uploadResp)
	if err != nil {
		return &entities.UploadResponse{
			Result: entities.Result{
				Code:    4,
				Message: "Lỗi server",
			},
		}, nil
	}
	switch uploadResp.Result.Code {
	case 1:
		return &entities.UploadResponse{
			Result: entities.Result{
				Code:    1,
				Message: "tệp rỗng 1 ",
			},
		}, nil
	case 2:
		return &entities.UploadResponse{
			Result: entities.Result{
				Code:    2,
				Message: "Không có tệp được tải lên 1",
			},
		}, nil
	case 3:
		return &entities.UploadResponse{
			Result: entities.Result{
				Code:    3,
				Message: "Tệp không hợp lệ 1",
			},
		}, nil
	default:
		return &uploadResp, nil
	}
}

func SetListFile(ctx context.Context, files []*multipart.FileHeader) ([]*entities.UploadResponse, error) {

	var responses = make([]*entities.UploadResponse, 0)
	for _, file := range files {
		response, err := SetByCurlImage(ctx, file)
		if err != nil || response.Result.Code != 0 {
			return nil, err
		}
		responses = append(responses, response)
	}
	return responses, nil
}

// SetByCurlImageQr gửi dữ liệu hình ảnh đến dịch vụ khác
func setByCurlImageQrToService(imageData []byte) *entities.UploadResponse {
	// URL của dịch vụ lưu trữ hình ảnh
	url := strings.TrimSpace(strings.Trim(configs.Get().FileLc, " "))
	fileConvert := strconv.FormatInt(GenerateUniqueKey(), 10)

	// Tạo buffer để tạo multipart form data
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileName := fmt.Sprintf("%s_%s.png", "qr_code", fileConvert)

	// Thêm hình ảnh vào form data
	fileWriter, err := bodyWriter.CreateFormFile("file", fileName)
	if err != nil {
		return &entities.UploadResponse{
			Result: entities.Result{
				Code:    4,
				Message: "Server error",
			},
		}
	}
	_, err = fileWriter.Write(imageData)
	if err != nil {
		return &entities.UploadResponse{
			Result: entities.Result{
				Code:    4,
				Message: "Server error",
			},
		}
	}

	// Kết thúc form data
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	// Gửi yêu cầu POST đến dịch vụ lưu trữ hình ảnh
	resp, err := http.Post(url, contentType, bodyBuf)
	if err != nil {
		return &entities.UploadResponse{
			Result: entities.Result{
				Code:    404,
				Message: fmt.Sprintf("Resource not found: %s", err),
			},
		}
	}
	defer resp.Body.Close()

	// Đọc phản hồi
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &entities.UploadResponse{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}
	}

	// Giải mã phản hồi thành cấu trúc
	var uploadResp entities.UploadResponse
	err = json.Unmarshal(respBody, &uploadResp)
	if err != nil {
		return &entities.UploadResponse{
			Result: entities.Result{
				Code:    4,
				Message: "Server error",
			},
		}
	}

	return &uploadResp
}
