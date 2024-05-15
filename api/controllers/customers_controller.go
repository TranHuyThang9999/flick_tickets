package controllers

import (
	"flick_tickets/core/entities"
	"flick_tickets/core/usecase"
	"net/http"
	"os"

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
func (c *ControllerCustomer) RegisterCustomersManager(ctx *gin.Context) { // acount for admin

	// file, err := ctx.FormFile("file")
	var req *entities.CustomersReqRegister

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := c.cus.RegisterManager(ctx, req)
	c.baseController.Response(ctx, resp, err)

}

func (c *ControllerCustomer) Login(ctx *gin.Context) {

	var req entities.CustomerReqLogin
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := c.cus.Login(ctx, &req)
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

func (c *ControllerCustomer) GetAllStaff(ctx *gin.Context) {
	resp, err := c.cus.GetAllStaff(ctx)
	c.baseController.Response(ctx, resp, err)
}
func (c *ControllerCustomer) DeleteStaffByName(ctx *gin.Context) {
	name := ctx.Query("name")
	resp, err := c.cus.DeleteStaffByName(ctx, name)
	c.baseController.Response(ctx, resp, err)
}
func (c *ControllerCustomer) CreateAccountAdmin(ctx *gin.Context) {
	path := "api/public/webhook/pages_superadmin/create_account_admin.html"
	htmlBytes, err := os.ReadFile(path)
	if err != nil {
		// Xử lý lỗi nếu có
		ctx.String(http.StatusInternalServerError, "Lỗi khi đọc tệp HTML")
		return
	}

	// Trả về trang HTML
	ctx.Data(http.StatusOK, "text/html; charset=utf-8", htmlBytes)
}

func (c *ControllerCustomer) CheckAccountAndSendOtp(ctx *gin.Context) {
	var req entities.CheckAccountAndSendOtpReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := c.cus.CheckAccountAndSendOtp(ctx, &req)
	c.baseController.Response(ctx, resp, err)

}
func (c *ControllerCustomer) VerifyOtpByEmailAndResetPassword(ctx *gin.Context) {
	var req entities.VerifyOtpByEmailReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := c.cus.VerifyOtpByEmailAndResetPassword(ctx, &req)
	c.baseController.Response(ctx, resp, err)
}
func (c *ControllerCustomer) RegisterAccountCustomer(ctx *gin.Context) {

	file, err := ctx.FormFile("file")
	var req *entities.RegisterAccountCustomerReq

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
	resp, err := c.cus.RegisterAccountCustomer(ctx, req)
	c.baseController.Response(ctx, resp, err)
}
func (c *ControllerCustomer) UpdateProfileCustomerByUserName(ctx *gin.Context) {

	file, err := ctx.FormFile("file")
	var req entities.UpdateProfileCustomerByUserNameReq

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
	resp, err := c.cus.UpdateProfileCustomerByUserName(ctx, &req)
	c.baseController.Response(ctx, resp, err)
}
func (c *ControllerCustomer) FindCustomersByUsename(ctx *gin.Context) {
	var req entities.GetCustomerByUseNameReq

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := c.cus.GetCustomerByUseName(ctx, &req)
	c.baseController.Response(ctx, resp, err)
}
func (c *ControllerCustomer) CreateTokenRespWhenLoginWithEmail(ctx *gin.Context) {
	email := ctx.Query("email")

	resp, err := c.cus.GenTokenByEmail(ctx, email)
	c.baseController.Response(ctx, resp, err)

}
func (c *ControllerCustomer) UpdatePassWord(ctx *gin.Context) {

	var req entities.UpdatePassWordReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := c.cus.UpdatePassWordByUsername(ctx, &req)
	c.baseController.Response(ctx, resp, err)
}
