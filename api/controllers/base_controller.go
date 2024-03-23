package controllers

import (
	"errors"
	"fmt"

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
