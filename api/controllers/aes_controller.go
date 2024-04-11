package controllers

import (
	"flick_tickets/core/usecase"

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

	token := ctx.Query("token")

	resp, err := aes.aes.CheckQrCode(ctx, token)
	aes.baseController.Response(ctx, resp, err)
}
