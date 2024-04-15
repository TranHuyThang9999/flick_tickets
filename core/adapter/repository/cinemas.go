package repository

import (
	"context"
	"flick_tickets/configs"
	"flick_tickets/core/adapter"
	"flick_tickets/core/domain"

	"gorm.io/gorm"
)

type CollectionCinemas struct {
	collection *gorm.DB
}

func NewCollectionCinemas(cf *configs.Configs, cm *adapter.PostGresql) domain.RepositoryCinemas {
	return &CollectionCinemas{
		collection: cm.CreateCollection(),
	}
}

// AddCinema implements domain.RepositoryCinemas.
func (c *CollectionCinemas) AddCinema(ctx context.Context, req *domain.Cinemas) error {
	result := c.collection.Create(req)
	return result.Error
}
func (c *CollectionCinemas) GetAllCinema(ctx context.Context) ([]*domain.Cinemas, error) {
	var cinemas = make([]*domain.Cinemas, 0)
	result := c.collection.Find(&cinemas)
	return cinemas, result.Error
}

func (c *CollectionCinemas) GetAllCinemaByName(ctx context.Context, name string) (*domain.Cinemas, error) {
	var cinemas *domain.Cinemas
	result := c.collection.Where("cinema_name = ?", name).First(&cinemas)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return cinemas, result.Error
}
func (c *CollectionCinemas) DeleteCinemaByName(ctx context.Context, name string) error {
	result := c.collection.Where("cinema_name = ?", name).Delete(&domain.Cinemas{})
	return result.Error
}
func (c *CollectionCinemas) UpdateColumnWidthHeightContainer(ctx context.Context, req *domain.CinemaReqUpdateSizeRoom) error {
	result := c.collection.Model(&domain.Cinemas{}).Where("cinema_name = ?", req.CinemaName).Updates(&req)
	return result.Error
}
