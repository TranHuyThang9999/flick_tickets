package domain

import "context"

type MovieTypes struct {
	Id            int64  `json:"id"`
	MovieTypeName string `json:"movieTypeName"`
}

type RepositoryMovieType interface {
	AddMoiveType(ctx context.Context, req *MovieTypes) error
	DeleteMovieTypeById(ctx context.Context, id int64) error
	FindAllMovie(ctx context.Context) ([]*MovieTypes, error)
	GetMovieTypeByName(ctx context.Context, name string) (*MovieTypes, error)
}
