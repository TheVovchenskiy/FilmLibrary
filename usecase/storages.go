package usecase

import (
	"context"
	"filmLibrary/model"
	"filmLibrary/pkg/sortOptions"
)

type MovieStorage interface {
	GetMovies(ctx context.Context, queryOptions map[sortOptions.SortOptionName]sortOptions.SortOptionValue) ([]model.DBMovie, error)
}

type StarStorage interface {
	GetMoviesStars(ctx context.Context, movieIds []int) ([]model.DBStar, error)
}
