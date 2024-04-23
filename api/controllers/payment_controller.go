package controllers

import (
	"flick_tickets/common/log"
	"flick_tickets/common/utils"
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

	var req entities.CheckoutRequestController
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := c.payment.CreatePayment(ctx, entities.CheckoutRequestType{
		OrderCode:    utils.GenerateUniqueKey(),
		Amount:       req.Amount,
		Description:  req.Description,
		CancelUrl:    req.CancelUrl,
		ReturnUrl:    req.ReturnUrl,
		Signature:    req.Signature,
		Items:        req.Items,
		BuyerName:    req.BuyerName,
		BuyerEmail:   req.BuyerEmail,
		BuyerPhone:   req.BuyerPhone,
		BuyerAddress: req.BuyerAddress,
		ExpiredAt:    utils.GenerateTimestampExpiredAt(15),
	})

	if err != nil {
		log.Error(err, err.Error())
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, resp)
}
