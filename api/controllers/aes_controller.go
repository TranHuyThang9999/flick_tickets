package controllers

import (
	"flick_tickets/core/entities"
	"flick_tickets/core/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControllerAes struct {
	aes *usecase.UseCaseAes
	*baseController
}

func NewControllerAes(
	aes *usecase.UseCaseAes,
	baseController *baseController,
) *ControllerAes {
	return &ControllerAes{
		aes:            aes,
		baseController: baseController,
	}
}
func (aes *ControllerAes) VerifyTickets(ctx *gin.Context) {

	// token := ctx.Param("token")
	var req entities.AesContentEncryptReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadGateway, err)
		return
	}
	resp, err := aes.aes.CheckQrCode(ctx, &req)
	aes.baseController.Response(ctx, resp, err)
}
