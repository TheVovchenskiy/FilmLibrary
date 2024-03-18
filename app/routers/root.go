package routers

import (
	"filmLibrary/app"
	"filmLibrary/configs"
	"filmLibrary/internal/repository"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

func Run() (err error) {
	db, err := app.GetPostgres()
	if err != nil {
		return
	}
	defer db.Close()

	movieStorage := repository.NewMoviesPg(db)

	// rootRouter := http.NewServeMux()
	rootRouter := mux.NewRouter()

	MountMovieRouter(rootRouter, movieStorage)

	fmt.Printf("\tstarting server at %d\n", configs.PORT)
	err = http.ListenAndServe(fmt.Sprintf(":%d", configs.PORT), rootRouter)

	return
}
