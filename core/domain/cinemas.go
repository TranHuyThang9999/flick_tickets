package domain

import (
	"context"
)

type Cinemas struct {
	Id             int64  `json:"id"`
	CinemaName     string `json:"cinema_name"`
	Description    string `json:"description"`
	Conscious      string `json:"conscious"`
	District       string `json:"district"`
	Commune        string `json:"commune"`
	AddressDetails string `json:"address_details"`
}
type RepositoryCinemas interface {
	AddCinema(ctx context.Context, req *Cinemas) error
	GetAllCinema(ctx context.Context) ([]*Cinemas, error)
	GetAllCinemaByName(ctx context.Context, name string) (*Cinemas, error)
	DeleteCinemaByName(ctx context.Context, name string) error
}
