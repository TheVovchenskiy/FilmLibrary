package routers

import (
	"filmLibrary/internal/rest"
	"filmLibrary/usecase"
	"net/http"
)

func MountMovieRouter(router *http.ServeMux, movieStorage *usecase.MovieStorage) {
	handler := rest.NewMovieHandler(movieStorage)

	router.HandleFunc("/movies", handler.HandleGetMovies)
}
