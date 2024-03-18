package rest

import (
	"context"
	"filmLibrary/model"
	"filmLibrary/pkg/responseTemplate"
	"filmLibrary/pkg/sanitizer"
	"filmLibrary/pkg/serverErrors"
	"filmLibrary/pkg/sortOptions"
	"filmLibrary/usecase"
	"net/http"
)

type MovieHandler struct {
	movieUsecase *usecase.MovieUsecase
}

func NewMovieHandler(movieStorage usecase.MovieStorage) *MovieHandler {
	return &MovieHandler{
		movieUsecase: usecase.NewMovieUsecase(movieStorage),
	}
}

func (movieHandler *MovieHandler) sanitizeMovies(movies ...model.APIMovie) (sanitizedMovies []model.APIMovie) {
	sanitizedMovies = make([]model.APIMovie, 0, len(movies))
	for _, movie := range movies {
		movie.Name = sanitizer.XSS.Sanitize(movie.Name)
		movie.Descriprion = sanitizer.XSS.Sanitize(movie.Descriprion)
		movie.ReleaseDate = sanitizer.XSS.Sanitize(movie.ReleaseDate)

		sanitizedMovies = append(sanitizedMovies, movie)
	}

	return
}

func (movieHandler *MovieHandler) HandleGetMovies(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		queryOptions, err := sortOptions.GetSortOptions(r.URL.Query())
		if err != nil {
			responseTemplate.ServeJsonError(w, err)
			return
		}

		movies, err := movieHandler.movieUsecase.GetAllMovies(context.Background(), queryOptions)
		if err != nil {
			responseTemplate.ServeJsonError(w, err)
			return
		}
		sanitizedMovies := movieHandler.sanitizeMovies(movies...)
		responseTemplate.MarshalAndSend(w, sanitizedMovies)
	default:
		responseTemplate.ServeJsonError(w, serverErrors.ErrMethodNotAllowed)
	}
}
