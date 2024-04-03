package usecase

import (
	"context"
	"flick_tickets/common/enums"
	"flick_tickets/common/log"
	"flick_tickets/common/utils"
	"flick_tickets/core/domain"
	"flick_tickets/core/entities"
)

type UseCaseTicker struct {
	ticket domain.RepositoryTickets
	trans  domain.RepositoryTransaction
	file   domain.RepositoryFileStorages
}

func NewUsecaseTicker(
	ticket domain.RepositoryTickets,
	trans domain.RepositoryTransaction,
	file domain.RepositoryFileStorages,

) *UseCaseTicker {
	return &UseCaseTicker{
		ticket: ticket,
		trans:  trans,
		file:   file,
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

	err = c.ticket.AddTicket(ctx, tx, &domain.Tickets{
		ID:          idTicket,
		UserId:      req.UserId,
		Name:        req.Name,
		Price:       req.Price,
		MaxTicket:   int64(req.Quantity),
		Quantity:    req.Quantity,
		Description: req.Description,
		Sale:        req.Sale,
		Showtime:    req.Showtime,
		ReleaseDate: req.ReleaseDate,
		CreatedAt:   utils.GenerateTimestamp(),
		UpdatedAt:   utils.GenerateTimestamp(),
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
	respFile, err := utils.SetListFile(ctx, req.File)
	if err != nil {
		return &entities.TicketRespUpload{
			Result: entities.Result{
				Code:    enums.UPLOAD_FILE_ERR_CODE,
				Message: enums.UPLOAD_FILE_ERR_MESS,
			},
		}, nil
	}
	log.Infof("file", respFile, len(respFile))
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
