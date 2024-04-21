package usecase

import (
	"context"
	"encoding/json"
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
	log.Infof("list time : ", listCinemasName)

	// listCinema, err := c.cinema.GetAllCinemaByName(ctx, req.Name)
	// if err != nil {
	// 	tx.Rollback()
	// 	return &entities.TicketRespUpload{
	// 		Result: entities.Result{
	// 			Code:    enums.DB_ERR_CODE,
	// 			Message: enums.DB_ERR_MESS,
	// 		},
	// 	}, nil
	// }
	//

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

	// for i := 0; i < len(listShowTimeInt); i++ {
	// 	reqListShowTime = append(reqListShowTime, &domain.ShowTime{
	// 		ID:         utils.GenerateUniqueKey(),
	// 		TicketID:   idTicket,
	// 		CinemaName: req.CinemaName,
	// 		MovieTime:  listShowTimeInt[i],
	// 		CreatedAt:  utils.GenerateTimestamp(),
	// 		UpdatedAt:  utils.GenerateTimestamp(),
	// 	})
	// }

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

	//set cache
	err = c.menory.SetObjectById(ctx, strconv.FormatInt(idTicket, 10), ticketAdd)
	if err != nil {
		log.Infof("error cache", err)
		tx.Rollback()
		return &entities.TicketRespUpload{
			Result: entities.Result{
				Code:    enums.CACHE_ERR_CODE,
				Message: enums.CACHE_ERR_MESS,
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
			err = c.file.AddInformationFileStorages(ctx, tx, &domain.FileStorages{
				ID:        utils.GenerateUniqueKey(),
				URL:       file.URL,
				TicketID:  idTicket,
				CreatedAt: utils.GenerateTimestamp(),
			})
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

	// Kiểm tra xem vé có tồn tại trong bộ nhớ cache không
	exists, err := c.menory.KeyExists(ctx, id)
	if err != nil {
		return &entities.TicketRespgetById{
			Result: entities.Result{
				Code:    enums.CACHE_ERR_CODE,
				Message: enums.CACHE_ERR_MESS,
			},
		}, nil
	}

	if !exists {
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

		// Lưu vé vào cache
		err = c.menory.SetObjectById(ctx, id, ticket)
		if err != nil {
			return &entities.TicketRespgetById{
				Result: entities.Result{
					Code:    enums.CACHE_ERR_CODE,
					Message: enums.CACHE_ERR_MESS,
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

	// Vé đã tồn tại trong cache, lấy thông tin vé từ cache trực tiếp
	dataString, err := c.menory.GetObjectById(ctx, id)
	if err != nil {
		return &entities.TicketRespgetById{
			Result: entities.Result{
				Code:    enums.CACHE_ERR_CODE,
				Message: enums.CACHE_ERR_MESS,
			},
		}, nil
	}

	// Chuyển đổi dữ liệu từ chuỗi JSON sang kiểu domain.Tickets
	var ticketUseCache *domain.Tickets
	err = json.Unmarshal([]byte(dataString), &ticketUseCache)
	if err != nil {
		return &entities.TicketRespgetById{
			Result: entities.Result{
				Code:    enums.ERROR_CONVERT_JSON_CODE,
				Message: enums.ERROR_CONVERT_JSON_MESS,
			},
		}, nil
	}

	// Trả về thông tin vé và kết quả thành công
	return &entities.TicketRespgetById{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Ticket:    ticketUseCache,
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
}

// func (c *UseCaseTicker) UpdateSizeRoom(ctx context.Context, req *entities.TicketReqUpdateSizeRoom) (*entities.TicketRespUpdateSizeRoom, error) {

// 	log.Infof("req update : ", req)

// 	err := c.ticket.UpdateSizeRoom(ctx, req.TicketId, req.WidthContainer, req.HeightContainer)
// 	if err != nil {
// 		return &entities.TicketRespUpdateSizeRoom{
// 			Result: entities.Result{
// 				Code:    enums.DB_ERR_CODE,
// 				Message: enums.DB_ERR_MESS,
// 			},
// 		}, err
// 	}
// 	return &entities.TicketRespUpdateSizeRoom{
// 		Result: entities.Result{
// 			Code:    enums.SUCCESS_CODE,
// 			Message: enums.SUCCESS_MESS,
// 		},
// 	}, nil
// }
