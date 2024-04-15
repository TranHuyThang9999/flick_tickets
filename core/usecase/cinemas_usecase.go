package usecase

import (
	"context"
	"flick_tickets/common/enums"
	"flick_tickets/common/utils"
	"flick_tickets/core/domain"
	"flick_tickets/core/entities"
)

type UseCaseCinemas struct {
	cm domain.RepositoryCinemas
}

func NewUseCaseCinemas(cm domain.RepositoryCinemas) *UseCaseCinemas {
	return &UseCaseCinemas{
		cm: cm,
	}
}
func (c *UseCaseCinemas) AddCinemas(ctx context.Context, req *entities.CinemasReq) (*entities.CinemasRes, error) {

	cinemaByName, err := c.cm.GetAllCinemaByName(ctx, req.CinemaName)
	if err != nil {
		return &entities.CinemasRes{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, err
	}
	if cinemaByName != nil {
		return &entities.CinemasRes{
			Result: entities.Result{
				Code:    enums.ROOM_EXSTIS_CODE,
				Message: enums.ROOM_EXSTIS_MESS,
			},
		}, nil
	}
	err = c.cm.AddCinema(ctx, &domain.Cinemas{
		Id:             utils.GenerateUniqueKey(),
		CinemaName:     req.CinemaName,
		Description:    req.Description,
		Conscious:      req.Conscious,
		District:       req.District,
		Commune:        req.Commune,
		AddressDetails: req.AddressDetails,
	})
	if err != nil {
		return &entities.CinemasRes{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, err
	}

	return &entities.CinemasRes{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
	}, nil

}
func (c *UseCaseCinemas) GetAllCinema(ctx context.Context) (*entities.CinemasResGetAll, error) {
	resp, err := c.cm.GetAllCinema(ctx)
	if err != nil {
		return &entities.CinemasResGetAll{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, err
	}
	if len(resp) == 0 {
		return &entities.CinemasResGetAll{
			Result: entities.Result{
				Code:    enums.DATA_EMPTY_ERR_CODE,
				Message: enums.DATA_EMPTY_ERR_MESS,
			},
		}, nil
	}
	return &entities.CinemasResGetAll{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Cinemas: resp,
	}, nil
}
func (c *UseCaseCinemas) DeleteCinemaByName(ctx context.Context, name string) (*entities.CinemasRespDelete, error) {
	err := c.cm.DeleteCinemaByName(ctx, name)
	if err != nil {
		return &entities.CinemasRespDelete{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	return &entities.CinemasRespDelete{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
	}, nil
}
func (c *UseCaseCinemas) GetAllCinemaByName(ctx context.Context, name string) (*entities.CinemasRespGetByName, error) {
	cinemaByName, err := c.cm.GetAllCinemaByName(ctx, name)
	if err != nil {
		return &entities.CinemasRespGetByName{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, err
	}
	if cinemaByName == nil {
		return &entities.CinemasRespGetByName{
			Result: entities.Result{
				Code:    enums.DATA_EMPTY_ERR_CODE,
				Message: enums.DATA_EMPTY_ERR_MESS,
			},
		}, nil
	}
	return &entities.CinemasRespGetByName{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Cinema: cinemaByName,
	}, nil
}
func (c *UseCaseCinemas) UpdateColumnWidthHeightContainer(
	ctx context.Context,
	req *entities.CinemaReqUpdateSizeRoom,
) (*entities.CinemaRespUpdateSizeRoom, error) {
	err := c.cm.UpdateColumnWidthHeightContainer(ctx, &domain.CinemaReqUpdateSizeRoom{
		CinemaName:      req.CinemaName,
		WidthContainer:  req.WidthContainer,
		HeightContainer: req.HeightContainer,
	})
	if err != nil {
		return &entities.CinemaRespUpdateSizeRoom{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, err
	}
	return &entities.CinemaRespUpdateSizeRoom{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
	}, nil
}
