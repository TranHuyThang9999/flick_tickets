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
	cart domain.RepositoryCarts
}

func NewUseCaseCart(cart domain.RepositoryCarts) *UseCaseCart {
	return &UseCaseCart{
		cart: cart,
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
	listCart, err := u.cart.FindCartByForm(ctx, req)
	if err != nil {
		return &entities.CartFindByFormResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	return &entities.CartFindByFormResp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Carts: listCart,
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
