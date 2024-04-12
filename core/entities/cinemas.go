package entities

import "flick_tickets/core/domain"

type CinemasReq struct {
	CinemaName     string `form:"cinema_name"`
	Description    string `form:"description"`
	Conscious      string `form:"conscious"`
	District       string `form:"district"`
	Commune        string `form:"commune"`
	AddressDetails string `form:"address_details"`
}
type CinemasRes struct {
	Result Result `json:"result"`
}
type CinemasResGetAll struct {
	Result  Result            `json:"result"`
	Cinemas []*domain.Cinemas `json:"cinemas"`
}
type CinemasRespDelete struct {
	Result Result `json:"result"`
}
type CinemasRespGetByName struct {
	Cinema *domain.Cinemas `json:"cinema"`
	Result Result          `json:"result"`
}
