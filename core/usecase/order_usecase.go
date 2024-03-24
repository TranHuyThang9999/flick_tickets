package usecase

import (
	"context"
	"flick_tickets/common/enums"
	"flick_tickets/common/log"
	"flick_tickets/common/utils"
	"flick_tickets/core/domain"
	"flick_tickets/core/entities"
	"flick_tickets/core/mapper"
	"fmt"
)

type UseCaseOrder struct {
	order   domain.RepositoryOrder
	tickets domain.RepositoryTickets
	trans   domain.RepositoryTransaction
	aes     *UseCaseAes
}

func NewUsecaseOrder(
	order domain.RepositoryOrder,
	tickets domain.RepositoryTickets,
	trans domain.RepositoryTransaction,
	aes *UseCaseAes,

) *UseCaseOrder {
	return &UseCaseOrder{
		order:   order,
		tickets: tickets,
		trans:   trans,
		aes:     aes,
	}
}
func (u *UseCaseOrder) RegisterTicket(ctx context.Context, req *entities.OrdersReq) (*entities.OrdersResponseResgister, error) {
	log.Infof("req : ", req)
	// Bắt đầu giao dịch
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
				Code:    enums.SUCCESS_CODE,
				Message: "Không tìm thấy vé",
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
		ID:          utils.GenerateUniqueKey(),
		Email:       req.Email,
		Seats:       req.Seats,
		TicketID:    ticket.ID,
		Showtime:    ticket.Showtime,
		ReleaseDate: ticket.ReleaseDate,
		Description: ticket.Description,
		Status:      ticket.Status,
		Sale:        ticket.Sale,
		Price:       ticket.Price,
		CreatedAt:   ticket.CreatedAt,
	})
	if err != nil {
		return &entities.OrdersResponseResgister{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: err.Error(),
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

	resp, err := u.aes.GeneratesTokenWithAesToQrCodeAndSendQrWithEmail(&entities.TokenRequestSendQrCode{
		Content:   fmt.Sprintf("%v-%v-%v-%v", ticket.ID, ticket.Showtime, ticket.Price+(ticket.Price*float64(ticket.Sale)), req.Seats),
		FromEmail: req.Email,
		Title:     "Xin gửi bạn mã QR Code vé xem phim tại dạp vui lòng không để lộ ra ngoài",
	})
	if resp.Result.Code != 0 {
		return &entities.OrdersResponseResgister{
			Result: entities.Result{
				Code:    resp.Result.Code,
				Message: resp.Result.Message,
			},
		}, nil
	}

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
