package controllers

import (
	"errors"
	"fmt"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type baseController struct {
	validate *validator.Validate
}

// NewBaseController tạo một baseController mới
func NewBaseController(validate *validator.Validate) *baseController {
	return &baseController{
		validate: validate,
	}
}

// validateRequest kiểm tra tính hợp lệ của request
func (b *baseController) validateRequest(obj interface{}) error {
	validate := validator.New()
	err := validate.Struct(obj)
	if err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			var errorMsg string
			for _, e := range validationErrs {
				fieldName := e.StructField()
				// errTag := e.Tag()
				errorMsg = fmt.Sprintf("Invalid request : %s", fieldName)
				break
			}
			return errors.New(errorMsg)
		}
		return err
	}
	return nil
}
func (b *baseController) Response(ctx *gin.Context, resp interface{}, err error) {
	if err != nil {
		ctx.JSON(505, err)
		return
	}
	ctx.JSON(200, resp)
}
func (b *baseController) GetUploadedFiles(c *gin.Context) ([]*multipart.FileHeader, error) {
	form, err := c.MultipartForm()
	if err != nil {
		return nil, err
	}

	files, ok := form.File["file"]
	if !ok || len(files) == 0 {
		return nil, nil
	}

	var uploadedFiles []*multipart.FileHeader
	for _, file := range files {
		if file.Size == 0 {
			return nil, fmt.Errorf("Uploaded file is empty")
		}
		uploadedFiles = append(uploadedFiles, file)
	}

	return uploadedFiles, nil
}
