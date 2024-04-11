package entities

import (
	"mime/multipart"
)

type SendOtpResponse struct {
	Result    Result `json:"result"`
	CreatedAt int    `json:"created_at"`
}

type Customers struct {
	UserName    string                `form:"user_name"`
	Password    string                `form:"password"`
	File        *multipart.FileHeader `form:"file"`
	Address     string                `form:"address"`
	Age         int                   `form:"age"`
	Email       string                `form:"email"`
	PhoneNumber string                `form:"phone_number"`
}

type CustomersReqByForm struct {
	ID          int64  `form:"id"`
	UserName    string `form:"user_name"`
	Address     string `form:"address"`
	Age         int    `form:"age"`
	Email       string `form:"email"`
	PhoneNumber string `form:"phone_number"`
	CreatedAt   int    `form:"created_at"`
	UpdatedAt   int    `form:"updated_at"`
	IsActive    bool   `form:"is_active"`
}

type CustomersReqRegister struct {
	UserName    string                `form:"user_name"`
	File        *multipart.FileHeader `form:"file"`
	Address     string                `form:"address"`
	Age         int                   `form:"age"`
	Email       string                `form:"email"`
	PhoneNumber string                `form:"phone_number"`
	// IsActive    bool                  `form:"is_active"`
	// ExpiredTime int                   `form:"expired_time"`
	// Role int `form:"role"`
}

type CustomersReqRegisterResp struct {
	Result Result `json:"result"`
	Id     int64  `json:"id"`
}

type CustomersReqVerifyOtp struct {
	Email string `form:"email"`
	Otp   int64  `form:"otp"`
}
type CustomersRespVerifyOtp struct {
	Result Result `json:"result"`
}
type CustomerReqLogin struct {
	UserName string `form:"user_name"`
	Password string `form:"password"`
}
type CustomerRespLogin struct {
	Result    Result    `json:"result"`
	JwtToken  *JwtToken `json:"jwt_token"`
	CreatedAt int       `json:"created_at"`
}
type CustomersReqRegisterAdminForStaff struct {
	UserName string `form:"user_name"`
	// Password    string                `form:"password"`
	File        *multipart.FileHeader `form:"file"`
	Address     string                `form:"address"`
	Age         int                   `form:"age"`
	Email       string                `form:"email"`
	PhoneNumber string                `form:"phone_number"`
	// IsActive    bool                  `form:"is_active"`
	ExpiredTime int `form:"expired_time"`
	// Role        int                   `form:"role"`
}
type CustomersRespRegisterAdmin struct {
	Result Result `json:"result"`
	Id     int64  `json:"id"`
}

type CustomersRespFindByForm struct {
	ID          int64  `json:"id"`
	UserName    string `json:"user_name"`
	AvatarUrl   string `json:"avatar_url"`
	Address     string `json:"address"`
	Age         int    `json:"age"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	IsActive    bool   `json:"is_active"`
	CreatedAt   int    `json:"created_at"`
}
type CustomersFindByFormResp struct {
	Result    Result                     `json:"result"`
	Customers []*CustomersRespFindByForm `json:"customers"`
}
type CustomerDeleteResp struct {
	Result Result `json:"result"`
}
