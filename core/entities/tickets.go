package entities

import "mime/multipart"

type TicketreqFindByForm struct {
	ID          int64   `form:"id"`
	Name        string  `form:"name"`
	Price       float64 `form:"price"`
	MaxTicket   int64   `form:"max_ticket"`
	Quantity    int     `form:"quantity"`
	Description string  `form:"description"`
	Sale        int     `form:"sale"`
	Showtime    int     `form:"showtime"`
	ReleaseDate int     `form:"release_date"`
	CreatedAt   int     `form:"created_at"`
	UpdatedAt   int     `form:"updated_at"`
}
type TicketReqUpload struct {
	Name  string  `form:"name"`
	Price float64 `form:"price"`
	//MaxTicket    int64                 `form:"max_ticket"`
	Quantity     int                   `form:"quantity"`
	Description  string                `form:"description"`
	Sale         int                   `form:"sale"`
	Showtime     int                   `form:"showtime"`
	ReleaseDate  int                   `form:"release_date"`
	File         *multipart.FileHeader `form:"file"`
	SelectedSeat string                `form:"selected_seat"`
	UserId       int64                 `form:"user_id"`
}
type TicketRespUpload struct {
	Result    Result `json:"result"`
	CreatedAt int    `json:"created_at"`
}
