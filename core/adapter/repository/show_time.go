package repository

import (
	"context"
	"flick_tickets/configs"
	"flick_tickets/core/adapter"
	"flick_tickets/core/domain"

	"gorm.io/gorm"
)

type CollectionShowTIme struct {
	collection *gorm.DB
}

// AddShowTime implements domain.RepositoryShowTime.
func (c *CollectionShowTIme) AddShowTime(ctx context.Context, req *domain.ShowTime) error {
	result := c.collection.Create(req)
	return result.Error
}

// DeleteShowTimeByTicketId implements domain.RepositoryShowTime.
func (c *CollectionShowTIme) DeleteShowTimeByTicketId(ctx context.Context, req *domain.ShowTimeDelete) error {

	result := c.collection.Where(" id = ? and movie_time= ? and ticket_id = ? ", req.ID, req.MovieTime, req.TicketID).Delete(&domain.ShowTime{})
	return result.Error
}

func (c *CollectionShowTIme) GetTimeUseCheckAddTicket(ctx context.Context, req *domain.ShowTimeCheckList) ([]*domain.ShowTime, error) {
	var ShowTimeList []*domain.ShowTime
	result := c.collection.Where("cinema_name = ? and movie_time =? ", req.CinemaName, req.MovieTime).Find(&ShowTimeList)
	return ShowTimeList, result.Error
}

func NewCollectionShowTime(cf *configs.Configs, sh *adapter.PostGresql) domain.RepositoryShowTime {
	return &CollectionShowTIme{
		collection: sh.CreateCollection(),
	}
}
