package controllers

import (
	"flick_tickets/core/entities"
	"flick_tickets/core/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControllerTicket struct {
	ticket *usecase.UseCaseTicker
	*baseController
}

func NewControllerTicket(
	ticket *usecase.UseCaseTicker,
	baseController *baseController,
) *ControllerTicket {
	return &ControllerTicket{
		ticket:         ticket,
		baseController: baseController,
	}
}
func (c *ControllerTicket) AddTicket(ctx *gin.Context) {

	var req entities.TicketReqUpload

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	files, err := c.baseController.GetUploadedFiles(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.File = files
	resp, err := c.ticket.AddTicket(ctx, &req)
	c.baseController.Response(ctx, resp, err)

}
func (c *ControllerTicket) GetTicketById(ctx *gin.Context) {

	id := ctx.Query("id")

	resp, err := c.ticket.GetTicketById(ctx, id)
	c.baseController.Response(ctx, resp, err)

}
