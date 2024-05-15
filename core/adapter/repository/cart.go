package repository

import (
	"context"
	"flick_tickets/configs"
	"flick_tickets/core/adapter"
	"flick_tickets/core/domain"

	"gorm.io/gorm"
)

type Collectioncart struct {
	collection *gorm.DB
}

func NewCollectionCart(cf *configs.Configs, cart *adapter.PostGresql) domain.RepositoryCarts {
	return &Collectioncart{
		collection: cart.CreateCollection(),
	}
}

// AddCarts implements domain.RepositoryCarts.
func (c *Collectioncart) AddCarts(ctx context.Context, req *domain.Carts) error {
	result := c.collection.Create(req)
	return result.Error
}

// DeleteCartById implements domain.RepositoryCarts.
func (c *Collectioncart) DeleteCartById(ctx context.Context, cartId int64) error {
	result := c.collection.Where("id = ?", cartId).Delete(&domain.Carts{})
	return result.Error
}

// FindCartByForm implements domain.RepositoryCarts.
func (c *Collectioncart) FindCartByForm(ctx context.Context, req *domain.CartFindByFormReq) ([]*domain.Carts, error) {
	var listItem = make([]*domain.Carts, 0)
	result := c.collection.Where(&domain.Carts{
		Id:            req.Id,
		UserName:      req.UserName,
		ShowTimeId:    req.ShowTimeId,
		SeatsPosition: req.SeatsPosition,
		Price:         req.Price,
		// CreatedAt:     req.CreatedAt,
		// UpdatedAt:     req.UpdatedAt,
	}).Find(&listItem)
	return listItem, result.Error
}

// UpdateCartById implements domain.RepositoryCarts.
func (c *Collectioncart) UpdateCartById(ctx context.Context, req *domain.Carts) error {
	result := c.collection.Where("id = ?", req.Id).Updates(req)
	return result.Error
}
func (c *Collectioncart) DeleteCartByShowTimeId(ctx context.Context, show_time_id int64) error {
	result := c.collection.Where("show_time_id = ?", show_time_id).Delete(&domain.Carts{})
	return result.Error
}
