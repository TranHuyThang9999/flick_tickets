package entities

import "mime/multipart"

type Users struct {
	UserName    string                `json:"user_name"`
	Age         int                   `json:"age"`
	AvatarUrl   *multipart.FileHeader `json:"avatar_url"`
	Role        int                   `json:"role"`
	IsActive    int                   `json:"is_active"`
	ExpiredTime int                   `json:"expired_time"`
}
type UserResp struct {
	Result  Result `json:"result"`
	Created int    `json:"created"`
}
