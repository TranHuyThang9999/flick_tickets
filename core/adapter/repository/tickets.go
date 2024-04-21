package repository

import (
	"context"
	"errors"
	"flick_tickets/configs"
	"flick_tickets/core/adapter"
	"flick_tickets/core/domain"

	"gorm.io/gorm"
)

type CollectionTickets struct {
	collection *gorm.DB
}

func NewCollectionTickets(cf *configs.Configs, ticket *adapter.PostGresql) domain.RepositoryTickets {
	return &CollectionTickets{
		collection: ticket.CreateCollection(),
	}
}

func (u *CollectionTickets) AddTicket(ctx context.Context, tx *gorm.DB, req *domain.Tickets) error {
	result := tx.Create(req)
	return result.Error
}
func (u *CollectionTickets) GetAllTickets(ctx context.Context, req *domain.TicketreqFindByForm) ([]*domain.Tickets, error) {
	var tickets []*domain.Tickets
	result := u.collection.Where(&domain.Tickets{
		ID:            req.ID,
		Name:          req.Name,
		Description:   req.Description,
		Sale:          req.Sale,
		ReleaseDate:   req.ReleaseDate,
		Status:        req.Status,
		MovieDuration: req.MovieDuration,
		AgeLimit:      req.AgeLimit,
		Director:      req.Director,
		Actor:         req.Actor,
		Producer:      req.Producer,
		CreatedAt:     req.CreatedAt,
		UpdatedAt:     req.UpdatedAt,
	}).Find(&tickets)
	return tickets, result.Error
}
func (u *CollectionTickets) UpdateTicketById(ctx context.Context, tx *gorm.DB, req *domain.TicketReqUpdateById) error {
	result := tx.Where("id=?", req.ID).Updates(&req)
	return result.Error
}
func (u *CollectionTickets) DeleteTicketsById(ctx context.Context, tx *gorm.DB, id int64) error {
	result := tx.Where("id=?", id).Delete(&domain.Tickets{})
	return result.Error
}
func (u *CollectionTickets) UpdateTicketQuantity(ctx context.Context, tx *gorm.DB, id int64, quantity int) error {

	if err := tx.Model(&domain.Tickets{}).Where("id = ?", id).UpdateColumn("quantity", quantity).Error; err != nil {
		return err
	}
	return nil
}
func (u *CollectionTickets) GetTicketById(ctx context.Context, id int64) (*domain.Tickets, error) {
	var ticket *domain.Tickets
	result := u.collection.Where("id = ?", id).First(&ticket)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Trả về nil nếu không tìm thấy bản ghi
			return nil, nil
		}
		// Xử lý lỗi khác nếu có
		return nil, result.Error
	}
	return ticket, nil
}

// func (u *CollectionTickets) UpdateTicketSelectedSeat(ctx context.Context, tx *gorm.DB, id int64, selected_seat string) error {

// 	if err := tx.Model(&domain.Tickets{}).Where("id = ?", id).UpdateColumn("selected_seat", selected_seat).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
