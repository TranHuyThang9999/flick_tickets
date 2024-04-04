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
}

func NewUsecaseTicker(
	ticket domain.RepositoryTickets,
	trans domain.RepositoryTransaction,
	file domain.RepositoryFileStorages,
	menory cache.RepositoryCache,
	showTime domain.RepositoryShowTime,

) *UseCaseTicker {
	return &UseCaseTicker{
		ticket:   ticket,
		trans:    trans,
		file:     file,
		menory:   menory,
		showTime: showTime,
	}
}
func (c *UseCaseTicker) AddTicket(ctx context.Context, req *entities.TicketReqUpload) (*entities.TicketRespUpload, error) {

	tx, err := c.trans.BeginTransaction(ctx)
	if err != nil {
		return &entities.TicketRespUpload{
			Result: entities.Result{
				Code:    enums.TRANSACTION_INVALID_CODE,
				Message: enums.TRANSACTION_INVALID_MESS,
			},
		}, nil
	}

	var idTicket int64 = utils.GenerateUniqueKey()
	ticketAdd := &domain.Tickets{
		ID:          idTicket,
		Name:        req.Name,
		Price:       req.Price,
		MaxTicket:   int64(req.Quantity),
		Quantity:    req.Quantity,
		Description: req.Description,
		Sale:        req.Sale,
		ReleaseDate: req.ReleaseDate,
		Status:      req.Status,
		CreatedAt:   utils.GenerateTimestamp(),
		UpdatedAt:   utils.GenerateTimestamp(),
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

	showTimes, err := c.showTime.GetTimeUseCheckAddTicket(ctx, &domain.ShowTimeCheckList{
		CinemaName: req.CinemaName,
		MovieTime:  req.MovieTime,
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
	log.Infof("len", len(showTimes))
	if len(showTimes) >= 1 {
		log.Info("*************** 0")
		return &entities.TicketRespUpload{
			Result: entities.Result{
				Code:    enums.SHOW_TIME_CODE,
				Message: enums.SHOW_TIME_MESS,
			},
		}, nil
	}

	//check show time
	listShowTimeInt, err := mapper.ParseToIntSlice(req.MovieTime)
	if err != nil {
		tx.Rollback()
		return &entities.TicketRespUpload{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}

	//check show
	statusCheckList := mapper.HasDuplicate(listShowTimeInt)
	log.Infof("$$$ ", statusCheckList)
	if statusCheckList {
		log.Info("*************** 1")
		return &entities.TicketRespUpload{
			Result: entities.Result{
				Code:    enums.SHOW_TIME_CODE,
				Message: enums.SHOW_TIME_MESS,
			},
		}, nil
	}

	// add show time
	err = c.showTime.AddShowTime(ctx, &domain.ShowTime{
		ID:         utils.GenerateUniqueKey(),
		TicketID:   idTicket,
		CinemaName: req.CinemaName,
		MovieTime:  mapper.ConvertIntArrayToString(listShowTimeInt), //array
	})
	if err != nil {
		log.Info("*************** 2")
		tx.Rollback()
		return &entities.TicketRespUpload{
			Result: entities.Result{
				Code:    enums.SHOW_TIME_CODE,
				Message: enums.SHOW_TIME_MESS,
			},
		}, nil
	}
	//set cache
	err = c.menory.SetObjectById(ctx, strconv.FormatInt(idTicket, 10), ticketAdd)
	log.Infof("error cache", err)
	if err != nil {
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

	tickets, err := c.ticket.GetAllTickets(ctx, req)
	if err != nil {
		return &entities.TicketRespGetAll{
			Result: entities.Result{
				Code:    enums.CACHE_ERR_CODE,
				Message: enums.CACHE_ERR_MESS,
			},
		}, nil
	}
	return &entities.TicketRespGetAll{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Tickets: tickets,
	}, nil
}
