package usecase

import (
	"context"
	"flick_tickets/common/enums"
	"flick_tickets/common/log"
	"flick_tickets/common/utils"
	"flick_tickets/core/domain"
	"flick_tickets/core/entities"
	"flick_tickets/core/mapper"
	"sort"
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

	number := mapper.ConvertStringToInt(id) //ticket id

	// Lấy danh sách thời gian chiếu từ cơ sở dữ liệu
	listShowTime, err := s.st.GetShowTimeByTicketId(ctx, int64(number))
	if err != nil {
		// Trả về lỗi nếu không thể truy cập vào dữ liệu
		return &entities.ShowTimeByTicketIdresp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, err
	}
	log.Infof("data : ", len(listShowTime))
	// Kiểm tra xem danh sách thời gian chiếu có rỗng không
	if len(listShowTime) == 0 {
		return &entities.ShowTimeByTicketIdresp{
			Result: entities.Result{
				Code:    enums.DATA_EMPTY_ERR_CODE,
				Message: enums.DATA_EMPTY_ERR_MESS,
			},
		}, nil
	}

	// Lấy danh sách rạp chiếu
	listCinema, err := s.cinema.GetAllCinema(ctx)
	if err != nil {
		return &entities.ShowTimeByTicketIdresp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, err
	}
	log.Infof("cinema ", listCinema)
	// Kiểm tra xem danh sách rạp chiếu có rỗng không
	if len(listCinema) == 0 {
		return &entities.ShowTimeByTicketIdresp{
			Result: entities.Result{
				Code:    enums.DATA_EMPTY_ERR_CODE,
				Message: enums.DATA_EMPTY_ERR_MESS,
			},
		}, nil
	}

	// Tạo bản đồ các rạp chiếu từ danh sách rạp chiếu
	mapCinemaByName := make(map[string]*domain.Cinemas)
	for _, cinema := range listCinema {
		mapCinemaByName[cinema.CinemaName] = cinema
	}

	// Tạo danh sách chi tiết thời gian chiếu
	var listRespDetail []*entities.ShowTime
	for _, showTime := range listShowTime {
		log.Infof("daTa : ", showTime.CinemaName)
		cinema := mapCinemaByName[showTime.CinemaName]
		log.Infof("data  c", cinema)
		if cinema == nil {
			// Nếu không tìm thấy thông tin về rạp chiếu, gán các trường thông tin về rạp chiếu bằng chuỗi rỗng
			cinema = &domain.Cinemas{
				// Gán các trường thông tin về rạp chiếu bằng chuỗi rỗng
				CinemaName:      "",
				Description:     "",
				Conscious:       "",
				District:        "",
				Commune:         "",
				AddressDetails:  "",
				WidthContainer:  0, // hoặc giá trị mặc định khác nếu thích
				HeightContainer: 0, // hoặc giá trị mặc định khác nếu thích
			}
		}

		// Thêm chi tiết thời gian chiếu vào danh sách
		listRespDetail = append(listRespDetail, &entities.ShowTime{
			ID:              showTime.ID,
			TicketID:        showTime.TicketID,
			CinemaName:      showTime.CinemaName,
			MovieTime:       showTime.MovieTime,
			Description:     cinema.Description,
			Conscious:       cinema.Conscious,
			District:        cinema.District,
			Commune:         cinema.Commune,
			AddressDetails:  cinema.AddressDetails,
			WidthContainer:  cinema.WidthContainer,
			HeightContainer: cinema.HeightContainer,
			SelectedSeat:    showTime.SelectedSeat,
			Quantity:        showTime.Quantity,
			OriginalNumber:  showTime.OriginalNumber,
			Price:           showTime.Price,
		})
	}
	log.Infof("listRespDetail", listRespDetail)
	// Sắp xếp danh sách thời gian chiếu theo thời gian của phim
	sort.Slice(listRespDetail, func(i, j int) bool {
		return int(listRespDetail[i].ID) > int(listRespDetail[j].ID)
	})

	// Trả về danh sách thời gian chiếu đã được chế biến và không có lỗi
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
