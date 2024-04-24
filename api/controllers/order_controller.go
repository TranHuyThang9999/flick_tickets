package controllers

import (
	"flick_tickets/core/entities"
	"flick_tickets/core/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControllerOrder struct {
	order *usecase.UseCaseOrder
	*baseController
}

func NewControllerOrder(
	order *usecase.UseCaseOrder,
	baseController *baseController,
) *ControllerOrder {
	return &ControllerOrder{
		order:          order,
		baseController: baseController,
	}
}

func (order *ControllerOrder) OrdersTicket(ctx *gin.Context) {
	var req entities.OrdersReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := order.order.RegisterTicket(ctx, &req)
	if err != nil {
		// Xử lý lỗi
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	// Lưu giá trị từ phản hồi vào cookie
	cookie := http.Cookie{
		Name:     "cookieName",
		Value:    string(resp.OrderId), // Thay resp.SomeValue bằng giá trị bạn muốn lưu vào cookie
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
	}
	ctx.SetCookie(cookie.Name, cookie.Value, 0, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)

	order.baseController.Response(ctx, resp, err)
}

func (order *ControllerOrder) GetOrderById(ctx *gin.Context) {

	idTicket := ctx.Query("id")

	resp, err := order.order.GetOrderById(ctx, idTicket)

	order.baseController.Response(ctx, resp, err)

}

func (order *ControllerOrder) UpsertOrderById(ctx *gin.Context) {

	var req entities.OrderReqUpSert

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := order.order.UpsertOrderById(ctx, &req)

	order.baseController.Response(ctx, resp, err)
}

func (order *ControllerOrder) SubmitSendTicketByEmail(ctx *gin.Context) {

	var req entities.OrderSendTicketAfterPaymentReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := order.order.SendticketAfterPayment(ctx, &req)

	order.baseController.Response(ctx, resp, err)
}
