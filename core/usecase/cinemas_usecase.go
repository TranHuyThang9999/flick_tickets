package usecase

import (
	"context"
	"flick_tickets/core/domain"
	"flick_tickets/core/entities"
)

type UseCaseCinemas struct {
	cm domain.RepositoryCinemas
}

func NewUseCaseCinemas(cm domain.RepositoryCinemas) *UseCaseCinemas {
	return &UseCaseCinemas{
		cm: cm,
	}
}
func (c *UseCaseCinemas) AddCinemas(ctx context.Context, req *entities.CinemasReq) (*entities.CinemasRes, error) {
	panic("UseCaseCinemas")
}
