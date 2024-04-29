package controllers

import (
	"flick_tickets/core/usecase"

	"github.com/gin-gonic/gin"
)

type ControllerAuth struct {
	*baseController
	jwtUseCase *usecase.UseCaseJwt
}

func NewControllerAuth(
	baseController *baseController,
	jwtUseCase *usecase.UseCaseJwt,
) *ControllerAuth {
	return &ControllerAuth{
		baseController: baseController,
		jwtUseCase:     jwtUseCase,
	}
}
func (c *ControllerAuth) CheckToken(ctx *gin.Context) {
	token := ctx.Param("token")
	resp, err := c.jwtUseCase.Decrypt(token)
	c.baseController.Response(ctx, resp, err)
}
