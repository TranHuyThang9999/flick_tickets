package routers

import (
	"flick_tickets/api/controllers"
	"flick_tickets/api/middlewares"
	"flick_tickets/configs"
	"flick_tickets/core/events/sockets"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type ApiRouter struct {
	Engine *gin.Engine
}

func NewApiRouter(
	cf *configs.Configs,
	user *controllers.ControllersUser,
	auth *controllers.ControllerAuth,
	ticket *controllers.ControllerTicket,
	file_lc *controllers.ControllerFileLc,
	order *controllers.ControllerOrder,
	aes *controllers.ControllerAes,
	middlewares *middlewares.MiddleWare,
	customer *controllers.ControllerCustomer,
	managerClient *sockets.ManagerClient,
	showTime *controllers.ControllerShowTime,
	cinema *controllers.ControllerCinemas,
	addresPublic *controllers.ControllerAddress,

) *ApiRouter {
	engine := gin.New()
	gin.DisableConsoleColor()

	engine.Use(gin.Logger())
	engine.Use(cors.AllowAll())
	engine.Use(gin.Recovery())
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())
	r := engine.RouterGroup.Group("/manager")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	//admin
	r.POST("/user/register", user.AddUser)
	r.POST("/user/login", auth.LoginUser)

	adminGroup := r.Group("/")
	// adminGroup.Use(middlewares.Authenticate())
	// {
	adminGroup.POST("/user/upload/ticket", ticket.AddTicket)
	//	}

	r.GET("/user/ticket", ticket.GetTicketById)
	r.GET("customers/ticket", ticket.GetAllTickets)
	r.GET("/user/load", file_lc.GetListFileById)
	r.POST("/user/verify/", aes.VerifyTickets)
	// r.PUT("/user/update/size/room", ticket.UpdateSizeRoom)
	//	r.GET("/ws", managerClient.ServerWs) //auto pool
	//cinema
	r.POST("/user/add/cinema", cinema.AddCinema)
	r.GET("/user/get/cinema", cinema.GetAllCinema)
	r.DELETE("/user/delete/cinema", cinema.DeleteCinemaByName)
	r.GET("/user/getByName", cinema.GetAllCinemaByName)
	r.PUT("/user/update/width/height", cinema.UpdateColumnWidthHeightContainer)
	// user
	r.POST("/customer/order/ticket", order.OrdersTicket)
	r.GET("/customer/look/order/ticket", order.GetOrderById)
	//customer
	r.POST("/customer/send", customer.SendOtptoEmail)
	r.POST("/customer/verify/", customer.CheckOtpByEmail)
	r.POST("/customer/manager/register", customer.RegisterCustomersManager)
	r.POST("/customer/manager/login", customer.LoginCustomerManager)
	r.POST("/customer/staff/register", customer.CreateAccountAdminManagerForStaff)
	r.POST("/customer/staff/login", customer.LoginCustomerStaff)
	r.GET("customer/staff/getall", customer.GetAllStaff)
	r.DELETE("/user/staff/delete", customer.DeleteStaffByName)
	//show time
	r.POST("/use/add/time", showTime.AddShowTime)
	r.DELETE("/use/delete/time", showTime.DeleteShowTime)
	r.GET("/user/getlist/time", showTime.GetShowTimeByTicketId)
	//r.Use(middlewares.Authenticate())

	//address public
	r.GET("/public/customer/cities", addresPublic.GetAllCity)
	r.GET("/public/customer/districts", addresPublic.GetAllDistrictsByCityName)
	r.GET("/public/customer/communes", addresPublic.GetAllCommunesByDistrictName)

	return &ApiRouter{
		Engine: engine,
	}
}
