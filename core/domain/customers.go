package domain

import (
	"context"

	"gorm.io/gorm"
)

type Customers struct {
	ID          int64  `json:"id"`
	UserName    string `json:"user_name"`
	Password    string `json:"password"`
	AvatarUrl   string `json:"avatar_url"`
	Address     string `json:"address"`
	Age         int    `json:"age"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	OTP         int64  `json:"otp"`
	IsActive    bool   `json:"is_active"`
	ExpiredTime int    `json:"expired_time"`
	Role        int    `json:"role"`
	CreatedAt   int    `json:"created_at"`
	UpdatedAt   int    `json:"updated_at"`
}
type CustomersFindByForm struct {
	ID          int64  `form:"id"`
	UserName    string `form:"user_name"`
	Password    string `form:"password"`
	AvatarUrl   string `form:"avatar_url"`
	Address     string `form:"address"`
	Age         int    `form:"age"`
	Email       string `form:"email"`
	PhoneNumber string `form:"phone_number"`
	OTP         int64  `form:"otp"`
	CreatedAt   int    `form:"created_at"`
	UpdatedAt   int    `form:"updated_at"`
	Role        int    `form:"role"`
}
type CustomerFindByUseName struct {
	ID          int64  `json:"id"`
	UserName    string `json:"user_name"`
	AvatarUrl   string `json:"avatar_url"`
	Address     string `json:"address"`
	Age         int    `json:"age"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   int    `json:"created_at"`
	UpdatedAt   int    `json:"updated_at"`
}
type RepositoryCustomers interface {
	FindCustomers(ctx context.Context, req *CustomersFindByForm) ([]*Customers, error)
	RegisterCustomers(ctx context.Context, tx *gorm.DB, req *Customers) error
	UpdateProfile(ctx context.Context, tx *gorm.DB, req *Customers) error
	UpdateWhenCheckOtp(ctx context.Context, otp int64, email string) error
	GetCustomersByEmailUseCheckOtp(ctx context.Context, email string, otp int64) (*Customers, error)
	GetCustomerByEmail(ctx context.Context, email string) (*Customers, error)
	DeleteStaffByName(ctx context.Context, name string) error
	FindCustomersByUsename(ctx context.Context, name string) (*CustomerFindByUseName, error)
	FindCustomersByRole(ctx context.Context, user_name string, password string, role int) (*CustomerFindByUseName, error)
	FindByUserName(ctx context.Context, email string) (*CustomerFindByUseName, error)
	FindAccountResetPassWord(ctx context.Context, userName, email string, role int) (int64, error)
	UpdatePassWord(ctx context.Context, user_name, newPassword string) error
}
