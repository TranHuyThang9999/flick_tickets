package entities

import "flick_tickets/core/domain"

type FileRequest struct {
	TicketId  int64  `json:"ticket"`
	Url       string `json:"url"`
	CreatedAt int    `json:"created"`
}

type UploadResponse struct {
	Result    Result `json:"result"`
	ID        string `json:"id"`
	URL       string `json:"url"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type ResponseGetListFileByObjetId struct {
	Result Result                 `json:"result"`
	Files  []*domain.FileStorages `json:"files"`
}
