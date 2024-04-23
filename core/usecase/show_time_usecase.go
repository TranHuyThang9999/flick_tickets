package usecase

import (
	"context"
	"flick_tickets/common/enums"
	"flick_tickets/common/utils"
	"flick_tickets/core/domain"
	"flick_tickets/core/entities"
	"flick_tickets/core/mapper"
)

type UseCaseShowTime struct {
	st     domain.RepositoryShowTime
	cinema domain.RepositoryCinemas
}

func NewUseCaseShowTime(
	st domain.RepositoryShowTime,
	cinema domain.RepositoryCinemas,
) *UseCaseShowTime {
	return &UseCaseShowTime{
		st:     st,
		cinema: cinema,
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
		ID:             utils.GenerateUniqueKey(),
		TicketID:       req.TicketID,
		CinemaName:     req.CinemaName,
		MovieTime:      req.MovieTime,
		SelectedSeat:   req.SelectedSeat,
		OriginalNumber: req.Quantity,
		Quantity:       req.Quantity,
		Price:          req.Price,
		CreatedAt:      utils.GenerateTimestamp(),
		UpdatedAt:      utils.GenerateTimestamp(),
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
func (s *UseCaseShowTime) GetShowTimeByTicketId(ctx context.Context, id string) (*entities.ShowTimeByTicketIdresp, error) {

	number := mapper.ConvertStringToInt(id)
	var listRespDetail []*entities.ShowTime
	var listCinema []*domain.Cinemas
	var mapCinemaByName = make(map[string]*domain.Cinemas)

	listShowTime, err := s.st.GetShowTimeByTicketId(ctx, int64(number))
	if err != nil {
		return &entities.ShowTimeByTicketIdresp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	if len(listShowTime) == 0 {
		return &entities.ShowTimeByTicketIdresp{
			Result: entities.Result{
				Code:    enums.DATA_EMPTY_ERR_CODE,
				Message: enums.DATA_EMPTY_ERR_MESS,
			},
		}, nil
	}
	listCinema, err = s.cinema.GetAllCinema(ctx)
	if err != nil {
		return &entities.ShowTimeByTicketIdresp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	if len(listCinema) == 0 {
		return &entities.ShowTimeByTicketIdresp{
			Result: entities.Result{
				Code:    enums.DATA_EMPTY_ERR_CODE,
				Message: enums.DATA_EMPTY_ERR_MESS,
			},
		}, nil
	}

	// Build mapCinema from the list of cinemas
	for _, cinema := range listCinema {
		mapCinemaByName[cinema.CinemaName] = cinema
	}

	for i := 0; i < len(listShowTime); i++ {
		cinema := mapCinemaByName[listShowTime[i].CinemaName]
		if cinema == nil {
			return &entities.ShowTimeByTicketIdresp{
				Result: entities.Result{
					Code:    enums.DATA_EMPTY_ERR_CODE,
					Message: enums.DATA_EMPTY_ERR_MESS,
				},
			}, nil
		}
		listRespDetail = append(listRespDetail, &entities.ShowTime{
			ID:              listShowTime[i].ID,
			TicketID:        listShowTime[i].TicketID,
			CinemaName:      listShowTime[i].CinemaName,
			MovieTime:       listShowTime[i].MovieTime,
			Description:     cinema.Description,
			Conscious:       cinema.Conscious,
			District:        cinema.District,
			Commune:         cinema.Commune,
			AddressDetails:  cinema.AddressDetails,
			WidthContainer:  cinema.WidthContainer,
			HeightContainer: cinema.HeightContainer,
			SelectedSeat:    listShowTime[i].SelectedSeat,
			Quantity:        listShowTime[i].Quantity,
			OriginalNumber:  listShowTime[i].OriginalNumber,
			Price:           listShowTime[i].Price,
		})
	}

	return &entities.ShowTimeByTicketIdresp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Showtimes: listRespDetail,
	}, nil
}
func (s *UseCaseShowTime) DetailShowTime(ctx context.Context, id string) (*entities.ShowTimeDetail, error) {

	showTimeId := mapper.ConvertStringToInt(id)

	showTime, err := s.st.GetInformationShowTimeForTicketByTicketId(ctx, int64(showTimeId))
	if err != nil {
		return &entities.ShowTimeDetail{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	return &entities.ShowTimeDetail{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		ShowTime: showTime,
	}, nil
}
