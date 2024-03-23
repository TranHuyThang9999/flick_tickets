package domain

import (
	"context"

	"gorm.io/gorm"
)

type Users struct {
	Id          int64  `json:"id"`
	UserName    string `json:"user_name"`
	Password    string `json:"password"`
	Age         int    `json:"age"`
	Address     string `json:"address"`
	AvatarUrl   string `json:"avatar_url"`
	Role        int    `json:"role"`
	IsActive    int    `json:"is_active"`
	ExpiredTime int    `json:"expired_time"`
	CreatedAt   int    `json:"created_at"`
	UpdatedAt   int    `json:"updated_at"`
}

type UsersReqByForm struct {
	Id        int64  `json:"id"`
	UserName  string `json:"user_name"`
	Age       int    `json:"age"`
	Address   string `json:"address"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
}
type UserUpdate struct {
	UserName    string `json:"user_name"`
	Password    string `json:"password"`
	Age         int    `json:"age"`
	AvatarUrl   string `json:"avatar_url"`
	Role        int    `json:"role"`
	IsActive    int    `json:"is_active"`
	ExpiredTime int    `json:"expired_time"`
	UpdatedAt   int    `json:"updated_at"`
}

type RepositoryUser interface {
	AddUser(ctx context.Context, tx *gorm.DB, user *Users) error
	GetAllUserStaffs(ctx context.Context, user *UsersReqByForm) ([]*Users, error)
	DeleteUserByUsernameStaff(ctx context.Context, tx *gorm.DB, userName string) error
	UpdateUserByUsernameStaff(ctx context.Context, tx *gorm.DB, user *UserUpdate) error
}
