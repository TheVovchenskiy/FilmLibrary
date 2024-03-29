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
	SearchMovie(ctx context.Context, searchQuery string) ([]model.DBMovie, error)
	// GetMoviesByStarsIds(ctx context.Context, starsIds []int) ([]model.DBMovie, error)
}

type StarStorage interface {
	GetAllStars(ctx context.Context) ([]model.APIStar, error)
	AddStarWithMovies(ctx context.Context, star model.DBStar) (int, error)
	UpdateStar(ctx context.Context, starId int, updateData map[string]interface{}) error
	DeleteStar(ctx context.Context, starId int) error
	// GetMoviesStars(ctx context.Context, movieIds []int) ([]model.DBStar, error)
}
