package controllers

import (
	"flick_tickets/core/entities"
	"flick_tickets/core/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControllerShowTime struct {
	st *usecase.UseCaseShowTime
	*baseController
}

func NewControllerShowTIme(
	st *usecase.UseCaseShowTime,
	baseController *baseController,
) *ControllerShowTime {
	return &ControllerShowTime{
		st:             st,
		baseController: baseController,
	}
}
func (c *ControllerShowTime) AddShowTime(ctx *gin.Context) {

	var req entities.ShowTimeAddReq

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := c.st.AddShowTime(ctx, &req)
	c.baseController.Response(ctx, resp, err)
}
func (c *ControllerShowTime) DeleteShowTime(ctx *gin.Context) {

	var req entities.ShowTimeDelete

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := c.st.DeleteShowTime(ctx, &req)
	c.baseController.Response(ctx, resp, err)
}
