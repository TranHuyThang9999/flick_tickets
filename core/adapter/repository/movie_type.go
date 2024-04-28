package repository

import (
	"context"
	"flick_tickets/configs"
	"flick_tickets/core/adapter"
	"flick_tickets/core/domain"

	"gorm.io/gorm"
)

type CollectionMovie struct {
	collection *gorm.DB
}

func NewCollectionMovie(cf *configs.Configs, moive *adapter.PostGresql) domain.RepositoryMovieType {
	return &CollectionMovie{
		collection: moive.CreateCollection(),
	}
}

// AddMoiveType implements domain.RepositoryMovieType.
func (c *CollectionMovie) AddMoiveType(ctx context.Context, req *domain.MovieTypes) error {
	result := c.collection.Create(req)
	return result.Error
}

// DeleteMovieTypeById implements domain.RepositoryMovieType.
func (c *CollectionMovie) DeleteMovieTypeById(ctx context.Context, id int64) error {
	result := c.collection.Where("id = ?", id).Delete(&domain.MovieTypes{})
	return result.Error
}

// FindAllMovie implements domain.RepositoryMovieType.
func (c *CollectionMovie) FindAllMovie(ctx context.Context) ([]*domain.MovieTypes, error) {
	var listMoive = make([]*domain.MovieTypes, 0)
	result := c.collection.Find(&listMoive)
	return listMoive, result.Error
}
func (c *CollectionMovie) GetMovieTypeByName(ctx context.Context, name string) (*domain.MovieTypes, error) {
	var movieType *domain.MovieTypes
	result := c.collection.Where("movie_type_name = ? ", name).First(&movieType)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return movieType, result.Error
}
