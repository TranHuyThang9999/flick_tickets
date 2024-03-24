package controllers

import (
	"flick_tickets/core/usecase"

	"github.com/gin-gonic/gin"
)

type ControllerFileLc struct {
	file *usecase.UseCaseFileStore
	*baseController
}

func NewControllerFileLc(
	file *usecase.UseCaseFileStore,
	baseController *baseController,
) *ControllerFileLc {
	return &ControllerFileLc{
		file:           file,
		baseController: baseController,
	}
}
func (lc *ControllerFileLc) GetListFileById(ctx *gin.Context) {

	id := ctx.Query("id")

	resp, err := lc.file.GetListFileByObjectId(ctx, id)
	lc.baseController.Response(ctx, resp, err)

}
