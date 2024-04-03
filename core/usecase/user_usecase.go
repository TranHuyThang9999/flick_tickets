package usecase

import (
	"context"
	"flick_tickets/common/enums"
	"flick_tickets/common/utils"
	"flick_tickets/core/domain"
	"flick_tickets/core/entities"
)

type UseCaseUser struct {
	user  domain.RepositoryUser
	file  domain.RepositoryFileStorages
	trans domain.RepositoryTransaction
}

func NewUseCaseUser(
	user domain.RepositoryUser,
	file domain.RepositoryFileStorages,
	trans domain.RepositoryTransaction,
) *UseCaseUser {
	return &UseCaseUser{
		user:  user,
		file:  file,
		trans: trans,
	}
}

func (u *UseCaseUser) AddUserd(ctx context.Context, user *entities.Users) (*entities.UserResp, error) {

	var idUser int64 = utils.GenerateUniqueKey()

	resp, err := u.GetAllUserById(ctx, &entities.UserRequestFindByForm{
		UserName: user.UserName,
	})

	if err != nil {
		return &entities.UserResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, err
	}
	if len(resp.Users) > 0 {
		return &entities.UserResp{
			Result: entities.Result{
				Code:    enums.USER_EXITS_CODE,
				Message: enums.USER_EXITS_CODE_MESS,
			},
		}, nil
	}
	tx, err := u.trans.BeginTransaction(ctx)
	if err != nil {
		return &entities.UserResp{
			Result: entities.Result{
				Code:    enums.TRANSACTION_INVALID_CODE,
				Message: enums.TRANSACTION_INVALID_MESS,
			},
		}, err
	}
	passwordEd, err := utils.HashPassword(user.Password)
	if err != nil {
		return &entities.UserResp{
			Result: entities.Result{
				Code:    enums.HASH_PASSWORD_ERR_CODE,
				Message: enums.HASH_PASSWORD_ERR_MESS,
			},
		}, err
	}

	respFile, err := utils.SetByCurlImage(ctx, user.File)
	if respFile.Result.Code != 0 || err != nil {
		return &entities.UserResp{
			Result:  respFile.Result,
			Created: utils.GenerateTimestamp(),
		}, nil
	}
	err = u.file.AddInformationFileStorages(ctx, tx, &domain.FileStorages{
		ID:        utils.GenerateUniqueKey(),
		TicketID:  idUser,
		URL:       respFile.URL,
		CreatedAt: utils.GenerateTimestamp(),
	})
	if err != nil {
		tx.Rollback()
		return &entities.UserResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, err
	}
	err = u.user.AddUser(ctx, tx, &domain.Users{
		Id:          idUser,
		UserName:    user.UserName,
		Password:    passwordEd,
		Age:         user.Age,
		Address:     user.Address,
		AvatarUrl:   respFile.URL,
		Role:        user.Role,
		IsActive:    enums.ACCOUNT_ACTIVE,
		ExpiredTime: user.ExpiredTime,
		CreatedAt:   utils.GenerateTimestamp(),
		UpdatedAt:   utils.GenerateTimestamp(),
	})

	if err != nil {
		tx.Rollback()
		return &entities.UserResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, err
	}
	tx.Commit()
	return &entities.UserResp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Created: utils.GenerateTimestamp(),
	}, nil
}
func (u *UseCaseUser) GetAllUserById(ctx context.Context, req *entities.UserRequestFindByForm) (*entities.UserRespFindByForm, error) {

	users, err := u.user.GetAllUserStaffs(ctx, &domain.UsersReqByForm{
		Id:        req.Id,
		UserName:  req.UserName,
		Age:       req.Age,
		Address:   req.Address,
		CreatedAt: utils.GenerateTimestamp(),
	})
	if err != nil {
		return &entities.UserRespFindByForm{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, err
	}
	return &entities.UserRespFindByForm{
		Users: users,
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Created: utils.GenerateTimestamp(),
	}, nil
}
