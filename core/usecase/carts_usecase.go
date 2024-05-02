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
}

func NewUseCaseCart(cart domain.RepositoryCarts, show_time domain.RepositoryShowTime, cinema domain.RepositoryCinemas) *UseCaseCart {
	return &UseCaseCart{
		cart:      cart,
		show_time: show_time,
		cinema:    cinema,
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

func (u *UseCaseCart) FindCartByForm(ctx context.Context, req *domain.CartFindByFormReq) (*entities.CartFindByFormResp, error) { //requied user_name

	var listShowTimeId = make([]int64, 0)
	var listCinemaName = make([]string, 0)
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

	for i := 0; i < len(listCart); i++ {
		listShowTimeId = append(listShowTimeId, listCart[i].ShowTimeId)
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
	for i := 0; i < len(listShowTime); i++ {
		listCinemaName = append(listCinemaName, listShowTime[i].CinemaName)
	}
	for i, cart := range listCart {
		cinema, err := u.cinema.GetAllCinemaByName(ctx, listCinemaName[i])
		if err != nil {
			return &entities.CartFindByFormResp{
				Result: entities.Result{
					Code:    enums.DB_ERR_CODE,
					Message: enums.DB_ERR_MESS,
				},
			}, nil
		}

		cart := &entities.Carts{
			Id:             cart.Id,
			ShowTimeId:     cart.ShowTimeId,
			SeatsPosition:  cart.SeatsPosition,
			Price:          cart.Price,
			MovieTime:      0, // Không có thông tin thời gian phim, bạn có thể cập nhật sau
			CinemaName:     cinema.CinemaName,
			Description:    cinema.Description,
			Conscious:      cinema.Conscious,
			District:       cinema.Description, // Lưu ý: có vẻ như có lỗi ở đây
			Commune:        cinema.Commune,
			AddressDetails: cinema.AddressDetails,
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
