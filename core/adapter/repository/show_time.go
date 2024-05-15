package repository

import (
	"context"
	"flick_tickets/configs"
	"flick_tickets/core/adapter"
	"flick_tickets/core/domain"

	"gorm.io/gorm"
)

type CollectionShowTime struct {
	collection *gorm.DB
}

// UpsertListShowTime implements domain.RepositoryShowTime.
func NewCollectionShowTime(cf *configs.Configs, sh *adapter.PostGresql) domain.RepositoryShowTime {
	return &CollectionShowTime{
		collection: sh.CreateCollection(),
	}
}

// AddShowTime implements domain.RepositoryShowTime.
func (c *CollectionShowTime) AddShowTime(ctx context.Context, req *domain.ShowTime) error {
	result := c.collection.Create(req)
	return result.Error
}
func (c *CollectionShowTime) AddListShowTime(ctx context.Context, tx *gorm.DB, req []*domain.ShowTime) error {
	result := tx.Create(req)
	return result.Error
}

// DeleteShowTimeByTicketId implements domain.RepositoryShowTime.
func (c *CollectionShowTime) DeleteShowTimeByTicketId(ctx context.Context, req *domain.ShowTimeDelete) error {

	result := c.collection.Where(" id = ? and movie_time= ? and ticket_id = ? ", req.ID, req.MovieTime, req.TicketID).Delete(&domain.ShowTime{})
	return result.Error
}

func (c *CollectionShowTime) GetTimeUseCheckAddTicket(ctx context.Context, req *domain.ShowTimeCheckList) (*domain.ShowTime, error) {
	var ShowTimeList *domain.ShowTime
	result := c.collection.Where("cinema_name = ? and movie_time = ? ", req.CinemaName, req.MovieTime).First(&ShowTimeList)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return ShowTimeList, result.Error
}

func (c *CollectionShowTime) FindDuplicateShowTimes(ctx context.Context, movieTimes []int, cinemaName []string) ([]*domain.ShowTime, error) {
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
func (c *CollectionShowTime) GetShowTimeByTicketId(ctx context.Context, ticketId int64) ([]*domain.ShowTime, error) {

	var listShowTimeByTicketId = make([]*domain.ShowTime, 0)
	result := c.collection.Where("ticket_id= ? ", ticketId).Find(&listShowTimeByTicketId)
	return listShowTimeByTicketId, result.Error
}
func (c *CollectionShowTime) GetAll(ctx context.Context) ([]*domain.ShowTime, error) {
	var listShowTimeByTicketId = make([]*domain.ShowTime, 0)
	result := c.collection.Find(&listShowTimeByTicketId)
	return listShowTimeByTicketId, result.Error
}
func (c *CollectionShowTime) UpdateSelectedSeat(ctx context.Context) error {
	panic("")
}
func (c *CollectionShowTime) GetInformationShowTimeForTicketByTicketId(ctx context.Context, id int64) (*domain.ShowTime, error) {
	var showTimeByTicketId *domain.ShowTime
	result := c.collection.Where("id = ?", id).Find(&showTimeByTicketId)
	return showTimeByTicketId, result.Error
}
func (c *CollectionShowTime) UpdateQuantitySeat(ctx context.Context, tx *gorm.DB, showTimeId int64, quantity int, selected_seat string) error {
	if err := tx.Model(&domain.ShowTime{}).Where("id = ?", showTimeId).
		UpdateColumns(map[string]interface{}{
			"quantity":      quantity,
			"selected_seat": selected_seat,
		}).Error; err != nil {
		return err
	}
	return nil
}
func (c *CollectionShowTime) UpdatePriceShowTimeByTicketId(ctx context.Context, tx *gorm.DB, ticketId int64, price float64) error {
	result := tx.Model(&domain.ShowTime{}).Where("", ticketId).UpdateColumn("price", price)
	return result.Error
}
func (c *CollectionShowTime) DeleteByTicketIdAndNameCinema(ctx context.Context, tx *gorm.DB, ticketId int64, nameCinema []string) error {
	// Xóa các bản ghi trong bảng CollectionShowTime dựa trên ticketId và nameCinema
	if err := tx.Table("CollectionShowTime").
		Where("ticket_id = ?", ticketId).
		Where("name_cinema IN (?)", nameCinema).
		Delete(nil).Error; err != nil {
		return err
	}

	return nil
}
func (c *CollectionShowTime) GetShowTimeByNameCinema(ctx context.Context, cinema_name string) ([]*domain.ShowTime, error) {
	var listCinema = make([]*domain.ShowTime, 0)
	result := c.collection.Where("cinema_name = ?", cinema_name).Find(&listCinema)
	return listCinema, result.Error
}
func (c *CollectionShowTime) GetListShowTimeByListId(ctx context.Context, ids []int64) ([]*domain.ShowTime, error) {
	var listCinema = make([]*domain.ShowTime, 0)
	result := c.collection.Where("id in(?)", ids).Find(&listCinema)
	return listCinema, result.Error
}

func (c *CollectionShowTime) DeleteShowTimesByTicketId(ctx context.Context, tx *gorm.DB, ticketId int64) error {
	result := tx.Where("ticket_id = ?", ticketId).Delete(&domain.ShowTime{})
	return result.Error
}
func (c *CollectionShowTime) UpsertListShowTime(ctx context.Context, req []*domain.ShowTime) error {
	result := c.collection.Create(req)
	return result.Error
}
func (c *CollectionShowTime) DeleteShowTimeByid(ctx context.Context, tx *gorm.DB, show_time_id int64) error {
	result := tx.Where("id = ?", show_time_id).Delete(&domain.ShowTime{})
	return result.Error
}
func (c *CollectionShowTime) GetShowTimeById(ctx context.Context, show_time_id int64) (*domain.ShowTime, error) {
	var showTime *domain.ShowTime
	result := c.collection.Where("id = ? ", show_time_id).First(&showTime)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return showTime, result.Error
}
func (c *CollectionShowTime) UpdateShowTimeById(ctx context.Context, req *domain.ShowTimeUpdateReq) error {
	result := c.collection.Model(&domain.ShowTime{}).Where("id  =?", req.ID).Updates(&req)
	return result.Error
}

// FindDuplicateShowTime tìm các lịch chiếu trùng lặp dựa trên thời gian phim và tên rạp.
func (c *CollectionShowTime) FindDuplicateShowTimeUseUpdate(ctx context.Context, movieTime int, cinemaName string) ([]*domain.ShowTime, error) {
	var result []*domain.ShowTime

	err := c.collection.Table("show_times").
		Select("movie_time, cinema_name, COUNT(*) AS record_count").
		Where("cinema_name = ?", cinemaName).
		Where("movie_time = ?", movieTime).
		Group("movie_time, cinema_name").
		Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}
func (c *CollectionShowTime) GetAllShowTime(ctx context.Context) ([]*domain.ShowTime, error) {
	var listShowTime = make([]*domain.ShowTime, 0)
	result := c.collection.Find(&listShowTime)
	return listShowTime, result.Error
}
