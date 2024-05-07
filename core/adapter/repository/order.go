package repository

import (
	"context"
	"flick_tickets/configs"
	"flick_tickets/core/adapter"
	"flick_tickets/core/domain"
	"time"

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
func (c *CollectionOrder) UpsertOrder(ctx context.Context, email string, orderId int64, status int) error {
	err := c.collection.Model(&domain.Orders{}).
		Where("id = ? OR email = ?", orderId, email).
		UpdateColumn("status", status).Error
	return err
}
func (c *CollectionOrder) GetOrderByEmail(ctx context.Context, email string) (*domain.Orders, error) {
	var order *domain.Orders
	result := c.collection.Where("email = ?", email).First(&order)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return order, result.Error
}
func (c *CollectionOrder) UpdateOrderWhenCancel(ctx context.Context, tx *gorm.DB, id int64, status int) error {
	result := tx.Model(&domain.Orders{}).Where("id = ?", id).UpdateColumn("status", status)
	return result.Error
}
func (c *CollectionOrder) GetAllOrder(ctx context.Context, req *domain.OrdersReqByForm) ([]*domain.Orders, error) {
	var listOrder = make([]*domain.Orders, 0)
	result := c.collection.Where(&domain.Orders{
		ID:             req.ID,
		ShowTimeID:     req.ShowTimeID,
		ReleaseDate:    req.ReleaseDate,
		Email:          req.Email,
		OTP:            req.OTP,
		Description:    req.Description,
		Status:         req.Status,
		Price:          req.Price,
		Seats:          req.Seats,
		Sale:           req.Sale,
		MovieTime:      req.MovieTime,
		AddressDetails: req.AddressDetails,
		CreatedAt:      req.CreatedAt,
		UpdatedAt:      req.UpdatedAt,
	}).Order("created_at asc").Limit(req.Limit).Offset(req.Offset).Find(&listOrder)
	return listOrder, result.Error
}
func (c *CollectionOrder) TriggerOrder(ctx context.Context) ([]*domain.Orders, error) {
	var orders []*domain.Orders

	// Lấy timestamp hiện tại
	currentTimeStamp := time.Now().Unix()

	// Tính timestamp của thời điểm 15 phút trước đây
	fifteenMinutesAgo := currentTimeStamp - (15 * 60)

	// Truy vấn các đơn hàng có status = 7 và created_at nhỏ hơn thời điểm 15 phút trước đó
	result := c.collection.Where("status = ? AND created_at < ?", 7, fifteenMinutesAgo).Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}
