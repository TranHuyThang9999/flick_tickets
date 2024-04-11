package usecase

import (
	"context"
	"flick_tickets/common/enums"
	"flick_tickets/common/utils"
	"flick_tickets/core/domain"
	"flick_tickets/core/entities"
)

type UseCaseShowTime struct {
	st domain.RepositoryShowTime
}

func NewUseCaseShowTime(
	st domain.RepositoryShowTime,
) *UseCaseShowTime {
	return &UseCaseShowTime{
		st: st,
	}
}
func (s *UseCaseShowTime) AddShowTime(ctx context.Context, req *entities.ShowTimeAddReq) (*entities.ShowTimeAddResponse, error) {

	resp, err := s.st.GetTimeUseCheckAddTicket(ctx, &domain.ShowTimeCheckList{
		//	TicketID:   req.TicketID,
		CinemaName: req.CinemaName,
		MovieTime:  req.MovieTime,
	})
	if err != nil {
		return &entities.ShowTimeAddResponse{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	if len(resp) > 1 {
		return &entities.ShowTimeAddResponse{
			Result: entities.Result{
				Code:    enums.SHOW_TIME_CODE,
				Message: enums.SHOW_TIME_MESS,
			},
		}, nil
	}
	err = s.st.AddShowTime(ctx, &domain.ShowTime{
		ID:         utils.GenerateUniqueKey(),
		TicketID:   req.TicketID,
		CinemaName: req.CinemaName,
		MovieTime:  req.MovieTime,
		CreatedAt:  utils.GenerateTimestamp(),
		UpdatedAt:  utils.GenerateTimestamp(),
	})

	if err != nil {
		return &entities.ShowTimeAddResponse{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	return &entities.ShowTimeAddResponse{
		Result: entities.Result{
			Code:    enums.DB_ERR_CODE,
			Message: enums.DB_ERR_MESS,
		},
	}, nil
}

func (s *UseCaseShowTime) DeleteShowTime(ctx context.Context, req *entities.ShowTimeDelete) (*entities.ShowTimeDeleteResponse, error) {
	err := s.st.DeleteShowTimeByTicketId(ctx, &domain.ShowTimeDelete{
		ID:        req.ID,
		TicketID:  req.TicketID,
		MovieTime: req.MovieTime,
	})
	if err != nil {
		return &entities.ShowTimeDeleteResponse{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	return &entities.ShowTimeDeleteResponse{
		Result: entities.Result{
			Code:    enums.DB_ERR_CODE,
			Message: enums.DB_ERR_MESS,
		},
	}, nil
}
