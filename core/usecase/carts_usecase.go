package usecase

import (
	"context"
	"flick_tickets/common/enums"
	"flick_tickets/common/log"
	"flick_tickets/common/utils"
	"flick_tickets/core/domain"
	"flick_tickets/core/entities"
	"flick_tickets/core/mapper"
)

type UseCaseCart struct {
	cart      domain.RepositoryCarts
	show_time domain.RepositoryShowTime
	cinema    domain.RepositoryCinemas
	ticket    domain.RepositoryTickets
}

func NewUseCaseCart(
	cart domain.RepositoryCarts,
	show_time domain.RepositoryShowTime,
	cinema domain.RepositoryCinemas,
	ticket domain.RepositoryTickets,
) *UseCaseCart {
	return &UseCaseCart{
		cart:      cart,
		show_time: show_time,
		cinema:    cinema,
		ticket:    ticket,
	}
}

func (u *UseCaseCart) AddCart(ctx context.Context, req *entities.CartsAddReq) (*entities.CartsAddResp, error) {
	err := u.cart.AddCarts(ctx, &domain.Carts{
		Id:            utils.GenerateUniqueKey(),
		UserName:      req.UserName,
		ShowTimeId:    req.ShowTimeId,
		SeatsPosition: req.SeatsPosition,
		Price:         req.Price,
		CreatedAt:     utils.GenerateTimestamp(),
		UpdatedAt:     utils.GenerateTimestamp(),
	})
	log.Infof("req : ", req)
	if err != nil {
		return &entities.CartsAddResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	return &entities.CartsAddResp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
	}, nil
}

func (u *UseCaseCart) FindCartByForm(ctx context.Context, req *domain.CartFindByFormReq) (*entities.CartFindByFormResp, error) {

	var listShowTimeId = make([]int64, 0)
	var listCartResp = make([]*entities.Carts, 0)

	// Tìm kiếm giỏ hàng dựa trên yêu cầu
	listCart, err := u.cart.FindCartByForm(ctx, req)
	if err != nil {
		log.Errorf(err, "Error finding carts: %v")
		return &entities.CartFindByFormResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	if len(listCart) == 0 {
		return &entities.CartFindByFormResp{
			Result: entities.Result{
				Code:    enums.DATA_EMPTY_ERR_CODE,
				Message: enums.DATA_EMPTY_ERR_MESS,
			},
		}, nil
	}

	// Tạo map để lưu trữ thông tin rạp chiếu phim dựa trên tên rạp
	cinemaMap := make(map[int64]*domain.Cinemas)
	ticketMap := make(map[int64]*domain.Tickets)
	movieTimeMap := make(map[int64]int64)

	for _, cart := range listCart {
		listShowTimeId = append(listShowTimeId, cart.ShowTimeId)
	}
	listShowTime, err := u.show_time.GetListShowTimeByListId(ctx, listShowTimeId)
	if err != nil {
		log.Errorf(err, "Error getting show time information: %v")
		return &entities.CartFindByFormResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}

	for _, showTime := range listShowTime {
		// Kiểm tra xem thông tin rạp đã được lưu trữ trong map chưa
		if _, ok := cinemaMap[showTime.ID]; !ok {
			cinema, err := u.cinema.GetAllCinemaByName(ctx, showTime.CinemaName)
			if err != nil {
				log.Errorf(err, "Error getting cinema information: %v")
				return &entities.CartFindByFormResp{
					Result: entities.Result{
						Code:    enums.DB_ERR_CODE,
						Message: enums.DB_ERR_MESS,
					},
				}, nil
			}

			// Lưu thông tin rạp vào map
			cinemaMap[showTime.ID] = cinema
		}
		ticket, err := u.ticket.GetTicketById(ctx, showTime.TicketID)
		if err != nil {
			log.Errorf(err, "Error getting cinema information: %v")
			return &entities.CartFindByFormResp{
				Result: entities.Result{
					Code:    enums.DB_ERR_CODE,
					Message: enums.DB_ERR_MESS,
				},
			}, nil
		}
		ticketMap[showTime.ID] = ticket
		movieTimeMap[showTime.ID] = int64(showTime.MovieTime)
	}

	for _, cart := range listCart {
		cinema := cinemaMap[cart.ShowTimeId] // Lấy thông tin rạp từ map
		ticket := ticketMap[cart.ShowTimeId] //
		listMovieTime := movieTimeMap[cart.ShowTimeId]
		cart := &entities.Carts{
			Id:             cart.Id,
			ShowTimeId:     cart.ShowTimeId,
			SeatsPosition:  cart.SeatsPosition,
			Price:          cart.Price,
			MovieTime:      int(listMovieTime),
			CinemaName:     cinema.CinemaName,
			Description:    cinema.Description,
			Conscious:      cinema.Conscious,
			District:       cinema.District,
			Commune:        cinema.Commune,
			AddressDetails: cinema.AddressDetails,
			MovieName:      ticket.Name,
			MovieDuration:  ticket.MovieDuration,
			AgeLimit:       ticket.AgeLimit,
		}
		listCartResp = append(listCartResp, cart)
	}

	// Trả về kết quả thành công cùng với danh sách giỏ hàng đã được tìm thấy
	return &entities.CartFindByFormResp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Carts: listCartResp,
	}, nil
}

func (u *UseCaseCart) UpdateCartById(ctx context.Context, req *entities.CartsUpdateReq) (*entities.CartsUpdateResp, error) {
	err := u.cart.UpdateCartById(ctx, &domain.Carts{
		Id:            req.Id,
		UserName:      req.UserName,
		ShowTimeId:    req.ShowTimeId,
		SeatsPosition: req.SeatsPosition,
		UpdatedAt:     utils.GenerateTimestamp(),
	})
	if err != nil {
		return &entities.CartsUpdateResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	return &entities.CartsUpdateResp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
	}, nil
}

func (u *UseCaseCart) DeleteCartById(ctx context.Context, cartId string) (*entities.CartsDeleteResp, error) {

	err := u.cart.DeleteCartById(ctx, int64(mapper.ConvertStringToInt(cartId)))
	if err != nil {
		return &entities.CartsDeleteResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	return &entities.CartsDeleteResp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
	}, nil
}
