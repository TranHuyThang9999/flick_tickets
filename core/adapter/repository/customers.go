package repository

import (
	"context"
	"flick_tickets/configs"
	"flick_tickets/core/adapter"
	"flick_tickets/core/domain"

	"gorm.io/gorm"
)

type CollectionCustomers struct {
	collection *gorm.DB
}

func NewCollectionCustomer(cf *configs.Configs, cus *adapter.PostGresql) domain.RepositoryCustomers {
	return &CollectionCustomers{
		collection: cus.CreateCollection(),
	}
}

// FindCustomers implements domain.RepositoryCustomers.
func (c *CollectionCustomers) FindCustomers(ctx context.Context, req *domain.CustomersFindByForm) ([]*domain.Customers, error) {

	var customers []*domain.Customers

	result := c.collection.Where(&domain.Customers{
		ID:          req.ID,
		UserName:    req.UserName,
		Password:    req.Password,
		AvatarUrl:   req.AvatarUrl,
		Address:     req.Address,
		Age:         req.Age,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		OTP:         req.OTP,
		CreatedAt:   req.CreatedAt,
		UpdatedAt:   req.UpdatedAt,
	}).Find(&customers)

	return customers, result.Error
}

// RegisterCustomers implements domain.RepositoryCustomers.
func (c *CollectionCustomers) RegisterCustomers(ctx context.Context, tx *gorm.DB, req *domain.Customers) error {
	result := tx.Create(req)
	return result.Error
}

// UpdateProfile implements domain.RepositoryCustomers.
func (c *CollectionCustomers) UpdateProfile(ctx context.Context, tx *gorm.DB, req *domain.Customers) error {
	result := tx.Where("id=?", req.ID).Updates(req)
	return result.Error
}

func (c *CollectionCustomers) UpdateWhenCheckOtp(ctx context.Context, otp int64, email string) error {
	result := c.collection.Model(&domain.Customers{}).Where("email = ? AND otp = ?", email, otp).UpdateColumn("is_active", true)
	return result.Error
}
func (c *CollectionCustomers) GetCustomersByEmailUseCheckOtp(ctx context.Context, email string, otp int64) (*domain.Customers, error) {
	var customer *domain.Customers
	result := c.collection.Where("email=? and otp =?", email, otp).First(&customer)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return customer, result.Error
}
func (c *CollectionCustomers) GetCustomerByEmail(ctx context.Context, email string) (*domain.Customers, error) {
	var customer *domain.Customers
	result := c.collection.Where("email=? ", email).First(&customer)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return customer, result.Error
}
