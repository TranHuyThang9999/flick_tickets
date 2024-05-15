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
		Role:        req.Role,
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
func (c *CollectionCustomers) DeleteStaffByName(ctx context.Context, name string) error {
	result := c.collection.Where("user_name = ?", name).Delete(&domain.Customers{})
	return result.Error
}
func (c *CollectionCustomers) FindCustomersByUsename(ctx context.Context, user_name string) (*domain.CustomerFindByUseName, error) {
	var customer *domain.CustomerFindByUseName
	result := c.collection.Model(&domain.Customers{}).Where("user_name=? ", user_name).First(&customer)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return customer, result.Error
}
func (c *CollectionCustomers) FindCustomersByRole(ctx context.Context, user_name string, password string, role int) (*domain.CustomerFindByUseName, error) {
	var customer *domain.CustomerFindByUseName
	result := c.collection.Model(&domain.Customers{}).Where("user_name = ? and password = ? and role = ?", user_name, password, role).First(&customer)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return customer, result.Error
}
func (c *CollectionCustomers) FindByUserName(ctx context.Context, email string) (*domain.CustomerFindByUseName, error) {
	var customer *domain.CustomerFindByUseName
	result := c.collection.Model(&domain.Customers{}).Where("email = ? ", email).First(&customer)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return customer, result.Error
}
func (c *CollectionCustomers) FindAccountResetPassWord(ctx context.Context, userName, email string, role int) (int64, error) {
	var countResp int64
	result := c.collection.Model(&domain.Customers{}).
		Where("user_name = ? and email = ? and role = ? ", userName, email, role).Count(&countResp)

	return countResp, result.Error
}
func (c *CollectionCustomers) UpdatePassWord(ctx context.Context, user_name, newPassword string) error {

	result := c.collection.Model(&domain.Customers{}).Where("user_name = ?", user_name).UpdateColumn("password", newPassword)
	return result.Error
}
