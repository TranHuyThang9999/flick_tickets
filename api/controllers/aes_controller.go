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

	var req entities.TokenReqCheckQrCode
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := aes.aes.CheckQrCode(&req)
	aes.baseController.Response(ctx, resp, err)
}
