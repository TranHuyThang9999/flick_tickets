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

func NewCollectionShowTime(cf *configs.Configs, sh *adapter.PostGresql) domain.RepositoryShowTime {
	return &CollectionShowTIme{
		collection: sh.CreateCollection(),
	}
}

// AddShowTime implements domain.RepositoryShowTime.
func (c *CollectionShowTIme) AddShowTime(ctx context.Context, req *domain.ShowTime) error {
	result := c.collection.Create(req)
	return result.Error
}
func (c *CollectionShowTIme) AddListShowTime(ctx context.Context, tx *gorm.DB, req []*domain.ShowTime) error {
	result := tx.Create(req)
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

func (c *CollectionShowTIme) FindDuplicateShowTimes(ctx context.Context, movieTimes []int, cinemaName []string) ([]*domain.ShowTime, error) {
	var result []*domain.ShowTime

	err := c.collection.Table("show_times").
		Select("movie_time, cinema_name, COUNT(*) AS record_count").
		Where("cinema_name in (?)", cinemaName).
		Where("movie_time IN (?)", movieTimes).
		Group("movie_time, cinema_name").
		Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}
func (c *CollectionShowTIme) GetShowTimeByTicketId(ctx context.Context, ticketId int64) ([]*domain.ShowTime, error) {

	var listShowTimeByTicketId = make([]*domain.ShowTime, 0)
	result := c.collection.Where("ticket_id= ? ", ticketId).Find(&listShowTimeByTicketId)
	return listShowTimeByTicketId, result.Error
}
func (c *CollectionShowTIme) GetAll(ctx context.Context) ([]*domain.ShowTime, error) {
	var listShowTimeByTicketId = make([]*domain.ShowTime, 0)
	result := c.collection.Find(&listShowTimeByTicketId)
	return listShowTimeByTicketId, result.Error
}
