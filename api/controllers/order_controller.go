package controllers

import (
	"flick_tickets/core/domain"
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

	ticketId := ctx.Query("id")

	resp, err := order.order.GetOrderById(ctx, ticketId)

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
func (order *ControllerOrder) UpdateOrderWhenCancel(ctx *gin.Context) {

	var req entities.OrderCancelBtyIdreq

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := order.order.UpdateOrderWhenCancel(ctx, &req)

	order.baseController.Response(ctx, resp, err)
}
func (order *ControllerOrder) GetAllOrder(ctx *gin.Context) {

	var req domain.OrdersReqByForm
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := order.order.GetAllOrder(ctx, &req)

	order.baseController.Response(ctx, resp, err)
}
func (order *ControllerOrder) TriggerOrder(ctx *gin.Context) {

	err := order.order.TriggerOrder(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error server": err})
		return
	}
	ctx.JSON(200, "trigger sucess")
}
func (order *ControllerOrder) GetOrderHistory(ctx *gin.Context) {
	var req entities.OrderHistoryReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := order.order.OrderHistory(ctx, &req)

	order.baseController.Response(ctx, resp, err)
}
func (order *ControllerOrder) OrderRevenueByMovieName(ctx *gin.Context) {
	var req entities.OrderRevenueReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := order.order.OrderRevenueByMovieName(ctx, &req)

	order.baseController.Response(ctx, resp, err)
}
func (order *ControllerOrder) GetAllMovieNameFromOrder(ctx *gin.Context) {

	resp, err := order.order.GetAllMovieNameFromOrder(ctx)
	order.baseController.Response(ctx, resp, err)

}
func (order *ControllerOrder) GetAllCinemaByMovieName(ctx *gin.Context) {
	movie_name := ctx.Query("cinema_name")
	resp, err := order.order.GetAllCinemaByMovieName(ctx, movie_name)
	order.baseController.Response(ctx, resp, err)

}
func (order *ControllerOrder) GetAllOrderStatistical(ctx *gin.Context) {
	var req entities.OrderStatisticalReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := order.order.StatisticalOrder(ctx, &req)
	order.baseController.Response(ctx, resp, err)
}
