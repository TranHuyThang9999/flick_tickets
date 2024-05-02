package controllers

import (
	"flick_tickets/core/domain"
	"flick_tickets/core/entities"
	"flick_tickets/core/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControllerCarts struct {
	cart *usecase.UseCaseCart
	*baseController
}

func NewControllerCart(cart *usecase.UseCaseCart, base *baseController) *ControllerCarts {
	return &ControllerCarts{
		cart:           cart,
		baseController: base,
	}
}

func (u *ControllerCarts) AddCart(ctx *gin.Context) {

	var req entities.CartsAddReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := u.cart.AddCart(ctx, &req)
	u.baseController.Response(ctx, resp, err)
}

func (u *ControllerCarts) FindByFormcart(ctx *gin.Context) {

	var req domain.CartFindByFormReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := u.cart.FindCartByForm(ctx, &req)
	u.baseController.Response(ctx, resp, err)
}

func (u *ControllerCarts) DeleteCartById(ctx *gin.Context) {

	id := ctx.Param("id")

	resp, err := u.cart.DeleteCartById(ctx, id)
	u.baseController.Response(ctx, resp, err)

}
func (u *ControllerCarts) UpdateCartById(ctx *gin.Context) {

	var req entities.CartsUpdateReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := u.cart.UpdateCartById(ctx, &req)
	u.baseController.Response(ctx, resp, err)
}
