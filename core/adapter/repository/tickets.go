package repository

import (
	"context"
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
		ID:          req.ID,
		Name:        req.Name,
		Price:       req.Price,
		Quantity:    req.Quantity,
		Description: req.Description,
		Sale:        req.Sale,
		Showtime:    req.Showtime,
		ReleaseDate: req.ReleaseDate,
		CreatedAt:   req.CreatedAt,
		UpdatedAt:   req.UpdatedAt,
	}).Find(tickets)
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
