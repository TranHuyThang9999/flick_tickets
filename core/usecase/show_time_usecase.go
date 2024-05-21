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
	"time"
)

type UseCaseShowTime struct {
	st     domain.RepositoryShowTime
	cinema domain.RepositoryCinemas
	trans  domain.RepositoryTransaction
	cart   domain.RepositoryCarts
}

func NewUseCaseShowTime(
	st domain.RepositoryShowTime,
	cinema domain.RepositoryCinemas,
	trans domain.RepositoryTransaction,
	cart domain.RepositoryCarts,
) *UseCaseShowTime {
	return &UseCaseShowTime{
		st:     st,
		cinema: cinema,
		trans:  trans,
		cart:   cart,
	}
}
func (s *UseCaseShowTime) AddShowTime(ctx context.Context, req *entities.ShowTimeAddReq) (*entities.ShowTimeAddResponse, error) {

	// check show time
	listShowTimeInt, err := mapper.ParseToIntSlice(req.MovieTime)
	if err != nil {
		return &entities.ShowTimeAddResponse{
			Result: entities.Result{
				Code:    enums.CONVERT_STRING_TO_ARRAY_CODE,
				Message: enums.CONVERT_STRING_TO_ARRAY_MESS,
			},
		}, nil
	}

	// check show
	statusCheckList := mapper.HasDuplicate(listShowTimeInt)
	if statusCheckList {
		return &entities.ShowTimeAddResponse{
			Result: entities.Result{
				Code:    enums.SHOW_TIME_CODE,
				Message: enums.SHOW_TIME_MESS,
			},
		}, nil
	}

	listCinemasName := mapper.ConvertListToStringSlice(req.CinemaName)

	checkTime, err := s.st.FindDuplicateShowTimes(ctx, listShowTimeInt, listCinemasName)
	if err != nil {
		return &entities.ShowTimeAddResponse{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}

	if len(checkTime) > 0 {
		return &entities.ShowTimeAddResponse{
			Result: entities.Result{
				Code:    enums.SHOW_TIME_CODE,
				Message: enums.SHOW_TIME_MESS,
			},
		}, nil
	}
	// /add show time
	var reqListShowTime = make([]*domain.ShowTime, 0)

	for i := 0; i < len(listCinemasName); i++ {
		for j := 0; j < len(listShowTimeInt); j++ {
			reqListShowTime = append(reqListShowTime, &domain.ShowTime{
				ID:             utils.GenerateUniqueKey(),
				TicketID:       req.TicketID,
				SelectedSeat:   "",
				Quantity:       req.Quantity,
				OriginalNumber: req.Quantity,
				CinemaName:     listCinemasName[i],
				MovieTime:      listShowTimeInt[j],
				Price:          req.Price,
				Discount:       req.Discount,
				CreatedAt:      utils.GenerateTimestamp(),
				UpdatedAt:      utils.GenerateTimestamp(),
			})
		}
	}
	err = s.st.UpsertListShowTime(ctx, reqListShowTime)
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
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
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
func (s *UseCaseShowTime) GetShowTimeByTicketId(ctx context.Context, ticketId string) (*entities.ShowTimeByTicketIdresp, error) {

	number := mapper.ConvertStringToInt(ticketId) // Convert ticket ID to int
	var timeNowTypetimestamp = time.Now().Unix()  // Get the current timestamp

	// Fetch the list of showtimes by ticket ID from the database
	listShowTime, err := s.st.GetShowTimeByTicketId(ctx, int64(number))
	if err != nil {
		// Return an error if unable to access the data
		return &entities.ShowTimeByTicketIdresp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, err
	}

	// Check if the list of showtimes is empty
	if len(listShowTime) == 0 {
		return &entities.ShowTimeByTicketIdresp{
			Result: entities.Result{
				Code:    enums.DATA_EMPTY_ERR_CODE,
				Message: enums.DATA_EMPTY_ERR_MESS,
			},
		}, nil
	}

	// Fetch the list of cinemas
	listCinema, err := s.cinema.GetAllCinema(ctx)
	if err != nil {
		return &entities.ShowTimeByTicketIdresp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, err
	}

	// Check if the list of cinemas is empty
	if len(listCinema) == 0 {
		return &entities.ShowTimeByTicketIdresp{
			Result: entities.Result{
				Code:    enums.DATA_EMPTY_ERR_CODE,
				Message: enums.DATA_EMPTY_ERR_MESS,
			},
		}, nil
	}

	// Create a map of cinemas by their name
	mapCinemaByName := make(map[string]*domain.Cinemas)
	for _, cinema := range listCinema {
		mapCinemaByName[cinema.CinemaName] = cinema
	}

	// Create a list to hold the showtime details
	var listRespDetail []*entities.ShowTime
	for _, showTime := range listShowTime {
		cinema := mapCinemaByName[showTime.CinemaName]
		if cinema == nil {
			// If cinema information is not found, create a default cinema object with empty fields
			cinema = &domain.Cinemas{
				CinemaName:      "",
				Description:     "",
				Conscious:       "",
				District:        "",
				Commune:         "",
				AddressDetails:  "",
				WidthContainer:  0,
				HeightContainer: 0,
			}
		}

		// Filter showtimes to include only those with MovieTime greater than the current timestamp
		if showTime.MovieTime > int(timeNowTypetimestamp) {
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
				Discount:        showTime.Discount,
			})
		}
	}

	// Check if the filtered list of showtimes is empty
	if len(listRespDetail) == 0 {
		return &entities.ShowTimeByTicketIdresp{
			Result: entities.Result{
				Code:    enums.DATA_EMPTY_ERR_CODE,
				Message: enums.DATA_EMPTY_ERR_MESS,
			},
		}, nil
	}

	// Sort the filtered list of showtimes by MovieTime in ascending order
	sort.Slice(listRespDetail, func(i, j int) bool {
		return listRespDetail[i].MovieTime < listRespDetail[j].MovieTime
	})

	// Return the filtered and sorted list of showtimes
	return &entities.ShowTimeByTicketIdresp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Showtimes: listRespDetail,
	}, nil
}

