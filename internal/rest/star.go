package rest

import (
	"context"
	"encoding/json"
	"filmLibrary/model"
	"filmLibrary/pkg/responseTemplate"
	"filmLibrary/pkg/sanitizer"
	"filmLibrary/pkg/serverErrors"
	"filmLibrary/pkg/utils"
	"filmLibrary/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type StarHandler struct {
	starUsecase *usecase.StarUsecase
}

func NewStarHandler(starStorage usecase.StarStorage) *StarHandler {
	return &StarHandler{
		starUsecase: usecase.NewStarUsecase(starStorage),
	}
}

func sanitizeStar(star model.APIStar) model.APIStar {
	star.Name = sanitizer.XSS.Sanitize(star.Name)
	star.Birthday = sanitizer.XSS.Sanitize(star.Birthday)
	star.Gender = model.Gender(sanitizer.XSS.Sanitize(string(star.Gender)))

	sanitizedMovies := make([]string, 0, len(star.Movies))
	for _, movie := range star.Movies {
		sanitizedMovies = append(sanitizedMovies, sanitizer.XSS.Sanitize(movie))
	}
	star.Movies = sanitizedMovies

	return star
}

func sanitizeStars(stars ...model.APIStar) (sanitizedMovies []model.APIStar) {
	sanitizedMovies = make([]model.APIStar, 0, len(stars))
	for _, star := range stars {
		sanitizedMovies = append(sanitizedMovies, sanitizeStar(star))
	}

	return
}

func (handler *StarHandler) HandleStars(w http.ResponseWriter, r *http.Request) {
	contextLogger := utils.GetContextLogger(r.Context())

	switch r.Method {
	case http.MethodGet:
		stars, err := handler.starUsecase.GetAllStars(r.Context())
		if err != nil {
			responseTemplate.ServeJsonError(w, err)
			return
		}

		sanitizedStars := sanitizeStars(stars...)
		responseTemplate.MarshalAndSend(w, sanitizedStars)
	case http.MethodPost:
		defer r.Body.Close()

		apiStar := model.APIStar{}
		err := json.NewDecoder(r.Body).Decode(&apiStar)
		if err != nil {
			contextLogger.WithFields(logrus.Fields{
				"err_msg": err,
			}).
				Info("json decoder error")
			responseTemplate.ServeJsonError(w, serverErrors.ErrInvalidRequest)
			return
		}

		newMovieId, err := handler.starUsecase.AddStar(r.Context(), apiStar)
		if err != nil {
			responseTemplate.ServeJsonError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprintf(`{"id":%d}`, newMovieId)))
	}
}

func (handler *StarHandler) HandleStar(w http.ResponseWriter, r *http.Request) {
	contextLogger := utils.GetContextLogger(r.Context())
	vars := mux.Vars(r)
	starId, err := strconv.Atoi(vars["id"])

	if err != nil {
		responseTemplate.ServeJsonError(w, serverErrors.ErrInvalidRequest)
		return
	}

	switch r.Method {
	case http.MethodPut:
		updateData := make(map[string]interface{})
		if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
			contextLogger.WithFields(logrus.Fields{
				"err_msg": err,
			}).
				Info("json decoder error")
			responseTemplate.ServeJsonError(w, err)
			return
		}

		err = handler.starUsecase.UpdateStar(r.Context(), starId, updateData)
		if err != nil {
			responseTemplate.ServeJsonError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)

	case http.MethodDelete:
		err = handler.starUsecase.DeleteStar(r.Context(), starId)
		if err != nil {
			responseTemplate.ServeJsonError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	default:
		responseTemplate.ServeJsonError(w, serverErrors.ErrMethodNotAllowed)
	}
}
