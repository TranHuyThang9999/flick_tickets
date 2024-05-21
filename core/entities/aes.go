package entities

type TokenRequestSendQrCode struct {
	Title     string                  `form:"title"`
	Content   string                  `form:"content"`
	FromEmail string                  `form:"fromEmail"`
	Order     *OrderSendTicketToEmail `form:"order"`
}
type TokenRespSendQrCode struct {
	Token     string `json:"token"`
	Result    Result `json:"result"`
	CreatedAt int    `json:"created_at"`
}
type TokenReqCheckQrCode struct {
	Token string `form:"token"`
}
type TokenResponseCheckQrCode struct {
	Result    Result                `json:"result"`
	Order     *OrderHistoryEntities `json:"order"`
	CreatedAt int                   `json:"created_at"`
}
type AesContentEncryptReq struct {
	Token string `form:"token"`
}
