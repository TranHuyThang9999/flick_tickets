package entities

import "flick_tickets/core/domain"

type ShowTimeAddReq struct {
	TicketID     int64   `form:"ticket_id"`
	CinemaName   string  `form:"cinema_name"`
	MovieTime    string  `form:"movie_time"`
	SelectedSeat string  `form:"selected_seat"`
	Quantity     int     `form:"quantity"`
	Price        float64 `form:"price"`
	Discount     int     `form:"discount"`
}
type ShowTimeDelete struct {
	ID        int64 `form:"id"`
	TicketID  int64 `form:"ticket_id"`
	MovieTime int   `form:"movie_time"`
}
type ShowTimeAddResponse struct {
	Result    Result `json:"result"`
	CreatedAt int    `json:"created_at"`
}
type ShowTimeDeleteResponse struct {
	Result    Result `json:"result"`
	CreatedAt int    `json:"created_at"`
}
type ShowTimeByTicketIdresp struct {
	Result    Result      `json:"result"`
	Showtimes []*ShowTime `json:"showtimes"`
}

type ShowTime struct {
	ID              int64   `json:"id"`
	TicketID        int64   `json:"ticket_id"`
	CinemaName      string  `json:"cinema_name"`
	MovieTime       int     `json:"movie_time"` //string
	Description     string  `json:"description"`
	Conscious       string  `json:"conscious"`
	District        string  `json:"district"`
	Commune         string  `json:"commune"`
	AddressDetails  string  `json:"address_details"`
	WidthContainer  int     `json:"width_container"`
	HeightContainer int     `json:"height_container"`
	SelectedSeat    string  `json:"selected_seat"`
	Quantity        int     `json:"quantity"`
	OriginalNumber  int     `json:"original_number"`
	Price           float64 `json:"price"`
	Discount        int     `json:"discount"`
}
type ShowTimeDetail struct {
	Result   Result           `json:"result"`
	ShowTime *domain.ShowTime `json:"show_time,omitempty"`
}
type ShowTimeDeleteByIdResp struct {
	Result    Result `json:"result"`
	CreatedAt int    `json:"created_at"`
}

type ShowTimeFindByIdResp struct {
	Result   Result       `json:"result"`
	ShowTime ShowTimeResp `json:"show_time"`
}

type ShowTimeResp struct {
	ID             int64   `json:"id"`
	TicketID       int64   `json:"ticket_id"`
	CinemaName     string  `json:"cinema_name"`
	Price          float64 `json:"price"`
	MovieTime      int     `json:"movie_time"` //string
	SelectedSeat   string  `json:"selected_seat"`
	Quantity       int     `json:"quantity"`
	OriginalNumber int     `json:"original_number"`
	Discount       int     `json:"discount"`
	CreatedAt      int     `json:"created_at"`
	UpdatedAt      int     `json:"updated_at"`
}
type ShowTimeUpdateByIdResp struct {
	Result Result `json:"result"`
}
