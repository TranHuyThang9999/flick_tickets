package entities

type TokenRequestSendQrCode struct {
	Title     string                  `form:"title"`
	Content   string                  `form:"content"`
	FromEmail string                  `form:"fromEmail"`
	Order     *OrderSendTicketToEmail `form:"order"`
}
type TokenRespSendQrCode struct {
	Token   string `json:"token"`
	Result  Result `json:"result"`
	Created int    `json:"created"`
}
type TokenReqCheckQrCode struct {
	Token string `form:"token"`
}
type TokenResponseCheckQrCode struct {
	Result  Result `json:"result"`
	Content string `json:"content"`
	Created int    `json:"created"`
}
