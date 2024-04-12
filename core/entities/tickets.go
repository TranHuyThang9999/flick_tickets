package entities

import (
	"flick_tickets/core/domain"
	"mime/multipart"
)

type TicketreqFindByForm struct {
	ID          int64   `form:"id"`
	Name        string  `form:"name"`
	Price       float64 `form:"price"`
	MaxTicket   int64   `form:"max_ticket"`
	Quantity    int     `form:"quantity"`
	Description string  `form:"description"`
	Sale        int     `form:"sale"`
	Showtime    string  `form:"showtime"`
	ReleaseDate int     `form:"release_date"`
	Status      int     `form:"status"`
	CreatedAt   int     `form:"created_at"`
	UpdatedAt   int     `form:"updated_at"`
}
type TicketReqUpload struct {
	Name          string                  `form:"name"`
	Price         float64                 `form:"price"`
	Quantity      int                     `form:"quantity"`
	Description   string                  `form:"description"`
	Sale          int                     `form:"sale"`
	Status        int                     `form:"status"`
	ReleaseDate   int                     `form:"release_date"`
	File          []*multipart.FileHeader `form:"file"`
	SelectedSeat  string                  `form:"selected_seat"`
	CinemaName    string                  `form:"cinema_name"`
	MovieTime     string                  `form:"movie_time"`
	MovieDuration int                     `form:"movieDuration"`
	AgeLimit      int                     `form:"age_limit"`
	Director      string                  `form:"director"`
	Actor         string                  `form:"actor"`
	Producer      string                  `form:"producer"`
}
type TicketRespUpload struct {
	Result    Result `json:"result"`
	CreatedAt int    `json:"created_at"`
}
type TicketRespgetById struct {
	Result    Result          `json:"result"`
	Ticket    *domain.Tickets `json:"ticket"`
	CreatedAt int             `json:"created_at"`
}

type Tickets struct {
	Ticket    *domain.Tickets        `json:"ticket"`
	ShowTimes []*domain.ShowTime     `json:"show_times"`
	ListUrl   []*domain.FileStorages `json:"list_url"`
}

type TicketRespGetAll struct {
	Result Result `json:"result"`
	//Tickets []*Tickets `json:"tickets"`
	ListTickets []*domain.Tickets `json:"list_tickets"`
}
