package controllers

import (
	"flick_tickets/core/entities"
	"flick_tickets/core/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControllerCustomer struct {
	*baseController
	cus *usecase.UseCaseCustomer
}

func NewControllerCustomer(
	baseController *baseController,
	cus *usecase.UseCaseCustomer,
) *ControllerCustomer {
	return &ControllerCustomer{
		baseController: baseController,
		cus:            cus,
	}
}
func (c *ControllerCustomer) SendOtptoEmail(ctx *gin.Context) {

	email := ctx.Query("email")
	resp, err := c.cus.SendOtpToEmail(ctx, email)
	c.baseController.Response(ctx, resp, err)

}
func (c *ControllerCustomer) CheckOtpByEmail(ctx *gin.Context) {

	var req entities.CustomersReqVerifyOtp
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := c.cus.CheckOtp(ctx, &req)
	c.baseController.Response(ctx, resp, err)

}
