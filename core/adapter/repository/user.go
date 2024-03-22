package repository

import (
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
func (c *CollectionUser) AddUser(user *domain.Users) error {
	result := c.collection.Create(user)
	return result.Error
}

// DeleteUserByUsernameStaff implements domain.RepositoryUser.
func (c *CollectionUser) DeleteUserByUsernameStaff(userName string) error {
	result := c.collection.Where("user_name =?", userName).Delete(&domain.Users{})
	return result.Error
}

// GetAllUserStaffs implements domain.RepositoryUser.
func (c *CollectionUser) GetAllUserStaffs(user *domain.UsersReqByForm) ([]*domain.Users, error) {
	var users = make([]*domain.Users, 0)
	result := c.collection.Where(&domain.Users{
		Id:          user.Id,
		UserName:    user.UserName,
		Age:         user.Age,
		AvatarUrl:   user.AvatarUrl,
		Role:        user.Role,
		IsActive:    user.IsActive,
		ExpiredTime: user.ExpiredTime,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}).Find(&users)
	return users, result.Error
}

// UpdateUserByUsernameStaff implements domain.RepositoryUser.
func (c *CollectionUser) UpdateUserByUsernameStaff(user *domain.UserUpdate) error {
	result := c.collection.Model(&domain.Users{}).Where("user_name = ?", user.UserName).Updates(user)
	return result.Error
}
