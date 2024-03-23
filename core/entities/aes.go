package entities

type TokenRequestSendQrCode struct {
	Id        int64  `json:"id"`
	FromEmail string `json:"fromEmail"`
}
type TokenRespSendQrCode struct {
	Token   string `json:"token"`
	Result  Result `json:"result"`
	Created int    `json:"created"`
}
type TokenReqCheckQrCode struct {
	Token string `json:"token"`
}
type TokenResponseCheckQrCode struct {
	Result  Result `json:"result"`
	Created int    `json:"created"`
}
