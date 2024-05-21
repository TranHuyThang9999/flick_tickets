package usecase

import (
	"context"
	"encoding/json"
	"flick_tickets/common/enums"
	"flick_tickets/common/log"
	"flick_tickets/common/utils"
	"flick_tickets/core/domain"
	"flick_tickets/core/entities"
	"flick_tickets/core/events/caching/cache"
	"flick_tickets/core/mapper"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type UseCaseOrder struct {
	order    domain.RepositoryOrder
	tickets  domain.RepositoryTickets
	trans    domain.RepositoryTransaction
	aes      *UseCaseAes
	menory   cache.RepositoryCache
	showTime domain.RepositoryShowTime
	cinema   domain.RepositoryCinemas
}

func NewUsecaseOrder(
	order domain.RepositoryOrder,
	tickets domain.RepositoryTickets,
	trans domain.RepositoryTransaction,
	aes *UseCaseAes,
	menory cache.RepositoryCache,
	showTime domain.RepositoryShowTime,
	cinema domain.RepositoryCinemas,

) *UseCaseOrder {
	return &UseCaseOrder{
		order:    order,
		tickets:  tickets,
		trans:    trans,
		aes:      aes,
		menory:   menory,
		showTime: showTime,
		cinema:   cinema,
	}
}
func (u *UseCaseOrder) RegisterTicket(ctx context.Context, req *entities.OrdersReq) (*entities.OrdersResponseResgister, error) {

	//idOrder := utils.GenerateUniqueKey()
	idOrder := req.Id
	// init transaction
	tx, err := u.trans.BeginTransaction(ctx)
	if err != nil {
		return &entities.OrdersResponseResgister{
			Result: entities.Result{
				Code:    enums.TRANSACTION_INVALID_CODE,
				Message: enums.TRANSACTION_INVALID_MESS,
			},
		}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()
	log.Infof("req  : ", req.ShowTimeId)
	showTimeForUserRegisterOrder, err := u.showTime.GetInformationShowTimeForTicketByTicketId(ctx, req.ShowTimeId)
	if err != nil {
		tx.Rollback()
		return &entities.OrdersResponseResgister{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE, //
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	log.Infof("data : ", showTimeForUserRegisterOrder)
	if showTimeForUserRegisterOrder.Quantity <= 0 {
		return &entities.OrdersResponseResgister{
			Result: entities.Result{
				Code:    enums.ORDER_REGISTER_TICKET_CODE,
				Message: enums.ORDER_REGISTER_TICKET_MESS,
			},
		}, nil
	}

	listSeat, err := mapper.ParseToIntSlice(showTimeForUserRegisterOrder.SelectedSeat) // string to array 12,32,432 =>[12,32,432]
	if err != nil {
		tx.Rollback()
		return &entities.OrdersResponseResgister{
			Result: entities.Result{
				Code:    enums.CONVERT_STRING_TO_ARRAY_CODE, //
				Message: enums.CONVERT_STRING_TO_ARRAY_MESS,
			},
		}, nil
	}

	listSeatsChoice, err := mapper.ParseToIntSlice(req.Seats)
	if err != nil {
		tx.Rollback()
		return &entities.OrdersResponseResgister{
			Result: entities.Result{
				Code:    enums.CONVERT_STRING_TO_ARRAY_CODE, //
				Message: enums.CONVERT_STRING_TO_ARRAY_MESS,
			},
		}, nil
	}
	log.Infof("list : ", listSeatsChoice, listSeat)
	checkDuplicate := mapper.HasDuplicateList(listSeat, listSeatsChoice)
	log.Infof("status", checkDuplicate)
	if !checkDuplicate {
		return &entities.OrdersResponseResgister{
			Result: entities.Result{
				Code:    enums.SHOW_TIME_ORDER_CODE,
				Message: enums.SHOW_TIME_ORDER_MESS,
			},
		}, nil
	}
	if err != nil {
		tx.Rollback()
		return &entities.OrdersResponseResgister{
			Result: entities.Result{
				Code:    enums.ERROR_CONVERT_JSON_CODE,
				Message: enums.ERROR_CONVERT_JSON_MESS,
			},
		}, nil
	} //
	// Lấy thông tin vé từ cơ sở dữ liệu
	ticket, err := u.tickets.GetTicketById(ctx, showTimeForUserRegisterOrder.TicketID)
	if err != nil {
		return &entities.OrdersResponseResgister{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, err
	}
	if ticket == nil {
		return &entities.OrdersResponseResgister{
			Result: entities.Result{
				Code:    enums.DATA_EMPTY_ERR_CODE,
				Message: enums.DATA_EMPTY_ERR_MESS,
			},
		}, nil
	}
	addressCinema, err := u.cinema.GetAllCinemaByName(ctx, showTimeForUserRegisterOrder.CinemaName)
	if err != nil {
		tx.Rollback()
		return &entities.OrdersResponseResgister{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	adsressRespnseData := entities.CinemaRespSendCustomer{
		CinemaName:     addressCinema.CinemaName,
		Description:    addressCinema.Description,
		Conscious:      addressCinema.Conscious,
		District:       addressCinema.District,
		Commune:        addressCinema.Commune,
		AddressDetails: addressCinema.AddressDetails,
	}
	addressCinemaTypeJson, err := json.Marshal(adsressRespnseData) //send email
	if err != nil {
		tx.Rollback()
		return &entities.OrdersResponseResgister{
			Result: entities.Result{
				Code:    enums.ERROR_CONVERT_JSON_CODE,
				Message: enums.ERROR_CONVERT_JSON_MESS,
			},
		}, nil
	}
	price := float64(ticket.Price) * float64(len(listSeatsChoice))

	// Đăng ký vé
	err = u.order.RegisterTicket(ctx, tx, &domain.Orders{
		ID:             idOrder,
		Email:          req.Email,
		Seats:          req.Seats,
		ShowTimeID:     req.ShowTimeId,
		ReleaseDate:    ticket.ReleaseDate,
		Description:    ticket.Description,
		MovieTime:      showTimeForUserRegisterOrder.MovieTime, //thoi gian chieu
		Status:         enums.ORDER_INIT,                       //need update
		Sale:           ticket.Sale,
		Price:          price,
		AddressDetails: string(addressCinemaTypeJson),
		CinemaName:     showTimeForUserRegisterOrder.CinemaName,
		MovieName:      ticket.Name,
		UpdatedAt:      utils.GenerateTimestamp(),
		CreatedAt:      utils.GenerateTimestamp(),
	})
	if err != nil {
		tx.Rollback()
		return &entities.OrdersResponseResgister{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}

	listSeat = append(listSeat, listSeatsChoice...)

	stringConvertEdSeat := mapper.ConvertIntArrayToString(listSeat) // listShowTime
	// Cập nhật số lượng vé sau khi đăng ký
	err = u.showTime.UpdateQuantitySeat(ctx, tx,
		showTimeForUserRegisterOrder.ID,
		showTimeForUserRegisterOrder.Quantity-len(listSeatsChoice),
		stringConvertEdSeat)
	if err != nil {
		tx.Rollback()
		return &entities.OrdersResponseResgister{
			Result: entities.Result{ //thay
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}

	err = tx.Commit().Error
	if err != nil {
		return &entities.OrdersResponseResgister{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, err
	}

	return &entities.OrdersResponseResgister{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		OrderId: idOrder,
	}, nil
}

func (u *UseCaseOrder) SendticketAfterPayment(ctx context.Context, req *entities.OrderSendTicketAfterPaymentReq) (*entities.OrderSendTicketAfterPaymentResp, error) {
	//send email
	log.Infof("req : ", req)
	order, err := u.order.GetOrderById(ctx, req.OrderId)
	if err != nil {
		return &entities.OrderSendTicketAfterPaymentResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	showTime, err := u.showTime.GetInformationShowTimeForTicketByTicketId(ctx, order.ShowTimeID)
	if err != nil {
		return &entities.OrderSendTicketAfterPaymentResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	log.Infof("showtime : ", showTime)
	ticket, err := u.tickets.GetTicketById(ctx, showTime.TicketID)
	if err != nil {
		return &entities.OrderSendTicketAfterPaymentResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	err = u.order.UpsertOrder(ctx, order.Email, order.ID, enums.ORDER_SUCESS)
	if err != nil {
		return &entities.OrderSendTicketAfterPaymentResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	resp, err := u.aes.GeneratesTokenWithAesToQrCodeAndSendQrWithEmail(&entities.TokenRequestSendQrCode{
		Content:   strconv.FormatInt(order.ID, 10),
		FromEmail: order.Email,
		Title:     "Xin gửi bạn mã QR Code vé xem phim tại dạp vui lòng không để lộ ra ngoài",
		Order: &entities.OrderSendTicketToEmail{
			ID:         order.ID,
			MoviceName: ticket.Name, // ten phim
			Price:      order.Price,
			Seats:      order.Seats,
			CinemaName: order.AddressDetails, // ten phong
			MovieTime:  showTime.MovieTime,
		},
	})
	log.Infof("data : ", resp)
	if err != nil || resp.Result.Code != 0 {
		return &entities.OrderSendTicketAfterPaymentResp{
			Result: entities.Result{
				Code:    enums.SEND_EMAIL_ERR_CODE,
				Message: enums.SEND_EMAIL_ERR_MESS,
			},
		}, nil
	}

	// Trả về nil nếu không có lỗi
	return &entities.OrderSendTicketAfterPaymentResp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
	}, nil
}

func (u *UseCaseOrder) GetOrderById(ctx context.Context, id string) (*entities.OrdersResponseGetById, error) {

	numberId, err := strconv.Atoi(id)

	if err != nil {
		return &entities.OrdersResponseGetById{
			Result: entities.Result{
				Code:    enums.CONVERT_TO_NUMBER_CODE,
				Message: enums.CONVERT_TO_NUMBER_MESS,
			},
		}, nil
	}

	order, err := u.order.GetOrderById(ctx, int64(numberId))
	if err != nil {
		return &entities.OrdersResponseGetById{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
			Created: int64(utils.GenerateTimestamp()),
		}, nil
	}
	if order == nil {
		return &entities.OrdersResponseGetById{
			Result: entities.Result{
				Code:    enums.DATA_EMPTY_ERR_CODE,
				Message: enums.DATA_EMPTY_ERR_MESS,
			},
			Created: int64(utils.GenerateTimestamp()),
		}, nil
	}
	return &entities.OrdersResponseGetById{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Orders:  order,
		Created: int64(utils.GenerateTimestamp()),
	}, nil

}
func (u *UseCaseOrder) UpsertOrderById(ctx context.Context, req *entities.OrderReqUpSert) (*entities.OrderRespUpSert, error) {

	err := u.order.UpsertOrder(ctx, req.Email, req.Id, enums.ORDER_SUCESS)
	if err != nil {
		return &entities.OrderRespUpSert{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	return &entities.OrderRespUpSert{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
	}, nil
}
func (u *UseCaseOrder) UpdateOrderWhenCancel(ctx context.Context, req *entities.OrderCancelBtyIdreq) (*entities.OrderCancelBtyIdresp, error) {

	tx, err := u.trans.BeginTransaction(ctx)
	if err != nil {
		return &entities.OrderCancelBtyIdresp{
			Result: entities.Result{
				Code:    enums.TRANSACTION_INVALID_CODE,
				Message: enums.TRANSACTION_INVALID_MESS,
			},
		}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	err = u.order.UpdateOrderWhenCancel(ctx, tx, req.OrderId, enums.ORDER_CANCEL)
	if err != nil {
		tx.Rollback()
		return &entities.OrderCancelBtyIdresp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}

	order, err := u.order.GetOrderById(ctx, req.OrderId)
	if err != nil {
		tx.Rollback()
		return &entities.OrderCancelBtyIdresp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}

	showtimeById, err := u.showTime.GetInformationShowTimeForTicketByTicketId(ctx, order.ShowTimeID)
	if err != nil {
		tx.Rollback()
		return &entities.OrderCancelBtyIdresp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	// newListSeatsGetShowTime := make([]int, 0)
	// newListSeatsGetPutOrder := make([]int, 0)
	newListSeatsGetShowTime, err := mapper.ParseToIntSlice(showtimeById.SelectedSeat)
	if err != nil {
		return &entities.OrderCancelBtyIdresp{
			Result: entities.Result{
				Code:    enums.CONVERT_STRING_TO_ARRAY_CODE,
				Message: enums.CONVERT_STRING_TO_ARRAY_MESS,
			},
		}, nil
	}
	newListSeatsGetPutOrder, err := mapper.ParseToIntSlice(order.Seats)
	if err != nil {
		return &entities.OrderCancelBtyIdresp{
			Result: entities.Result{
				Code:    enums.CONVERT_STRING_TO_ARRAY_CODE,
				Message: enums.CONVERT_STRING_TO_ARRAY_MESS,
			},
		}, nil
	}
	//newListSeatsGetShowTime = append(newListSeatsGetShowTime, newListSeatsGetPutOrder...)
	listremoveDuplicates := mapper.RemoveDuplicates(newListSeatsGetShowTime, newListSeatsGetPutOrder)
	err = u.showTime.UpdateQuantitySeat(ctx, tx,
		order.ShowTimeID,
		showtimeById.Quantity+len(mapper.ConvertListToStringSlice(order.Seats)),
		mapper.ConvertIntArrayToString(listremoveDuplicates),
	)

	if err != nil {
		tx.Rollback()
		return &entities.OrderCancelBtyIdresp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	tx.Commit()
	return &entities.OrderCancelBtyIdresp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
	}, nil
}
func (u *UseCaseOrder) GetAllOrder(ctx context.Context, req *domain.OrdersReqByForm) (*entities.OrderGetAll, error) {
	orders, err := u.order.GetAllOrder(ctx, req)
	if err != nil {
		return &entities.OrderGetAll{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	return &entities.OrderGetAll{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Total:  len(orders),
		Orders: orders,
	}, nil
}

func (u *UseCaseOrder) GetOrderByIdFromPayOs(ctx context.Context, orderID string) (*entities.PayMentResponseCheckOrder, error) {
	url := fmt.Sprintf("https://api-merchant.payos.vn/v2/payment-requests/%s", orderID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error(err, "error")
		return nil, err
	}
	log.Info(orderID)
	req.Header.Set("x-client-id", "c84c857d-160c-456a-91f2-384526d7a360")
	req.Header.Set("x-api-key", "f74461b1-d7d3-4fca-b918-fcb39524ce8c")
	req.Header.Set("Cookie", "connect.sid=s%3A-Sat8d9c-WFoxLgE3cJZTb9bi3oSwFC2.uiWQpbtmdJc8ARx1PevsohQW62U4QiaOgBPfX85%2F91s")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err, "error")
		return nil, err
	}
	defer resp.Body.Close()

	// Cập nhật phần "Cookie" từ phản hồi API trước đó
	newCookie := resp.Header.Get("Set-Cookie")
	req.Header.Set("Cookie", newCookie)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err, "error")
		return nil, err
	}

	var response entities.PayMentResponseCheckOrder
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Error(err, "error convert string to json")
		return nil, fmt.Errorf(enums.ERROR_CONVERT_JSON_MESS+"%s", err)
	}

	if response.Code == "101" {
		log.Info("Mã thanh toán không tồn tại")
		a := []entities.Transaction{}
		return &entities.PayMentResponseCheckOrder{
			Code: "101",
			Desc: "Mã thanh toán không tồn tại",
			Data: entities.Data{
				ID:                 orderID,
				OrderCode:          0,
				Amount:             0,
				AmountPaid:         0,
				AmountRemaining:    0,
				Status:             "",
				CreatedAt:          "",
				Transactions:       a,
				CanceledAt:         nil,
				CancellationReason: nil,
			},
			Signature: "",
		}, nil
	}

	if response.Data.Status == "EXPIRED" {
		a := []entities.Transaction{}
		log.Info("Đơn hàng đã quá hạn thời gian")
		return &entities.PayMentResponseCheckOrder{
			Code: "103",
			Desc: "Đơn hàng đã quá hạn thời gian",
			Data: entities.Data{
				ID:                 "",
				OrderCode:          mapper.ConvertStringToInt(orderID),
				Amount:             0,
				AmountPaid:         0,
				AmountRemaining:    0,
				Status:             "",
				CreatedAt:          "",
				Transactions:       a,
				CanceledAt:         nil,
				CancellationReason: nil,
			},
		}, nil
	} else if response.Data.Status == "CANCELLED" {
		a := []entities.Transaction{}
		return &entities.PayMentResponseCheckOrder{
			Code: "102",
			Desc: "Đơn hàng đã hủy",
			Data: entities.Data{
				ID:                 "0",
				OrderCode:          mapper.ConvertStringToInt(orderID),
				Amount:             0,
				AmountPaid:         0,
				AmountRemaining:    0,
				Status:             "",
				CreatedAt:          "",
				Transactions:       a,
				CanceledAt:         nil,
				CancellationReason: nil,
			},
			Signature: "",
		}, nil
	}

	return &response, nil
}

func (u *UseCaseOrder) TriggerOrder(ctx context.Context) error {

	var listOrderId = make([]int64, 0)

	order, err := u.order.TriggerOrder(ctx)
	if err != nil {
		log.Error(err, "error check order")
		return err
	}
	if len(order) == 0 {
		log.Info("data emptly")
		return nil
	}
	for _, v := range order {
		listOrderId = append(listOrderId, v.ID)
	}

	for i := 0; i < len(listOrderId); i++ {
		resp, err := u.GetOrderByIdFromPayOs(ctx, strconv.FormatInt(listOrderId[i], 10))
		if err != nil {
			log.Error(err, "error check list")
			return fmt.Errorf("error check order %v", err)
		}
		if resp.Code == "103" { // qua han ko thanh toan
			_, err := u.UpdateOrderWhenCancel(ctx, &entities.OrderCancelBtyIdreq{
				OrderId: listOrderId[i],
			})
			if err != nil {
				return fmt.Errorf("error check order %v", err)
			}
			return nil
		} else if resp.Code == "101" || resp.Code == "102" {
			return fmt.Errorf("error check order %v ", resp)
		} else { //thanh toan thanh cong
			_, err := u.SendticketAfterPayment(ctx, &entities.OrderSendTicketAfterPaymentReq{
				OrderId: listOrderId[i],
			})
			log.Info("update ok")
			if err != nil {
				return fmt.Errorf("error check order %v", err)
			}
		}
	}
	// 101 ko ton tai 102 da huy
	log.Infof("trigger ok")
	return nil
}
func (u *UseCaseOrder) OrderHistory(ctx context.Context, req *entities.OrderHistoryReq) (*entities.OrderHistoryResp, error) {

	var listOrdersResp = make([]*entities.OrderHistoryEntities, 0)

	listOrders, err := u.order.GetListOrderHistoeryByEmail(ctx, req.Email)
	if err != nil {
		return &entities.OrderHistoryResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	if len(listOrders) == 0 {
		return &entities.OrderHistoryResp{
			Result: entities.Result{
				Code:    enums.DATA_EMPTY_ERR_CODE,
				Message: enums.DATA_EMPTY_ERR_MESS,
			},
		}, nil
	}

	for i := 0; i < len(listOrders); i++ {

		listOrdersResp = append(listOrdersResp, &entities.OrderHistoryEntities{
			ID:             listOrders[i].ID,
			Email:          req.Email,
			ReleaseDate:    listOrders[i].ReleaseDate,
			Description:    listOrders[i].Description,
			Status:         listOrders[i].Status,
			Price:          listOrders[i].Price,
			Seats:          listOrders[i].Seats,
			MovieTime:      listOrders[i].MovieTime,
			AddressDetails: listOrders[i].AddressDetails,
			MovieName:      listOrders[i].MovieName,  // ten phim
			CinemaName:     listOrders[i].CinemaName, // ten rap
			CreatedAt:      listOrders[i].CreatedAt,
		})
	}

	return &entities.OrderHistoryResp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		OrderHistoryEntities: listOrdersResp,
	}, nil
}
func (u *UseCaseOrder) OrderRevenueByMovieName(ctx context.Context, req *entities.OrderRevenueReq) (*entities.OrderRevenueResp, error) {

	sum, err := u.order.GetrevenueOrderByMovieNameAndTimeDistance(ctx, &domain.OrderRevenue{
		CinemaName:        req.CinemaName,
		MovieName:         req.MovieName,
		TimeDistanceStart: req.TimeDistanceStart,
		TimeDistanceEnd:   req.TimeDistanceEnd,
	})
	if err != nil {
		return &entities.OrderRevenueResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	return &entities.OrderRevenueResp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Sum: sum,
	}, nil
}
func (u *UseCaseOrder) GetAllMovieNameFromOrder(ctx context.Context) (*entities.OrderGetAllFromOrderResp, error) {
	orders, err := u.order.GetAllMovieNameFromOrder(ctx)
	if err != nil {
		return &entities.OrderGetAllFromOrderResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}

	if len(orders) == 0 {
		return &entities.OrderGetAllFromOrderResp{
			Result: entities.Result{
				Code:    enums.DATA_EMPTY_ERR_CODE,
				Message: enums.DATA_EMPTY_ERR_MESS,
			},
		}, nil
	}

	return &entities.OrderGetAllFromOrderResp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Orders: orders,
	}, nil
}
func (u *UseCaseOrder) GetAllCinemaByMovieName(ctx context.Context, cinema_name string) (*entities.OrderGetAllFromOrderByCinemaNameResp, error) {
	orders, err := u.order.GetAllCinemaByMovieName(ctx, cinema_name)
	if err != nil {
		return &entities.OrderGetAllFromOrderByCinemaNameResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}

	if len(orders) == 0 {
		return &entities.OrderGetAllFromOrderByCinemaNameResp{
			Result: entities.Result{
				Code:    enums.DATA_EMPTY_ERR_CODE,
				Message: enums.DATA_EMPTY_ERR_MESS,
			},
		}, nil
	}

	return &entities.OrderGetAllFromOrderByCinemaNameResp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Orders: orders,
	}, nil

}
func (u *UseCaseOrder) StatisticalOrder(ctx context.Context, req *entities.OrderStatisticalReq) (*entities.OrderStatisticalResp, error) {

	var orderResp = make([]*domain.Orders, 0)
	listOrder, err := u.order.GetAllOrder(ctx, &domain.OrdersReqByForm{
		Status: 9,
	})
	log.Infof("len ", len(listOrder))
	if err != nil {
		return &entities.OrderStatisticalResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	for i := 0; i < len(listOrder); i++ {
		if listOrder[i].CreatedAt >= req.StartTime && listOrder[i].CreatedAt <= req.EndTime {
			orderResp = append(orderResp, listOrder[i])
		} else if req.EndTime == req.StartTime {
			if listOrder[i].CreatedAt == req.EndTime {
				orderResp = append(orderResp, listOrder[i])
			}
		}
	}
	return &entities.OrderStatisticalResp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Orders: orderResp,
	}, nil
}
