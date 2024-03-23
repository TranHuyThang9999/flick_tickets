package controllers

import (
	"flick_tickets/core/entities"
	"flick_tickets/core/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControllerTicket struct {
	ticket *usecase.UseCaseTicker
}

func NewControllerTicket(ticket *usecase.UseCaseTicker) *ControllerTicket {
	return &ControllerTicket{
		ticket: ticket,
	}
}
func (c *ControllerTicket) AddTicket(ctx *gin.Context) {

	var req entities.TicketReqUpload

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	file, err := ctx.FormFile("file")

	if err != nil && err != http.ErrMissingFile && err != http.ErrNotMultipart {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Không thể tải ảnh lên.",
		})
		return
	}
	req.File = file
	resp, err := c.ticket.AddTicket(ctx, &req)
	if err != nil {
		ctx.JSON(200, err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
