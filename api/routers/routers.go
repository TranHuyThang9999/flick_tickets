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
	payment *controllers.ControllerPayMent,
	moive *controllers.ControllerMovie,
	cart *controllers.ControllerCarts,
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
	//process file local
	r.GET("/customer/:token", auth.CheckToken) //file
	r.GET("/user/ticket", ticket.GetTicketById)
	r.GET("customers/ticket", ticket.GetAllTickets)
	r.GET("/user/load", file_lc.GetListFileById)
	r.PUT("/user/upload", file_lc.UpSertFileDescriptByTicketId)
	r.DELETE("/user/delete/file/:id", file_lc.DeleteFileById)
	r.GET("/user/verify/aes", aes.VerifyTickets)
	r.DELETE("/user/delete/ticket/:id", ticket.DeleteTicketsById)
	r.GET("/customer/ticket/action", ticket.GetAllTicketsAttachSale) // ko dung
	r.PUT("/use/ticket/updates", ticket.UpdateTicketById)
	r.GET("/customer/ticket/movie", ticket.GetAllTicketsByFilmName)
	// r.PUT("/user/update/size/room", ticket.UpdateSizeRoom)
	//	r.GET("/ws", managerClient.ServerWs) //auto pool
	//cinema
	r.POST("/user/add/cinema", cinema.AddCinema)
	r.GET("/user/get/cinema", cinema.GetAllCinema)
	r.DELETE("/user/delete/cinema", cinema.DeleteCinemaByName)
	r.GET("/user/getByName", cinema.GetAllCinemaByName)
	r.PUT("/user/update/width/height", cinema.UpdateColumnWidthHeightContainer)
	// user //order
	r.POST("/customer/order/ticket", order.OrdersTicket)
	r.GET("/customer/look/order/ticket", order.GetOrderById) //ko dung
	r.PUT("/customer/update/order", order.UpsertOrderById)
	r.PUT("/customer/order/send", order.SubmitSendTicketByEmail) //webhook
	r.PUT("/customer/order/calcel", order.UpdateOrderWhenCancel) // webhook
	r.GET("/user/order/getlist", order.GetAllOrder)              // thong ke co admin
	r.GET("/user/trigger", order.TriggerOrder)
	r.GET("/user/order/history", order.GetOrderHistory)
	r.GET("/user/order/revenue", order.OrderRevenueByMovieName)
	r.GET("/user/history", order.GetAllMovieNameFromOrder)
	r.GET("/user/history/movie/name", order.GetAllCinemaByMovieName)
	r.GET("/user/statistical", order.GetAllOrderStatistical)
	//customer
	r.POST("/customer/send/:email", customer.SendOtptoEmail)
	r.POST("/customer/verify/", customer.CheckOtpByEmail)
	r.POST("/customer/manager/register", customer.RegisterCustomersManager) //account admin => export include webhook
	r.POST("/customer/manager/login", customer.Login)
	r.POST("/customer/staff/register", customer.CreateAccountAdminManagerForStaff) // account staff
	r.GET("customer/staff/getall", customer.GetAllStaff)
	r.DELETE("/user/staff/delete", customer.DeleteStaffByName)
	r.GET("/supper_admin/create", customer.CreateAccountAdmin) // pessmiss super admin
	r.POST("/customer/check", customer.CheckAccountAndSendOtp) //sen opt
	r.PUT("/customer/reset", customer.VerifyOtpByEmailAndResetPassword)
	//
	r.POST("/customer/user/register", customer.RegisterAccountCustomer) //create account customer
	r.PUT("/customer/user/update", customer.UpdateProfileCustomerByUserName)
	r.GET("/customer/user/profile", customer.FindCustomersByUsename)
	r.GET("/customer/auth2", customer.CreateTokenRespWhenLoginWithEmail)
	r.PUT("/customer/reset/password", customer.UpdatePassWord) //reset password
	//show time
	r.POST("/use/add/time", showTime.AddShowTime)
	r.DELETE("/use/delete/time", showTime.DeleteShowTime) //ko dung
	r.GET("/user/getlist/time", showTime.GetShowTimeByTicketId)
	r.GET("/user/getlist/time/admin", showTime.GetShowTimeByTicketIdForAdmin)
	r.GET("/customer/detail/showtime", showTime.DetailShowTime)
	r.DELETE("use/delete/byid", showTime.DeleteShowTimeById)
	r.GET("/use/showtime", showTime.GetShowTimeById)
	r.PUT("/user/showtime/update", showTime.UpdateShowTimeById)
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
	//cart
	r.POST("/cart/add", cart.AddCart)
	r.GET("/cart/getlist", cart.FindByFormcart)
	r.PUT("/cart/update", cart.UpdateCartById)
	r.DELETE("/cart/delete/:id", cart.DeleteCartById)
	// Thêm công việc vào lịch để chạy mỗi 15 phút = 900s
	// scheduler := cron.New()
	// err := scheduler.AddFunc("*/900 * * * *", func() {
	// 	resp, err := http.Get("http://localhost:8080/manager/user/trigger")
	// 	if err != nil {
	// 		// Xử lý lỗi khi gọi API
	// 		log.Error(err, "error controller")
	// 		return
	// 	}
	// 	defer resp.Body.Close()
	// 	// Xử lý phản hồi nếu cần
	// })
	// if err != nil {
	// 	log.Error(err, "error")
	// }
	// // Bắt đầu lịch sau khi thêm công việc vào
	// scheduler.Start()

	return &ApiRouter{
		Engine: engine,
	}
}
