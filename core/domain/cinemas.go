package domain

import (
	"context"
)

type Cinemas struct {
	Id              int64  `json:"id"`
	CinemaName      string `json:"cinema_name"`
	Description     string `json:"description"`
	Conscious       string `json:"conscious"`
	District        string `json:"district"`
	Commune         string `json:"commune"`
	AddressDetails  string `json:"address_details"`
	WidthContainer  int    `json:"width_container"`
	HeightContainer int    `json:"height_container"`
}

type CinemaReqUpdateSizeRoom struct {
	CinemaName      string `json:"cinema_name"`
	WidthContainer  int    `json:"width_container"`
	HeightContainer int    `json:"height_container"`
}

type RepositoryCinemas interface {
	AddCinema(ctx context.Context, req *Cinemas) error
	GetAllCinema(ctx context.Context) ([]*Cinemas, error)
	GetAllCinemaByName(ctx context.Context, name string) (*Cinemas, error)
	DeleteCinemaByName(ctx context.Context, name string) error
	UpdateColumnWidthHeightContainer(ctx context.Context, req *CinemaReqUpdateSizeRoom) error
}
