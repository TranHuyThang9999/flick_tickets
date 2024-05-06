package controllers

import (
	"flick_tickets/core/entities"
	"flick_tickets/core/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControllerCinemas struct {
	*baseController
	cus *usecase.UseCaseCinemas
}

func NewControllerCinamas(
	base *baseController,
	cus *usecase.UseCaseCinemas,
) *ControllerCinemas {
	return &ControllerCinemas{
		baseController: base,
		cus:            cus,
	}
}
func (c *ControllerCinemas) AddCinema(ctx *gin.Context) {

	var req entities.CinemasReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := c.cus.AddCinemas(ctx, &req)
	c.baseController.Response(ctx, resp, err)
}

func (c *ControllerCinemas) GetAllCinema(ctx *gin.Context) {
	resp, err := c.cus.GetAllCinema(ctx)
	c.baseController.Response(ctx, resp, err)

}
func (c *ControllerCinemas) DeleteCinemaByName(ctx *gin.Context) {
	name := ctx.Query("name")
	resp, err := c.cus.DeleteCinemaByName(ctx, name)
	c.baseController.Response(ctx, resp, err)

}
func (c *ControllerCinemas) GetAllCinemaByName(ctx *gin.Context) {
	name := ctx.Query("name")
	resp, err := c.cus.GetAllCinemaByName(ctx, name)
	c.baseController.Response(ctx, resp, err)

}
func (c *ControllerCinemas) UpdateColumnWidthHeightContainer(ctx *gin.Context) {

	var req entities.CinemaReqUpdateSizeRoom
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := c.cus.UpdateColumnWidthHeightContainer(ctx, &req)
	c.baseController.Response(ctx, resp, err)

}
