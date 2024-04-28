package usecase

import (
	"context"
	"flick_tickets/common/enums"
	"flick_tickets/common/log"
	"flick_tickets/common/utils"
	"flick_tickets/core/domain"
	"flick_tickets/core/entities"
)

type UseCaseMovie struct {
	movie domain.RepositoryMovieType
}

func NewUseCaseMovie(
	movie domain.RepositoryMovieType,
) *UseCaseMovie {
	return &UseCaseMovie{
		movie: movie,
	}
}
func (u *UseCaseMovie) AddMovieType(ctx context.Context, req *entities.MovieTypesAddReq) (*entities.MovieTypesAddResp, error) {

	movieType, err := u.movie.GetMovieTypeByName(ctx, req.MovieTypeName)
	if err != nil {
		log.Error(err, "errpr 1")
		return &entities.MovieTypesAddResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	if movieType != nil {
		return &entities.MovieTypesAddResp{
			Result: entities.Result{
				Code:    enums.MOVIE_EXIST_CODE,
				Message: enums.MOVIE_EXIST_MESS,
			},
		}, nil
	}
	err = u.movie.AddMoiveType(ctx, &domain.MovieTypes{
		Id:            utils.GenerateUniqueKey(),
		MovieTypeName: req.MovieTypeName,
	})
	if err != nil {
		log.Error(err, "errpr 2")
		return &entities.MovieTypesAddResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}

	return &entities.MovieTypesAddResp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
	}, nil
}

func (u *UseCaseMovie) GetAllMovieType(ctx context.Context) (*entities.MovieGetAllResp, error) {
	movies, err := u.movie.FindAllMovie(ctx)
	if err != nil {
		return &entities.MovieGetAllResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	if len(movies) == 0 {
		return &entities.MovieGetAllResp{
			Result: entities.Result{
				Code:    enums.DATA_EMPTY_ERR_CODE,
				Message: enums.DATA_EMPTY_ERR_MESS,
			},
		}, nil
	}
	return &entities.MovieGetAllResp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Movie: movies,
	}, nil
}
