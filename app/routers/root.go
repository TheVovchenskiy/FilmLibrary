package routers

import (
	"filmLibrary/app"
	"filmLibrary/configs"
	"filmLibrary/internal/repository"
	"filmLibrary/internal/rest/middleware"
	"filmLibrary/pkg/logging"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

func Run() (err error) {
	db, err := app.GetPostgres()
	if err != nil {
		return
	}
	defer db.Close()

	logFile, err := os.OpenFile(configs.LOGS_DIR+configs.LOGFILE_NAME, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	logging.InitLogger(logFile)

	movieStorage := repository.NewMoviesPg(db)
	starStorage := repository.NewStarsPg(db)

	rootRouter := mux.NewRouter()

	MountMovieRouter(rootRouter, movieStorage)
	MountStarRouter(rootRouter, starStorage)

	loggedRouter := middleware.AccessLogMiddleware(rootRouter)
	requestIdRouter := middleware.RequestID(loggedRouter)
	recoverRouter := middleware.PanicRecoverMiddleware(logging.Logger, requestIdRouter)

	fmt.Printf("\tstarting server at %d\n", configs.PORT)
	err = http.ListenAndServe(fmt.Sprintf(":%d", configs.PORT), recoverRouter)

	return
}
