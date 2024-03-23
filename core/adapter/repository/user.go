package repository

import (
	"context"
	"flick_tickets/configs"
	"flick_tickets/core/adapter"
	"flick_tickets/core/domain"

	"gorm.io/gorm"
)

type CollectionUser struct {
	collection *gorm.DB
}

func NewCollectionUser(cf *configs.Configs, user *adapter.PostGresql) domain.RepositoryUser {
	return &CollectionUser{
		collection: user.CreateCollection(),
	}
}

// AddUser implements domain.RepositoryUser.
func (c *CollectionUser) AddUser(ctx context.Context, tx *gorm.DB, user *domain.Users) error {
	result := tx.Create(user)
	return result.Error
}

// DeleteUserByUsernameStaff implements domain.RepositoryUser.
func (c *CollectionUser) DeleteUserByUsernameStaff(ctx context.Context, tx *gorm.DB, userName string) error {
	result := tx.Where("user_name =?", userName).Delete(&domain.Users{})
	return result.Error
}

// GetAllUserStaffs implements domain.RepositoryUser.
func (c *CollectionUser) GetAllUserStaffs(ctx context.Context, user *domain.UsersReqByForm) ([]*domain.Users, error) {
	var users = make([]*domain.Users, 0)
	result := c.collection.Where(&domain.Users{
		Id:        user.Id,
		UserName:  user.UserName,
		Age:       user.Age,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}).Find(&users)
	return users, result.Error
}

// UpdateUserByUsernameStaff implements domain.RepositoryUser.
func (c *CollectionUser) UpdateUserByUsernameStaff(ctx context.Context, tx *gorm.DB, user *domain.UserUpdate) error {
	result := tx.Model(&domain.Users{}).Where("user_name = ?", user.UserName).Updates(user)
	return result.Error
}
