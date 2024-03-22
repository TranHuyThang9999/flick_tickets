package usecase

import (
	"context"
	"flick_tickets/common/utils"
	"flick_tickets/core/domain"
	"flick_tickets/core/entities"
)

type UseCaseUser struct {
	user domain.RepositoryUser
}

func NewUseCaseUser(user domain.RepositoryUser) *UseCaseUser {
	return &UseCaseUser{
		user: user,
	}
}

func (u *UseCaseUser) AddUser(ctx context.Context, user *entities.Users) (*entities.UserResp, error) {

	respFile := utils.SetByCurlImage(ctx, user.AvatarUrl)
	if respFile.Result.Code != 0 {
		return &entities.UserResp{
			Result:  respFile.Result,
			Created: utils.GenerateTimestamp(),
		}, nil
	}

	err := u.user.AddUser(&domain.Users{
		Id:          utils.GenerateUniqueKey(),
		UserName:    user.UserName,
		Age:         user.Age,
		AvatarUrl:   respFile.URL,
		Role:        1,
		IsActive:    user.IsActive,
		ExpiredTime: user.ExpiredTime,
		CreatedAt:   utils.GenerateTimestamp(),
		UpdatedAt:   utils.GenerateTimestamp(),
	})
	if err != nil {
		return &entities.UserResp{
			Result: entities.Result{
				Code:    1,
				Message: "error creating user",
			},
		}, err
	}
	return &entities.UserResp{}, nil
}
