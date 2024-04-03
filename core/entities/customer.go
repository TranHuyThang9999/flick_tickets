package entities

import "mime/multipart"

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
	Password    string `form:"password"`
	AvatarUrl   string `form:"avatar_url"`
	Address     string `form:"address"`
	Age         int    `form:"age"`
	Email       string `form:"email"`
	PhoneNumber string `form:"phone_number"`
	OTP         int64  `form:"otp"`
	CreatedAt   int    `form:"created_at"`
	UpdatedAt   int    `form:"updated_at"`
}
type CustomersReqVerifyOtp struct {
	Email string `form:"email"`
	Otp   int64  `form:"otp"`
}
type CustomersRespVerifyOtp struct {
	Result Result `json:"result"`
}
