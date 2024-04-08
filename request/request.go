package request

import "mime/multipart"

type InforFileReq struct {
	Files *multipart.FileHeader `form:"image"`
}

type InforFileResp struct {
	Result    Result `json:"result"`
	Id        string `json:"id"`
	Url       string `json:"url"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
