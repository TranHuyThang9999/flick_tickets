package controllers

import (
	"flick_tickets/core/entities"
	"flick_tickets/core/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControllersUser struct {
	user *usecase.UseCaseUser
	*baseController
}

func NewControllersUser(
	user *usecase.UseCaseUser,
	baseController *baseController,
) *ControllersUser {
	return &ControllersUser{
		user:           user,
		baseController: baseController,
	}
}
func (u *ControllersUser) AddUser(ctx *gin.Context) {

	var req entities.Users

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
	resp, err := u.user.AddUserd(ctx, &req)

	u.baseController.Response(ctx, resp, err)

}
