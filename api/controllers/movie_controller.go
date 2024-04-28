package controllers

import (
	"flick_tickets/core/entities"
	"flick_tickets/core/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControllerMovie struct {
	movie *usecase.UseCaseMovie
	*baseController
}

func NewControllerMovie(
	movie *usecase.UseCaseMovie,
	baseController *baseController,
) *ControllerMovie {
	return &ControllerMovie{
		movie: movie,
	}
}
func (c *ControllerMovie) AddMoiveType(ctx *gin.Context) {

	var req entities.MovieTypesAddReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := c.movie.AddMovieType(ctx, &req)
	c.baseController.Response(ctx, resp, err)
}
func (c *ControllerMovie) GetAllMovieType(ctx *gin.Context) {
	resp, err := c.movie.GetAllMovieType(ctx)
	c.baseController.Response(ctx, resp, err)

}
