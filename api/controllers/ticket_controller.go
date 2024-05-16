package controllers

import (
	"flick_tickets/core/domain"
	"flick_tickets/core/entities"
	"flick_tickets/core/events/sockets"
	"flick_tickets/core/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControllerTicket struct {
	ticket *usecase.UseCaseTicker
	*baseController
	socket *sockets.ManagerClient
}

func NewControllerTicket(
	ticket *usecase.UseCaseTicker,
	baseController *baseController,
	socket *sockets.ManagerClient,
) *ControllerTicket {
	return &ControllerTicket{
		ticket:         ticket,
		baseController: baseController,
		socket:         socket,
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

func (c *ControllerTicket) GetAllTickets(ctx *gin.Context) {
	var req domain.TicketreqFindByForm

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := c.ticket.GetAllTickets(ctx, &req)

	c.baseController.Response(ctx, resp, err)
}

func (c *ControllerTicket) DeleteTicketsById(ctx *gin.Context) {
	id := ctx.Param("id")
	resp, err := c.ticket.DeleteTicketsById(ctx, id)
	c.baseController.Response(ctx, resp, err)
}
func (c *ControllerTicket) GetAllTicketsAttachSale(ctx *gin.Context) {

	status := ctx.Query("status")
	resp, err := c.ticket.GetAllTicketsAttachSale(ctx, status)
	c.baseController.Response(ctx, resp, err)
}
func (c *ControllerTicket) UpdateTicketById(ctx *gin.Context) {

	var req entities.TicketReqUpdateById
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := c.ticket.UpdateTicketById(ctx, &req)

	c.baseController.Response(ctx, resp, err)
}
func (c *ControllerTicket) GetAllTicketsByFilmName(ctx *gin.Context) {

	var req entities.TicketFindByMovieNameReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := c.ticket.GetAllTicketsByFilmName(ctx, &req)

	c.baseController.Response(ctx, resp, err)
}
