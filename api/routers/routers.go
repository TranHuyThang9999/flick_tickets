package routers

import (
	"flick_tickets/api/controllers"
	"flick_tickets/api/middlewares"
	"flick_tickets/configs"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type ApiRouter struct {
	Engine *gin.Engine
}

func NewApiRouter(
	user *controllers.ControllersUser,
	auth *controllers.ControllerAuth,
	ticket *controllers.ControllerTicket,
	file_lc *controllers.ControllerFileLc,
	order *controllers.ControllerOrder,
	aes *controllers.ControllerAes,
	middlewares *middlewares.MiddleWare,
	customer *controllers.ControllerCustomer,
	cf *configs.Configs,
) *ApiRouter {
	engine := gin.New()
	gin.DisableConsoleColor()

	engine.Use(gin.Logger())
	engine.Use(cors.AllowAll())
	//middlewares.recovy
	engine.Use(gin.Recovery())

	r := engine.RouterGroup.Group("/manager")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	//admin
	r.POST("/user/register", user.AddUser)
	r.POST("/user/login", auth.LoginUser)
	r.POST("/user/upload/ticket", ticket.AddTicket)
	r.GET("/user/load", file_lc.GetListFileById)
	r.POST("/user/verify/", aes.VerifyTickets)
	// user
	r.POST("/customer/register/ticket", order.OrdersTicket)
	r.GET("/customer/order/ticket", order.GetOrderById)
	r.POST("/customer/send", customer.SendOtptoEmail)
	r.POST("/customer/verify/", customer.CheckOtpByEmail)
	//r.Use(middlewares.Authenticate())
	return &ApiRouter{
		Engine: engine,
	}
}
