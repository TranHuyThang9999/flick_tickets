package entities

type CinemasReq struct {
	CinemaName string `form:"cinema_name"`
}
type CinemasRes struct {
	Result Result `json:"result"`
}
