package entities

import "flick_tickets/core/domain"

type MovieTypesAddReq struct {
	MovieTypeName string `form:"movieTypeName"`
}
type MovieTypesAddResp struct {
	Result Result `json:"result"`
}
type MovieGetAllResp struct {
	Result Result               `json:"result"`
	Movie  []*domain.MovieTypes `json:"movie"`
}
