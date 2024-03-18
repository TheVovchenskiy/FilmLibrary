package routers

import (
	"filmLibrary/internal/rest"
	"filmLibrary/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

func MountMovieRouter(router *mux.Router, movieStorage usecase.MovieStorage) {
	handler := rest.NewMovieHandler(movieStorage)

	router.Handle("/movies", http.HandlerFunc(handler.HandleMovies)).Methods("GET", "POST")
	router.Handle("/movies/search", http.HandlerFunc(handler.HandleSearchMovies)).Methods("GET")
	router.Handle("/movies/{id}", http.HandlerFunc(handler.HandleMovie)).Methods("PUT", "DELETE")
}
