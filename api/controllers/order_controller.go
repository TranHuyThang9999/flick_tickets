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
