package usecase

import (
	"context"
	"encoding/json"
	"flick_tickets/common/enums"
	"flick_tickets/common/utils"
	"flick_tickets/core/domain"
	"flick_tickets/core/entities"
	"flick_tickets/core/events/caching/cache"
	"flick_tickets/core/mapper"
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

	idOrder := utils.GenerateUniqueKey()

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
	//showTime
	// keyshowTime, err := u.menory.KeyExists(ctx, string(req.ShowTimeId))
	// if err != nil {
	// 	tx.Rollback()
	// 	return &entities.OrdersResponseResgister{
	// 		Result: entities.Result{
	// 			Code:    enums.CACHE_ERR_CODE, //
	// 			Message: enums.CACHE_ERR_MESS,
	// 		},
	// 	}, nil
	// }
	// if keyshowTime {

	// }

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

	if showTimeForUserRegisterOrder.Quantity <= 0 {
		return &entities.OrdersResponseResgister{
			Result: entities.Result{
				Code:    enums.ORDER_REGISTER_TICKET_CODE,
				Message: enums.ORDER_REGISTER_TICKET_MESS,
			},
		}, nil
	}

	listSeat, err := mapper.ParseToIntSlice(showTimeForUserRegisterOrder.SelectedSeat)
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

	// for _, v := range listSeat {
	// 	if v == listSeat[] {
	// 		tx.Rollback()
	// 		return &entities.OrdersResponseResgister{
	// 			Result: entities.Result{
	// 				Code:    enums.TICKETS_REGISTERED_ERR_CODE,
	// 				Message: enums.TICKETS_REGISTERED_ERR_MESS,
	// 			},
	// 		}, nil
	// 	}
	// }

	if err != nil {
		tx.Rollback()
		return &entities.OrdersResponseResgister{
			Result: entities.Result{
				Code:    enums.ERROR_CONVERT_JSON_CODE,
				Message: enums.ERROR_CONVERT_JSON_MESS,
			},
		}, nil
	}
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
	addressCinemaTypeJson, err := json.Marshal(adsressRespnseData)
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
		ID:          idOrder,
		Email:       req.Email,
		Seats:       string(addressCinemaTypeJson),
		ShowTimeID:  req.ShowTimeId,
		ReleaseDate: ticket.ReleaseDate,
		Description: ticket.Description,
		MovieTime:   showTimeForUserRegisterOrder.MovieTime, //thoi gian chieu
		Status:      enums.ORDER_INIT,                       //need update
		Sale:        ticket.Sale,
		Price:       price,
		CreatedAt:   ticket.CreatedAt,
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

	//send email
	defer func() *entities.OrdersResponseResgister {
		resp, err := u.aes.GeneratesTokenWithAesToQrCodeAndSendQrWithEmail(&entities.TokenRequestSendQrCode{
			Content:   strconv.FormatInt(idOrder, 10),
			FromEmail: req.Email,
			Title:     "Xin gửi bạn mã QR Code vé xem phim tại dạp vui lòng không để lộ ra ngoài",
			Order: &entities.OrderSendTicketToEmail{
				ID:         idOrder,
				MoviceName: ticket.Name,
				Price:      ticket.Price,
				Seats:      req.Seats,
				CinemaName: string(addressCinemaTypeJson),
				MovieTime:  showTimeForUserRegisterOrder.MovieTime,
			},
		})

		if err != nil || resp.Result.Code != 0 {
			return &entities.OrdersResponseResgister{
				Result: entities.Result{
					Code:    enums.SEND_EMAIL_ERR_CODE,
					Message: enums.SEND_EMAIL_ERR_MESS,
				},
			}
		}

		// Trả về nil nếu không có lỗi
		return &entities.OrdersResponseResgister{
			Result: entities.Result{
				Code:    enums.SUCCESS_CODE,
				Message: enums.SUCCESS_MESS,
			},
		}
	}()

	listSeat = append(listSeat, listSeatsChoice...)

	stringConvertEdSeat := mapper.ConvertIntArrayToString(listSeat) // listShowTime
	// Cập nhật số lượng vé sau khi đăng ký
	err = u.showTime.UpdateQuantitySeat(ctx, tx,
		showTimeForUserRegisterOrder.ID,
		showTimeForUserRegisterOrder.Quantity-1,
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

	//update cache show time
	// err = u.menory.SetObjectById(ctx, strconv.FormatInt(showTimeForUserRegisterOrder.ID, 10), &domain.ShowTime{
	// 	ID:           showTimeForUserRegisterOrder.ID,
	// 	TicketID:     showTimeForUserRegisterOrder.TicketID,
	// 	CinemaName:   showTimeForUserRegisterOrder.CinemaName,
	// 	MovieTime:    showTimeForUserRegisterOrder.MovieTime,
	// 	SelectedSeat: showTimeForUserRegisterOrder.SelectedSeat,
	// 	Quantity:     showTimeForUserRegisterOrder.Quantity,
	// 	CreatedAt:    showTimeForUserRegisterOrder.CreatedAt,
	// 	UpdatedAt:    utils.GenerateTimestamp(),
	// })
	// if err != nil {
	// 	tx.Rollback()
	// 	return &entities.OrdersResponseResgister{
	// 		Result: entities.Result{
	// 			Code:    enums.CACHE_ERR_CODE,
	// 			Message: enums.CACHE_ERR_MESS,
	// 		},
	// 	}, nil
	// }
	// Commit giao dịch
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

	err := u.order.UpsertOrder(ctx, req.Id, enums.ORDER_SUCESS)
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
