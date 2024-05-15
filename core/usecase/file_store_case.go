package usecase

import (
	"context"
	"flick_tickets/common/enums"
	"flick_tickets/common/utils"
	"flick_tickets/core/domain"
	"flick_tickets/core/entities"
	"flick_tickets/core/mapper"
	"strconv"
)

type UseCaseFileStore struct {
	file  domain.RepositoryFileStorages
	trans domain.RepositoryTransaction
}

func NewUseCaseFile(
	file domain.RepositoryFileStorages,
	trans domain.RepositoryTransaction,
) *UseCaseFileStore {
	return &UseCaseFileStore{
		file:  file,
		trans: trans,
	}
}
func (u *UseCaseFileStore) GetListFileByObjectId(ctx context.Context, id string) (*entities.ResponseGetListFileByObjetId, error) {

	number, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return &entities.ResponseGetListFileByObjetId{
			Result: entities.Result{
				Code:    enums.CONVERT_TO_NUMBER_CODE,
				Message: enums.CONVERT_TO_NUMBER_MESS,
			},
		}, nil
	}

	listFile, err := u.file.GetListFileById(ctx, number)
	if err != nil {
		return &entities.ResponseGetListFileByObjetId{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	if len(listFile) == 0 {
		return &entities.ResponseGetListFileByObjetId{
			Result: entities.Result{
				Code:    enums.DATA_EMPTY_ERR_CODE,
				Message: enums.DATA_EMPTY_ERR_MESS,
			},
		}, nil
	}
	return &entities.ResponseGetListFileByObjetId{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Files: listFile,
	}, nil

}
func (u *UseCaseFileStore) UploadFileByTicketId(ctx context.Context, req *entities.UpSertFileDescriptReq) (*entities.UpSertFileDescriptResp, error) {

	var listFile = make([]*domain.FileStorages, 0)
	tx, err := u.trans.BeginTransaction(ctx)
	if err != nil {
		return &entities.UpSertFileDescriptResp{
			Result: entities.Result{
				Code:    enums.TRANSACTION_INVALID_CODE,
				Message: enums.TRANSACTION_INVALID_MESS,
			},
		}, nil
	}
	respFile, err := utils.SetListFile(ctx, req.File)
	if err != nil {
		return &entities.UpSertFileDescriptResp{
			Result: entities.Result{
				Code:    enums.UPLOAD_FILE_ERR_CODE,
				Message: enums.UPLOAD_FILE_ERR_MESS,
			},
		}, nil
	}
	if len(respFile) > 0 {
		for _, file := range respFile {

			listFile = append(listFile, &domain.FileStorages{
				ID:        utils.GenerateUniqueKey(),
				URL:       file.URL,
				TicketID:  req.TicketId,
				CreatedAt: utils.GenerateTimestamp(),
			})
		}

	}
	if len(listFile) > 0 {
		err = u.file.AddListInformationFileStorages(ctx, tx, listFile)
		if err != nil {
			tx.Rollback()
			return &entities.UpSertFileDescriptResp{
				Result: entities.Result{
					Code:    enums.DB_ERR_CODE,
					Message: enums.DB_ERR_MESS,
				},
			}, nil
		}
	}

	tx.Commit()
	return &entities.UpSertFileDescriptResp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
	}, nil
}
func (u *UseCaseFileStore) DeleteFileById(ctx context.Context, fileId string) (*entities.DeleteFileByIdResp, error) {

	err := u.file.DeleteFileByIdNotTransaction(ctx, int64(mapper.ConvertStringToInt(fileId)))
	if err != nil {
		return &entities.DeleteFileByIdResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	return &entities.DeleteFileByIdResp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
	}, nil
}
