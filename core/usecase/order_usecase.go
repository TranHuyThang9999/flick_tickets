package usecase

import (
	"context"
	"flick_tickets/common/enums"
	"flick_tickets/common/log"
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
}

func NewUsecaseOrder(
	order domain.RepositoryOrder,
	tickets domain.RepositoryTickets,
	trans domain.RepositoryTransaction,
	aes *UseCaseAes,
	menory cache.RepositoryCache,
	showTime domain.RepositoryShowTime,

) *UseCaseOrder {
	return &UseCaseOrder{
		order:    order,
		tickets:  tickets,
		trans:    trans,
		aes:      aes,
		menory:   menory,
		showTime: showTime,
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

	// Lấy thông tin vé từ cơ sở dữ liệu
	ticket, err := u.tickets.GetTicketById(ctx, req.TicketId)
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
	if ticket.Quantity <= 0 {
		return &entities.OrdersResponseResgister{
			Result: entities.Result{
				Code:    enums.SUCCESS_CODE,
				Message: "Hết vé",
			},
		}, nil
	}

	// Đăng ký vé
	err = u.order.RegisterTicket(ctx, tx, &domain.Orders{
		ID:          idOrder,
		Email:       req.Email,
		Seats:       req.Seats,
		TicketID:    ticket.ID,
		ReleaseDate: ticket.ReleaseDate,
		Description: ticket.Description,
		MovieTime:   req.MovieTime,
		Status:      ticket.Status, //need update
		Sale:        ticket.Sale,
		Price:       ticket.Price,
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
	listSeat, err := mapper.ParseToIntSlice(ticket.SelectedSeat)
	if err != nil {
		tx.Rollback()
		return &entities.OrdersResponseResgister{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE, //
				Message: err.Error(),
			},
		}, nil
	}

	for _, v := range listSeat {
		if v == req.Seats {
			tx.Rollback()
			return &entities.OrdersResponseResgister{
				Result: entities.Result{
					Code:    enums.TICKETS_REGISTERED_ERR_CODE,
					Message: enums.TICKETS_REGISTERED_ERR_MESS,
				},
			}, nil
		}
	}

	if err != nil {
		tx.Rollback()
		return &entities.OrdersResponseResgister{
			Result: entities.Result{
				Code:    enums.ERROR_CONVERT_JSON_CODE,
				Message: enums.ERROR_CONVERT_JSON_MESS,
			},
		}, nil
	}
	defer func() *entities.OrdersResponseResgister {
		log.Infof("id order : ", idOrder)
		log.Infof("req : ", req)
		resp, err := u.aes.GeneratesTokenWithAesToQrCodeAndSendQrWithEmail(&entities.TokenRequestSendQrCode{
			Content:   strconv.FormatInt(idOrder, 10),
			FromEmail: req.Email,
			Title:     "Xin gửi bạn mã QR Code vé xem phim tại dạp vui lòng không để lộ ra ngoài",
			Order: &entities.OrderSendTicketToEmail{
				ID:          idOrder, // ko co
				MoviceName:  ticket.Name,
				ReleaseDate: req.MovieTime,
				Price:       ticket.Price,
				Seats:       req.Seats, //ko co
				CinemaName:  req.CinemaName,
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

	listSeat = append(listSeat, req.Seats)
	// Cập nhật số lượng vé sau khi đăng ký
	ticketsAfter := ticket.Quantity - 1
	err = u.tickets.UpdateTicketQuantity(ctx, tx, ticket.ID, ticketsAfter)
	if err != nil {
		tx.Rollback()
		return &entities.OrdersResponseResgister{
			Result: entities.Result{
				Code:    enums.CONVERT_STRING_TO_ARRAY_CODE,
				Message: enums.CONVERT_STRING_TO_ARRAY_MESS,
			},
		}, nil
	}

	stringSeat := mapper.ConvertIntArrayToString(listSeat)

	err = u.tickets.UpdateTicketSelectedSeat(ctx, tx, ticket.ID, stringSeat)
	if err != nil {
		tx.Rollback()
		return &entities.OrdersResponseResgister{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: err.Error(),
			},
		}, nil
	}
	err = u.menory.SetObjectById(ctx, strconv.FormatInt(req.TicketId, 10), &domain.Tickets{
		ID:           ticket.ID,
		Name:         ticket.Name,
		Price:        ticket.Price,
		MaxTicket:    ticket.MaxTicket,
		Quantity:     ticketsAfter,
		Description:  ticket.Description,
		Sale:         ticket.Sale,
		ReleaseDate:  ticket.ReleaseDate,
		Status:       ticket.Status, //need update
		SelectedSeat: stringSeat,
		CreatedAt:    ticket.CreatedAt,
		UpdatedAt:    utils.GenerateTimestamp(),
	})
	if err != nil {
		tx.Rollback()
		return &entities.OrdersResponseResgister{
			Result: entities.Result{
				Code:    enums.CACHE_ERR_CODE,
				Message: enums.CACHE_ERR_MESS,
			},
		}, nil
	}
	// Commit giao dịch
	err = tx.Commit().Error
	if err != nil {
		return &entities.OrdersResponseResgister{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: err.Error(),
			},
		}, nil
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
