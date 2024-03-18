package usecase

import (
	"context"
	"filmLibrary/model"
	"filmLibrary/pkg/sortOptions"
)

type MovieStorage interface {
	GetMovies(ctx context.Context, queryOptions map[sortOptions.SortOptionName]sortOptions.SortOptionValue) ([]model.DBMovie, error)
	AddMovie(ctx context.Context, dbMovie model.DBMovie) (int, error)
	UpdateMovie(ctx context.Context, movieId int, updateData map[string]interface{}) error
	DeleteMovie(ctx context.Context, movieId int) error
}

type StarStorage interface {
	GetMoviesStars(ctx context.Context, movieIds []int) ([]model.DBStar, error)
}
