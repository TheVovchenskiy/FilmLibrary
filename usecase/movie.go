package usecase

import (
	"context"
	"filmLibrary/model"
)

type MovieStorage interface {
	GetMovies(ctx context.Context) (model.Movie, error)
}

type MovieUsecase struct {
	movieStorage *MovieStorage
}

func NewMovieUsecase(movieStorage *MovieStorage) *MovieUsecase {
	return &MovieUsecase{
		movieStorage: movieStorage,
	}
}

