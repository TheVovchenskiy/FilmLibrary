package routers

import (
	"filmLibrary/internal/rest"
	"filmLibrary/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

func MountStarRouter(router *mux.Router, starStorage usecase.StarStorage) {
	handler := rest.NewStarHandler(starStorage)

	router.Handle("/stars", http.HandlerFunc(handler.HandleStars)).Methods("GET", "POST")
	router.Handle("/stars/{id}", http.HandlerFunc(handler.HandleStar)).Methods("PUT", "DELETE")
}
