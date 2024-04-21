package controllers

import (
	"flick_tickets/common/log"
	"flick_tickets/core/entities"
	"flick_tickets/core/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControllerPayMent struct {
	payment *usecase.UseCasePayment
	*baseController
}

func NewControllerParment(
	payment *usecase.UseCasePayment,
	baseController *baseController,
) *ControllerPayMent {
	return &ControllerPayMent{
		payment:        payment,
		baseController: baseController,
	}
}
func (c *ControllerPayMent) CreatePayment(ctx *gin.Context) {

	var req entities.CheckoutRequestType
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := c.payment.CreatePayment(ctx, req)
	if err != nil {
		log.Error(err, err.Error())
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, resp)
}
