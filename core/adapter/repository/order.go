package repository

import (
	"context"
	"flick_tickets/configs"
	"flick_tickets/core/adapter"
	"flick_tickets/core/domain"

	"gorm.io/gorm"
)

type CollectionOrder struct {
	collection *gorm.DB
}

// CancelTicket implements domain.RepositoryOrder.
func (c *CollectionOrder) CancelTicket(ctx context.Context, id int64) error {
	panic("unimplemented")
}

// RegisterTicket implements domain.RepositoryOrder.
func (c *CollectionOrder) RegisterTicket(ctx context.Context, tx *gorm.DB, req *domain.Orders) error {
	result := tx.Create(req)
	return result.Error
}

func NewCollectionOrder(cf *configs.Configs, order *adapter.PostGresql) domain.RepositoryOrder {
	return &CollectionOrder{
		collection: order.CreateCollection(),
	}
}
