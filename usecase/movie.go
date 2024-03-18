package usecase

import (
	"context"
	"filmLibrary/model"
	"filmLibrary/pkg/sortOptions"
)

type MovieUsecase struct {
	movieStorage MovieStorage
}

func NewMovieUsecase(movieStorage MovieStorage) *MovieUsecase {
	return &MovieUsecase{
		movieStorage: movieStorage,
	}
}

func (u *MovieUsecase) dbMoviesToAPI(movies []model.DBMovie) []model.APIMovie {
	res := make([]model.APIMovie, 0, len(movies))
	for _, movie := range movies {
		res = append(res, *movie.ToAPI())
	}
	return res
}

func (u *MovieUsecase) GetAllMovies(
	ctx context.Context,
	queryOptions map[sortOptions.SortOptionName]sortOptions.SortOptionValue,
) ([]model.APIMovie, error) {
	dbMovies, err := u.movieStorage.GetMovies(ctx, queryOptions)
	if err != nil {
		return nil, err
	}

	return u.dbMoviesToAPI(dbMovies), nil
}
