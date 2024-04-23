package controllers

import (
	"flick_tickets/core/entities"
	"flick_tickets/core/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControllerCustomer struct {
	*baseController
	cus *usecase.UseCaseCustomer
}

func NewControllerCustomer(
	baseController *baseController,
	cus *usecase.UseCaseCustomer,
) *ControllerCustomer {
	return &ControllerCustomer{
		baseController: baseController,
		cus:            cus,
	}
}
func (c *ControllerCustomer) SendOtptoEmail(ctx *gin.Context) {

	email := ctx.Param("email")
	resp, err := c.cus.SendOtpToEmail(ctx, email)
	c.baseController.Response(ctx, resp, err)

}
func (c *ControllerCustomer) CheckOtpByEmail(ctx *gin.Context) {

	var req entities.CustomersReqVerifyOtp
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := c.cus.CheckOtp(ctx, &req)
	c.baseController.Response(ctx, resp, err)

}
func (c *ControllerCustomer) RegisterCustomersManager(ctx *gin.Context) {

	file, err := ctx.FormFile("file")
	var req *entities.CustomersReqRegister

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err != nil && err != http.ErrMissingFile && err != http.ErrNotMultipart {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Không thể tải ảnh lên.",
		})
		return
	}

	req.File = file
	resp, err := c.cus.RegisterManager(ctx, req)
	c.baseController.Response(ctx, resp, err)

}

func (c *ControllerCustomer) LoginCustomerManager(ctx *gin.Context) {

	var req *entities.CustomerReqLogin
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := c.cus.LoginCustomerManager(ctx, req)
	c.baseController.Response(ctx, resp, err)

}
func (c *ControllerCustomer) CreateAccountAdminManagerForStaff(ctx *gin.Context) {

	file, err := ctx.FormFile("file")
	var req *entities.CustomersReqRegisterAdminForStaff

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err != nil && err != http.ErrMissingFile && err != http.ErrNotMultipart {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Không thể tải ảnh lên.",
		})
		return
	}

	req.File = file
	resp, err := c.cus.CreateAccountAdminManagerForStaff(ctx, req)
	c.baseController.Response(ctx, resp, err)
}

func (c *ControllerCustomer) LoginCustomerStaff(ctx *gin.Context) {

	var req *entities.CustomerReqLogin
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := c.cus.LoginCustomerForStaff(ctx, req)
	c.baseController.Response(ctx, resp, err)

}
func (c *ControllerCustomer) GetAllStaff(ctx *gin.Context) {
	resp, err := c.cus.GetAllStaff(ctx)
	c.baseController.Response(ctx, resp, err)
}
func (c *ControllerCustomer) DeleteStaffByName(ctx *gin.Context) {
	name := ctx.Query("name")
	resp, err := c.cus.DeleteStaffByName(ctx, name)
	c.baseController.Response(ctx, resp, err)
}