func (s *UseCaseShowTime) GetShowTimeByTicketIdForAdmin(ctx context.Context, ticketId string) (*entities.ShowTimeByTicketIdresp, error) {

	number := mapper.ConvertStringToInt(ticketId) //ticket id
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
		cinema := mapCinemaByName[showTime.CinemaName]
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
			Discount:        showTime.Discount,
		})
	}
	log.Infof("listRespDetail", listRespDetail)

	// Sắp xếp danh sách thời gian chiếu theo thời gian của phim
	sort.Slice(listRespDetail, func(i, j int) bool {
		return int(listRespDetail[i].ID) > int(listRespDetail[j].ID)
	})
	return &entities.ShowTimeByTicketIdresp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Showtimes: listRespDetail,
	}, nil

	// Trả về danh sách thời gian chiếu đã được chế biến và không có lỗi
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
func (s *UseCaseShowTime) DeleteShowTimeById(ctx context.Context, show_time_id string) (*entities.ShowTimeDeleteByIdResp, error) {

	show_time_id_number := mapper.ConvertStringToInt(show_time_id)
	tx, err := s.trans.BeginTransaction(ctx)
	if err != nil {
		return &entities.ShowTimeDeleteByIdResp{
			Result: entities.Result{
				Code:    enums.TRANSACTION_INVALID_CODE,
				Message: enums.TRANSACTION_INVALID_MESS,
			},
		}, nil
	}
	err = s.st.DeleteShowTimeByid(ctx, tx, int64(show_time_id_number))
	if err != nil {
		return &entities.ShowTimeDeleteByIdResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	err = s.cart.DeleteCartByShowTimeId(ctx, int64(show_time_id_number))
	if err != nil {
		tx.Rollback()
		return &entities.ShowTimeDeleteByIdResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	tx.Commit()
	return &entities.ShowTimeDeleteByIdResp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		CreatedAt: utils.GenerateTimestamp(),
	}, nil
}
func (s *UseCaseShowTime) GetShowTimeById(ctx context.Context, id string) (*entities.ShowTimeFindByIdResp, error) {

	if id == "" {
		return &entities.ShowTimeFindByIdResp{
			Result: entities.Result{
				Code:    enums.INVALID_REQUEST_CODE,
				Message: enums.INVALID_REQUEST_MESS,
			},
		}, nil
	}

	showTime, err := s.st.GetShowTimeById(ctx, int64(mapper.ConvertStringToInt(id)))
	if err != nil {
		return &entities.ShowTimeFindByIdResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	return &entities.ShowTimeFindByIdResp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		ShowTime: entities.ShowTimeResp{
			ID:             showTime.ID,
			TicketID:       showTime.TicketID,
			CinemaName:     showTime.CinemaName,
			Price:          showTime.Price,
			MovieTime:      showTime.MovieTime,
			SelectedSeat:   showTime.SelectedSeat,
			Quantity:       showTime.Quantity,
			OriginalNumber: showTime.OriginalNumber,
			Discount:       showTime.Discount,
			CreatedAt:      showTime.CreatedAt,
			UpdatedAt:      showTime.UpdatedAt,
		},
	}, nil
}
func (s *UseCaseShowTime) UpdateShowTimeById(ctx context.Context, req *entities.ShowTimeUpdateByIdReq) (*entities.ShowTimeUpdateByIdResp, error) {
	// Kiểm tra nếu CinemaName hoặc MovieTime không được truyền vào, lấy từ API theo ID
	if req.CinemaName == "" || req.MovieTime == 0 {
		showTime, err := s.st.GetShowTimeById(ctx, req.ID)
		if err != nil {
			return &entities.ShowTimeUpdateByIdResp{
				Result: entities.Result{
					Code:    enums.DB_ERR_CODE,
					Message: enums.DB_ERR_MESS,
				},
			}, nil
		}
		if req.CinemaName == "" {
			req.CinemaName = showTime.CinemaName
		}
		if req.MovieTime == 0 {
			req.MovieTime = showTime.MovieTime
		}
	}
	log.Infof("data ", req.MovieTime)
	// Kiểm tra nếu không tìm thấy bản ghi để cập nhật, thêm mới
	showTimeGetCheck, err := s.st.GetTimeUseCheckAddTicket(ctx, &domain.ShowTimeCheckList{
		CinemaName: req.CinemaName,
		MovieTime:  req.MovieTime,
	})
	log.Infof("data", showTimeGetCheck)
	if err != nil {
		return &entities.ShowTimeUpdateByIdResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	if showTimeGetCheck == nil {
		err = s.st.UpdateShowTimeById(ctx, &domain.ShowTimeUpdateReq{
			ID:             req.ID,
			TicketID:       req.TicketID,
			CinemaName:     req.CinemaName,
			MovieTime:      req.MovieTime,
			Quantity:       req.Quantity,
			OriginalNumber: req.Quantity,
			Price:          req.Price,
			Discount:       req.Discount,
			UpdatedAt:      utils.GenerateTimestamp(),
		})
		if err != nil {
			return &entities.ShowTimeUpdateByIdResp{
				Result: entities.Result{
					Code:    enums.DB_ERR_CODE,
					Message: enums.DB_ERR_MESS,
				},
			}, nil
		}
		return &entities.ShowTimeUpdateByIdResp{
			Result: entities.Result{
				Code:    enums.SUCCESS_CODE,
				Message: enums.SUCCESS_MESS,
			},
		}, nil
	}

	// Kiểm tra nếu bản ghi tồn tại, nhưng không phải bản ghi cần cập nhật
	if showTimeGetCheck.ID != req.ID {
		return &entities.ShowTimeUpdateByIdResp{
			Result: entities.Result{
				Code:    enums.SHOW_TIME_CODE,
				Message: enums.SHOW_TIME_MESS,
			},
		}, nil
	}

	// Cập nhật bản ghi
	err = s.st.UpdateShowTimeById(ctx, &domain.ShowTimeUpdateReq{
		ID:             req.ID,
		TicketID:       req.TicketID,
		CinemaName:     req.CinemaName,
		MovieTime:      req.MovieTime,
		Quantity:       req.Quantity,
		OriginalNumber: req.Quantity,
		Price:          req.Price,
		Discount:       req.Discount,
		UpdatedAt:      utils.GenerateTimestamp(),
	})
	if err != nil {
		return &entities.ShowTimeUpdateByIdResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}

	return &entities.ShowTimeUpdateByIdResp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
	}, nil
}
