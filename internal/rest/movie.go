package rest

import (
	"context"
	"encoding/json"
	"filmLibrary/model"
	"filmLibrary/pkg/responseTemplate"
	"filmLibrary/pkg/sanitizer"
	"filmLibrary/pkg/searchOptions"
	"filmLibrary/pkg/serverErrors"
	"filmLibrary/pkg/sortOptions"
	"filmLibrary/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type MovieHandler struct {
	movieUsecase *usecase.MovieUsecase
}

func NewMovieHandler(movieStorage usecase.MovieStorage) *MovieHandler {
	return &MovieHandler{
		movieUsecase: usecase.NewMovieUsecase(movieStorage),
	}
}

func sanitizeMovie(movie model.APIMovie) model.APIMovie {
	movie.Name = sanitizer.XSS.Sanitize(movie.Name)
	movie.Descriprion = sanitizer.XSS.Sanitize(movie.Descriprion)
	movie.ReleaseDate = sanitizer.XSS.Sanitize(movie.ReleaseDate)

	// movie.Stars = sanitizeStars(movie.Stars...)

	return movie
}

func sanitizeMovies(movies ...model.APIMovie) (sanitizedMovies []model.APIMovie) {
	sanitizedMovies = make([]model.APIMovie, 0, len(movies))
	for _, movie := range movies {
		sanitizedMovies = append(sanitizedMovies, sanitizeMovie(movie))
	}

	return
}

func (movieHandler *MovieHandler) HandleMovies(w http.ResponseWriter, r *http.Request) {
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
		sanitizedMovies := sanitizeMovies(movies...)
		responseTemplate.MarshalAndSend(w, sanitizedMovies)
	case http.MethodPost:
		defer r.Body.Close()

		apiMovie := model.APIMovie{}
		err := json.NewDecoder(r.Body).Decode(&apiMovie)
		if err != nil {
			responseTemplate.ServeJsonError(w, serverErrors.ErrInvalidRequest)
			return
		}

		newMovieId, err := movieHandler.movieUsecase.AddMovie(context.Background(), apiMovie)
		if err != nil {
			responseTemplate.ServeJsonError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprintf(`{"id":%d}`, newMovieId)))
	default:
		responseTemplate.ServeJsonError(w, serverErrors.ErrMethodNotAllowed)
	}
}

func (movieHandler *MovieHandler) HandleMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieId, err := strconv.Atoi(vars["id"])

	if err != nil {
		responseTemplate.ServeJsonError(w, serverErrors.ErrInvalidRequest)
		return
	}
	switch r.Method {
	case http.MethodPut:
		updateData := make(map[string]interface{})
		if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
			responseTemplate.ServeJsonError(w, err)
			return
		}

		err = movieHandler.movieUsecase.UpdateMovie(context.Background(), movieId, updateData)
		if err != nil {
			responseTemplate.ServeJsonError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)

	case http.MethodDelete:
		err = movieHandler.movieUsecase.DeleteMovie(context.Background(), movieId)
		if err != nil {
			responseTemplate.ServeJsonError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	default:
		responseTemplate.ServeJsonError(w, serverErrors.ErrMethodNotAllowed)
	}
}

func (handler *MovieHandler) HandleSearchMovies(w http.ResponseWriter, r *http.Request) {
	searchQuery, err := searchOptions.GetSearchQuery(r.URL.Query())
	if err != nil {
		responseTemplate.ServeJsonError(w, err)
		return
	}

	movies, err := handler.movieUsecase.SearchMovies(context.Background(), searchQuery)
	if err != nil {
		responseTemplate.ServeJsonError(w, err)
		return
	}
	sanitizedMovies := sanitizeMovies(movies...)
	responseTemplate.MarshalAndSend(w, sanitizedMovies)
}
