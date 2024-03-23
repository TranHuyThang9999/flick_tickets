package controllers

import (
	"flick_tickets/core/entities"
	"flick_tickets/core/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControllerAuth struct {
	*baseController
	jwtUseCase *usecase.JwtUseCase
}

func NewControllerAuth(
	baseController *baseController,
	jwtUseCase *usecase.JwtUseCase,
) *ControllerAuth {
	return &ControllerAuth{
		baseController: baseController,
		jwtUseCase:     jwtUseCase,
	}
}
func (b *ControllerAuth) LoginUser(ctx *gin.Context) {

	var req entities.LoginReq

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Yêu cầu không hợp lệ"})
		return
	}
	if err := b.validateRequest(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := b.jwtUseCase.LoginUser(ctx, req.UserName, req.Password)
	if err != nil {
		ctx.JSON(200, err)
		return
	}
	ctx.JSON(200, resp)

}
