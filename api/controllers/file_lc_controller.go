package controllers

import (
	"flick_tickets/core/usecase"

	"github.com/gin-gonic/gin"
)

type ControllerFileLc struct {
	file *usecase.UseCaseFileStore
}

func NewControllerFileLc(file *usecase.UseCaseFileStore) *ControllerFileLc {
	return &ControllerFileLc{
		file: file,
	}
}
func (lc *ControllerFileLc) GetListFileById(ctx *gin.Context) {

	id := ctx.Query("id")

	resp, err := lc.file.GetListFileByObjectId(ctx, id)
	if err != nil {
		ctx.JSON(200, err)
		return
	}
	ctx.JSON(200, resp)
}
