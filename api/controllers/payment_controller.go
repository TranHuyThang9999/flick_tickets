package controllers

import (
	"flick_tickets/common/log"
	"flick_tickets/common/utils"
	"flick_tickets/core/entities"
	"flick_tickets/core/usecase"
	"fmt"
	"net/http"
	"os"

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
	// Tạo một id duy nhất cho cookie
	id := utils.GenerateUniqueKey()

	// Lưu id vào cookie và đặt nó để tồn tại ở domain cụ thể
	ctx.SetCookie("order_id", fmt.Sprintln(id), 3600, "/", "localhost:8080", false, true)

	var req entities.CheckoutRequestController
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := c.payment.CreatePayment(ctx, entities.CheckoutRequestType{
		OrderCode:    id,
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
		ShowTimeId:   req.ShowTimeId,
		Seats:        req.Seats,
		ExpiredAt:    utils.GenerateTimestampExpiredAt(15), // thoi gian ton tai QR code = 15 phut
	})

	if err != nil {
		log.Error(err, err.Error())
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, resp)
}

func (c *ControllerPayMent) GetPaymentOrderByIdFromPayOs(ctx *gin.Context) {

	id := ctx.Query("id")
	resp, err := c.payment.GetOrderByIdFromPayOs(ctx, id)
	c.baseController.Response(ctx, resp, err)

}

func (c *ControllerPayMent) ReturnUrlAfterPayment(ctx *gin.Context) {
	path := "api/public/webhook/payment/create_payment.html"
	htmlBytes, err := os.ReadFile(path)
	if err != nil {
		// Xử lý lỗi nếu có
		ctx.String(http.StatusInternalServerError, "Lỗi khi đọc tệp HTML")
		return
	}

	// Trả về trang HTML
	ctx.Data(http.StatusOK, "text/html; charset=utf-8", htmlBytes)

}
func (c *ControllerPayMent) ReturnUrlAftercanCelPayment(ctx *gin.Context) {
	path := "api/public/webhook/payment/cancel_payment.html"
	htmlBytes, err := os.ReadFile(path)
	if err != nil {
		// Xử lý lỗi nếu có
		ctx.String(http.StatusInternalServerError, "Lỗi khi đọc tệp HTML")
		return
	}

	// Trả về trang HTML
	ctx.Data(http.StatusOK, "text/html; charset=utf-8", htmlBytes)
}
