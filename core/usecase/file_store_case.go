package usecase

import (
	"context"
	"flick_tickets/common/enums"
	"flick_tickets/core/domain"
	"flick_tickets/core/entities"
	"strconv"
)

type UseCaseFileStore struct {
	file domain.RepositoryFileStorages
}

func NewUseCaseFile(
	file domain.RepositoryFileStorages,
) *UseCaseFileStore {
	return &UseCaseFileStore{
		file: file,
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
	return &entities.ResponseGetListFileByObjetId{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Files: listFile,
	}, nil

}
