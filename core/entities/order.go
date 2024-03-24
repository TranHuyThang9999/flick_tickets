package entities

type OrdersReq struct {
	TicketId int64  `form:"ticket_id"`
	Email    string `form:"email"`
	Seats    int    `form:"seats"`
}
type OrdersResponseResgister struct {
	Result  Result `json:"result"`
	Created int64  `json:"created"`
}
