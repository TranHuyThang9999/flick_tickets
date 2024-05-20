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

type UseCaseTicker struct {
	ticket   domain.RepositoryTickets
	trans    domain.RepositoryTransaction
	file     domain.RepositoryFileStorages
	menory   cache.RepositoryCache
	showTime domain.RepositoryShowTime
	cinema   domain.RepositoryCinemas
}

func NewUsecaseTicker(
	ticket domain.RepositoryTickets,
	trans domain.RepositoryTransaction,
	file domain.RepositoryFileStorages,
	menory cache.RepositoryCache,
	showTime domain.RepositoryShowTime,
	cinema domain.RepositoryCinemas,

) *UseCaseTicker {
	return &UseCaseTicker{
		ticket:   ticket,
		trans:    trans,
		file:     file,
		menory:   menory,
		showTime: showTime,
		cinema:   cinema,
	}
}
func (c *UseCaseTicker) AddTicket(ctx context.Context, req *entities.TicketReqUpload) (*entities.TicketRespUpload, error) {

	var idTicket int64 = utils.GenerateUniqueKey()
	var listImage = make([]*domain.FileStorages, 0)

	tx, err := c.trans.BeginTransaction(ctx)
	if err != nil {
		return &entities.TicketRespUpload{
			Result: entities.Result{
				Code:    enums.TRANSACTION_INVALID_CODE,
				Message: enums.TRANSACTION_INVALID_MESS,
			},
		}, nil
	}
	ticketAdd := &domain.Tickets{
		ID:            idTicket,
		Name:          req.Name,
		Price:         req.Price,
		Description:   req.Description,
		Sale:          req.Sale,
		ReleaseDate:   req.ReleaseDate,
		Status:        req.Status,
		MovieDuration: req.MovieDuration,
		AgeLimit:      req.AgeLimit,
		Director:      req.Director,
		Actor:         req.Actor,
		Producer:      req.Producer,
		MovieType:     req.MovieType,
		CreatedAt:     utils.GenerateTimestamp(),
		UpdatedAt:     utils.GenerateTimestamp(),
	}

	err = c.ticket.AddTicket(ctx, tx, ticketAdd)
	if err != nil {
		tx.Rollback()
		return &entities.TicketRespUpload{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	//check show time
	listShowTimeInt, err := mapper.ParseToIntSlice(req.MovieTime)
	// log.Infof("time : ", listShowTimeInt)
	if err != nil {
		tx.Rollback()
		return &entities.TicketRespUpload{
			Result: entities.Result{
				Code:    enums.CONVERT_STRING_TO_ARRAY_CODE,
				Message: enums.CONVERT_STRING_TO_ARRAY_MESS,
			},
		}, nil
	}

	//check show
	statusCheckList := mapper.HasDuplicate(listShowTimeInt)
	if statusCheckList {
		return &entities.TicketRespUpload{
			Result: entities.Result{
				Code:    enums.SHOW_TIME_CODE,
				Message: enums.SHOW_TIME_MESS,
			},
		}, nil
	}

	listCinemasName := mapper.ConvertListToStringSlice(req.CinemaName)

	checkTime, err := c.showTime.FindDuplicateShowTimes(ctx, listShowTimeInt, listCinemasName)
	if err != nil {
		tx.Rollback()
		return &entities.TicketRespUpload{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	if len(checkTime) > 1 {
		return &entities.TicketRespUpload{
			Result: entities.Result{
				Code:    enums.SHOW_TIME_CODE,
				Message: enums.SHOW_TIME_MESS,
			},
		}, nil
	}
	///add show time
	var reqListShowTime = make([]*domain.ShowTime, 0)

	for i := 0; i < len(listCinemasName); i++ {
		for j := 0; j < len(listShowTimeInt); j++ {
			reqListShowTime = append(reqListShowTime, &domain.ShowTime{
				ID:             utils.GenerateUniqueKey(),
				TicketID:       idTicket,
				SelectedSeat:   "",
				Quantity:       req.Quantity,
				OriginalNumber: req.Quantity,
				CinemaName:     listCinemasName[i],
				MovieTime:      listShowTimeInt[j],
				Price:          req.Price,
				Discount:       req.Sale,
				CreatedAt:      utils.GenerateTimestamp(),
				UpdatedAt:      utils.GenerateTimestamp(),
			})
		}
	}

	err = c.showTime.AddListShowTime(ctx, tx, reqListShowTime)
	if err != nil {
		tx.Rollback()
		return &entities.TicketRespUpload{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}

	respFile, err := utils.SetListFile(ctx, req.File)
	if err != nil {
		return &entities.TicketRespUpload{
			Result: entities.Result{
				Code:    enums.UPLOAD_FILE_ERR_CODE,
				Message: enums.UPLOAD_FILE_ERR_MESS,
			},
		}, nil
	}
	if len(respFile) > 0 {
		for _, file := range respFile {
			// err = c.file.AddInformationFileStorages(ctx, tx, &domain.FileStorages{
			// 	ID:        utils.GenerateUniqueKey(),
			// 	URL:       file.URL,
			// 	TicketID:  idTicket,
			// 	CreatedAt: utils.GenerateTimestamp(),
			// })
			// if err != nil {
			// 	tx.Rollback()
			// 	return &entities.TicketRespUpload{
			// 		Result: entities.Result{
			// 			Code:    enums.DB_ERR_CODE,
			// 			Message: enums.DB_ERR_MESS,
			// 		},
			// 	}, nil
			// }
			listImage = append(listImage, &domain.FileStorages{
				ID:        utils.GenerateUniqueKey(),
				URL:       file.URL,
				TicketID:  idTicket,
				CreatedAt: utils.GenerateTimestamp(),
			})
		}

	}
	if len(listImage) > 0 {
		err = c.file.AddListInformationFileStorages(ctx, tx, listImage)
		if err != nil {
			tx.Rollback()
			return &entities.TicketRespUpload{
				Result: entities.Result{
					Code:    enums.DB_ERR_CODE,
					Message: enums.DB_ERR_MESS,
				},
			}, nil
		}
	}

	tx.Commit()
	return &entities.TicketRespUpload{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
	}, nil
}

func (c *UseCaseTicker) GetTicketById(ctx context.Context, id string) (*entities.TicketRespgetById, error) {
	// Chuyển đổi id từ chuỗi sang số nguyên
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		return &entities.TicketRespgetById{
			Result: entities.Result{
				Code:    enums.CONVERT_TO_NUMBER_CODE,
				Message: enums.CONVERT_TO_NUMBER_MESS,
			},
		}, nil
	}

	// Lấy vé từ cơ sở dữ liệu
	ticket, err := c.ticket.GetTicketById(ctx, int64(idNumber))
	if err != nil {
		return &entities.TicketRespgetById{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}

	// Trả về thông tin vé và kết quả thành công
	return &entities.TicketRespgetById{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Ticket:    ticket,
		CreatedAt: utils.GenerateTimestamp(),
	}, nil
}

func (c *UseCaseTicker) GetAllTickets(ctx context.Context, req *domain.TicketreqFindByForm) (*entities.TicketRespGetAll, error) {

	// var listTicketResp = make([]*entities.Tickets, 0)
	// var listShowTime []*domain.ShowTime
	// var mapShowTimes = make(map[int64][]*domain.ShowTime)
	// var mapListFile = make(map[int64][]*domain.FileStorages)
	// var listFile []*domain.FileStorages

	// listTicket, err := c.ticket.GetAllTickets(ctx, req) //get all
	// if err != nil {
	// 	return &entities.TicketRespGetAll{
	// 		Result: entities.Result{
	// 			Code:    enums.DB_ERR_CODE,
	// 			Message: enums.DB_ERR_MESS,
	// 		},
	// 	}, err
	// }
	// if len(listTicket) == 0 {
	// 	return &entities.TicketRespGetAll{
	// 		Result: entities.Result{
	// 			Code:    enums.DATA_EMPTY_ERR_CODE,
	// 			Message: enums.DATA_EMPTY_ERR_MESS,
	// 		},
	// 	}, nil
	// }
	// listShowTime, err = c.showTime.GetAll(ctx)
	// if err != nil {
	// 	return &entities.TicketRespGetAll{
	// 		Result: entities.Result{
	// 			Code:    enums.DB_ERR_CODE,
	// 			Message: enums.DB_ERR_MESS,
	// 		},
	// 	}, err
	// }
	// listFile, err = c.file.GetAll(ctx)
	// if err != nil {
	// 	return &entities.TicketRespGetAll{
	// 		Result: entities.Result{
	// 			Code:    enums.DB_ERR_CODE,
	// 			Message: enums.DB_ERR_MESS,
	// 		},
	// 	}, err
	// }
	// // Xây dựng mapShowTimes từ danh sách showtime
	// for _, showTime := range listShowTime {
	// 	ticketID := showTime.TicketID
	// 	mapShowTimes[ticketID] = append(mapShowTimes[ticketID], showTime)
	// }
	// for _, v := range listFile {
	// 	ticketID := v.TicketID
	// 	mapListFile[ticketID] = append(mapListFile[ticketID], v)
	// }
	// for _, ticket := range listTicket {
	// 	showTimes := mapShowTimes[ticket.ID]
	// 	listFile := mapListFile[ticket.ID]
	// 	listTicketResp = append(listTicketResp, &entities.Tickets{
	// 		Ticket:    ticket,
	// 		ShowTimes: showTimes,
	// 		ListUrl:   listFile,
	// 	})
	// }

	// return &entities.TicketRespGetAll{
	// 	Result: entities.Result{
	// 		Code:    enums.SUCCESS_CODE,
	// 		Message: enums.SUCCESS_MESS,
	// 	},
	// 	Tickets: listTicketResp,
	// }, nil
	//
	if req.MovieTheaterName == "" {
		listTicket, err := c.ticket.GetAllTickets(ctx, req) //get all
		if err != nil {
			return &entities.TicketRespGetAll{
				Result: entities.Result{
					Code:    enums.DB_ERR_CODE,
					Message: enums.DB_ERR_MESS,
				},
			}, err
		}
		if len(listTicket) == 0 {
			return &entities.TicketRespGetAll{
				Result: entities.Result{
					Code:    enums.DATA_EMPTY_ERR_CODE,
					Message: enums.DATA_EMPTY_ERR_MESS,
				},
			}, nil
		}

		return &entities.TicketRespGetAll{
			Result: entities.Result{
				Code:    enums.SUCCESS_CODE,
				Message: enums.SUCCESS_MESS,
			},
			ListTickets: listTicket,
		}, nil
	} else {
		var listTicketId = make([]int64, 0)
		listShowTime, err := c.showTime.GetShowTimeByNameCinema(ctx, req.MovieTheaterName)
		if err != nil {
			return &entities.TicketRespGetAll{
				Result: entities.Result{
					Code:    enums.DB_ERR_CODE,
					Message: enums.DB_ERR_MESS,
				},
			}, err
		}
		for i := 0; i < len(listShowTime); i++ {
			listTicketId = append(listTicketId, listShowTime[i].TicketID)
		}
		ticket, err := c.ticket.GetlistTicketByListTicketId(ctx, listTicketId)
		if err != nil {
			return &entities.TicketRespGetAll{
				Result: entities.Result{
					Code:    enums.DB_ERR_CODE,
					Message: enums.DB_ERR_MESS,
				},
			}, err
		}
		log.Infof("list ", listTicketId)
		return &entities.TicketRespGetAll{
			Result: entities.Result{
				Code:    enums.SUCCESS_CODE,
				Message: enums.SUCCESS_MESS,
			},
			ListTickets: ticket,
		}, nil
	}

}

func (c *UseCaseTicker) DeleteTicketsById(ctx context.Context, ticketId string) (*entities.TicketRespDeleteById, error) {

	ticketIdNumber := mapper.ConvertStringToInt(ticketId)
	tx, err := c.trans.BeginTransaction(ctx)
	if err != nil {
		return &entities.TicketRespDeleteById{
			Result: entities.Result{
				Code:    enums.TRANSACTION_INVALID_CODE,
				Message: enums.TRANSACTION_INVALID_MESS,
			},
		}, nil
	}

	ticket, err := c.ticket.GetTicketById(ctx, int64(ticketIdNumber))
	if err != nil {
		return &entities.TicketRespDeleteById{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	err = c.ticket.DeleteTicketsById(ctx, tx, ticket.ID)
	if err != nil {
		return &entities.TicketRespDeleteById{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	err = c.file.DeleteFileByAnyIdObject(ctx, tx, ticket.ID)
	if err != nil {
		tx.Rollback()
		return &entities.TicketRespDeleteById{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	err = c.showTime.DeleteShowTimesByTicketId(ctx, tx, ticket.ID)
	if err != nil {
		tx.Rollback()
		return &entities.TicketRespDeleteById{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	tx.Commit()
	return &entities.TicketRespDeleteById{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
	}, nil
}
func (c *UseCaseTicker) UpdateTicketById(ctx context.Context, req *entities.TicketReqUpdateById) (*entities.TicketRespUpdateById, error) {

	err := c.ticket.UpdateTicketById(ctx, &domain.TicketReqUpdateById{
		ID:            req.ID,
		Name:          req.Name,
		Price:         req.Price,
		Description:   req.Description,
		Sale:          req.Sale,
		ReleaseDate:   req.ReleaseDate,
		Status:        req.Status,
		MovieDuration: req.MovieDuration,
		AgeLimit:      req.AgeLimit,
		Director:      req.Director,
		Actor:         req.Actor,
		Producer:      req.Producer,
		MovieType:     req.MovieType,
		UpdatedAt:     utils.GenerateTimestamp(),
	})
	log.Infof("req : ", req)
	if err != nil {
		return &entities.TicketRespUpdateById{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}

	return &entities.TicketRespUpdateById{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
	}, nil
}

func (c *UseCaseTicker) GetAllTicketsAttachSale(ctx context.Context, status string) (*entities.TicketGetAllByStatusResp, error) { // ko dung

	listTicket, err := c.ticket.GetListTicketWithSatus(ctx, mapper.ConvertStringToInt(status))
	if err != nil {
		return &entities.TicketGetAllByStatusResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	return &entities.TicketGetAllByStatusResp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		ListTickets: listTicket,
	}, nil
}
func (c *UseCaseTicker) GetAllTicketsByFilmName(ctx context.Context, req *entities.TicketFindByMovieNameReq) (
	*entities.TicketFindByMovieNameResp, error) {
	tickets, err := c.ticket.GetAllTicketsByFilmName(ctx, req.MovieName)
	if err != nil {
		return &entities.TicketFindByMovieNameResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	return &entities.TicketFindByMovieNameResp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Tickets: tickets,
	}, nil
}
