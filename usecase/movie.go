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

func (u *MovieUsecase) AddMovie(ctx context.Context, apiMovie model.APIMovie) (int, error) {
	dbMovie := apiMovie.ToDB()
	id, err := u.movieStorage.AddMovie(ctx, *dbMovie)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *MovieUsecase) UpdateMovie(ctx context.Context, movieId int, updateData map[string]interface{}) error {
	return u.movieStorage.UpdateMovie(ctx, movieId, updateData)
}

func (u *MovieUsecase) DeleteMovie(ctx context.Context, movieId int) error {
	return u.movieStorage.DeleteMovie(ctx, movieId)
}
