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

func NewCollectionOrder(cf *configs.Configs, order *adapter.PostGresql) domain.RepositoryOrder {
	return &CollectionOrder{
		collection: order.CreateCollection(),
	}
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
func (c *CollectionOrder) GetOrderById(ctx context.Context, id int64) (*domain.Orders, error) {
	var order *domain.Orders
	result := c.collection.Where("id = ?", id).First(&order)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return order, result.Error
}
func (c *CollectionOrder) UpsertOrder(ctx context.Context, orderId int64, status int) error {
	err := c.collection.Model(&domain.Orders{}).Where("id = ?", orderId).UpdateColumn("status", status)
	return err.Error
}
