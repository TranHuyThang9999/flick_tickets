package repository

import (
	"context"
	"flick_tickets/configs"
	"flick_tickets/core/adapter"
	"flick_tickets/core/domain"

	"gorm.io/gorm"
)

type CollectionCinemas struct {
	collection *gorm.DB
}

// AddCinema implements domain.RepositoryCinemas.
func (c *CollectionCinemas) AddCinema(ctx context.Context, tx *gorm.DB, req *domain.Cinemas) error {
	result := tx.Create(req)
	return result.Error
}

func NewCollectionCinemas(cf *configs.Configs, cm *adapter.PostGresql) domain.RepositoryCinemas {
	return &CollectionCinemas{
		collection: cm.CreateCollection(),
	}
}
