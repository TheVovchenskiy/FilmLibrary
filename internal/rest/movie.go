package rest

import (
	"filmLibrary/pkg/responseTemplate"
	"filmLibrary/pkg/serverErrors"
	"filmLibrary/usecase"
	"net/http"
)

type MovieHandler struct {
	usecase *usecase.MovieUsecase
}

func NewMovieHandler(movieStorage *usecase.MovieStorage) *MovieHandler {
	return &MovieHandler{
		usecase: usecase.NewMovieUsecase(movieStorage),
	}
}

func (api *MovieHandler) HandleGetMovies(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		responseTemplate.ServeJsonError(w, serverErrors.ErrMethodNotAllowed)
	}
}
