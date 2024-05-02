package routers

import (
	"flick_tickets/api/controllers"
	"flick_tickets/api/middlewares"
	"flick_tickets/common/log"
	"flick_tickets/configs"
	"flick_tickets/core/events/sockets"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
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
	payment *controllers.ControllerPayMent,
	moive *controllers.ControllerMovie,
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
	//r.POST("/user/register", user.AddUser)

	adminGroup := r.Group("/")
	// adminGroup.Use(middlewares.Authenticate())
	// {
	adminGroup.POST("/user/upload/ticket", ticket.AddTicket)
	//	}
	r.GET("/customer/:token", auth.CheckToken)
	r.GET("/user/ticket", ticket.GetTicketById)
	r.GET("customers/ticket", ticket.GetAllTickets)
	r.GET("/user/load", file_lc.GetListFileById)
	r.POST("/user/verify/", aes.VerifyTickets)
	r.DELETE("/user/delete/ticket/:id", ticket.DeleteTicketsById)
	r.GET("/customer/ticket/action", ticket.GetAllTicketsAttachSale)
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
	r.PUT("/customer/update/order", order.UpsertOrderById)
	r.PUT("/customer/order/send", order.SubmitSendTicketByEmail) //webhook
	r.PUT("/customer/order/calcel", order.UpdateOrderWhenCancel) // webhook
	r.GET("/user/order/getlist", order.GetAllOrder)
	r.GET("/user/trigger", order.TriggerOrder)

	//customer
	r.POST("/customer/send/:email", customer.SendOtptoEmail)
	r.POST("/customer/verify/", customer.CheckOtpByEmail)
	r.POST("/customer/manager/register", customer.RegisterCustomersManager) //account admin
	r.POST("/customer/manager/login", customer.Login)
	r.POST("/customer/staff/register", customer.CreateAccountAdminManagerForStaff) // account staff
	r.GET("customer/staff/getall", customer.GetAllStaff)
	r.DELETE("/user/staff/delete", customer.DeleteStaffByName)
	r.GET("/supper_admin/create", customer.CreateAccountAdmin)
	r.POST("/customer/check", customer.CheckAccountAndSendOtp)
	r.PUT("/customer/reset", customer.VerifyOtpByEmailAndResetPassword)
	//
	r.POST("/customer/user/register", customer.RegisterAccountCustomer)
	r.PUT("/customer/user/update", customer.UpdateProfileCustomerByUserName)
	r.GET("/customer/user/profile", customer.FindCustomersByUsename)
	//show time
	r.POST("/use/add/time", showTime.AddShowTime)
	r.DELETE("/use/delete/time", showTime.DeleteShowTime)
	r.GET("/user/getlist/time", showTime.GetShowTimeByTicketId)
	r.GET("/customer/detail/showtime", showTime.DetailShowTime)
	//r.Use(middlewares.Authenticate())

	//address public
	r.GET("/public/customer/cities", addresPublic.GetAllCity)
	r.GET("/public/customer/districts", addresPublic.GetAllDistrictsByCityName)
	r.GET("/public/customer/communes", addresPublic.GetAllCommunesByDistrictName)

	//payment
	r.POST("/public/customer/payment/pay", payment.CreatePayment)
	r.GET("/public/customer/payment/request", payment.GetPaymentOrderByIdFromPayOs)
	r.GET("/public/customer/payment/return", payment.ReturnUrlAfterPayment)
	r.GET("/public/customer/payment/calcel", payment.ReturnUrlAftercanCelPayment)
	//moive
	r.POST("/user/movie/add", moive.AddMoiveType)
	r.GET("/user/movie/getlist", moive.GetAllMovieType)

	// Thêm công việc vào lịch để chạy mỗi 15 phút = 900s
	scheduler := cron.New()
	err := scheduler.AddFunc("*/15 * * * *", func() {
		resp, err := http.Get("http://localhost:8080/manager/user/trigger")
		if err != nil {
			// Xử lý lỗi khi gọi API
			log.Error(err, "error controller")
			return
		}
		defer resp.Body.Close()
		// Xử lý phản hồi nếu cần
	})
	if err != nil {
		log.Error(err, "error")
	}
	// Bắt đầu lịch sau khi thêm công việc vào
	scheduler.Start()

	return &ApiRouter{
		Engine: engine,
	}
}
