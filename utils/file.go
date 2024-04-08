package utils

import (
	"file_store/request"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func SaveFile(file *multipart.FileHeader, filePath string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return nil
}

func SaveImage(file *multipart.FileHeader, baseDir string) (string, error) {
	fileExt := filepath.Ext(file.Filename)
	fileName := uuid.New().String() + fileExt
	filePath := filepath.Join(baseDir, fileName)

	if err := SaveFile(file, filePath); err != nil {
		return "", err
	}

	return fileName, nil
}

func SaveImage2(file *multipart.FileHeader) (*request.InforFileResp, error) {

	create_at := GetCurrentTimestamp()

	if file == nil {
		return &request.InforFileResp{
			Result: request.Result{
				Code:    2,
				Message: "Không có tệp được tải lên",
			},
		}, nil
	}

	fileExt := strings.ToLower(filepath.Ext(file.Filename))
	//if fileExt == ".exe" {
	//	return &request.InforFileResp{
	//		Result: request.Result{
	//			Code:    3,
	//			Message: "Tệp không hợp lệ",
	//		},
	//		CreatedAt: create_at,
	//		UpdatedAt: create_at,
	//	}, nil
	//}
	baseDir := "shader"
	//newFileName := GenUuid()
	newFileName := strconv.FormatInt(GenerateUniqueKey(), 10) //fileExt := filepath.Ext(file.Filename)
	fileName := newFileName + fileExt
	filePath := filepath.Join(baseDir, fileName)

	if err := SaveFile(file, filePath); err != nil {
		return nil, err
	}

	return &request.InforFileResp{
		Result: request.Result{
			Code:    0,
			Message: "upload successfully",
		},
		Id:        newFileName,
		Url:       "http://localhost:1234/manager/shader/huythang/" + fileName,
		CreatedAt: create_at,
		UpdatedAt: create_at,
	}, nil
}
