package rest_test

import (
	"bytes"
	"context"
	"encoding/json"
	"filmLibrary/internal/rest"
	"filmLibrary/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockStarUsecase struct {
	mock.Mock
}

func (m *MockStarUsecase) GetAllStars(ctx context.Context) ([]model.APIStar, error) {
	args := m.Called(ctx)
	return args.Get(0).([]model.APIStar), args.Error(1)
}

func (m *MockStarUsecase) AddStarWithMovies(ctx context.Context, star model.DBStar) (int, error) {
	args := m.Called(ctx, star)
	return args.Int(0), args.Error(1)
}
func (m *MockStarUsecase) UpdateStar(ctx context.Context, starId int, updateData map[string]interface{}) error {
	args := m.Called(ctx, starId, updateData)
	return args.Error(0)
}

func (m *MockStarUsecase) DeleteStar(ctx context.Context, starId int) error {
	args := m.Called(ctx, starId)
	return args.Error(0)
}

func TestHandleStarPost(t *testing.T) {
	mockUsecase := new(MockStarUsecase)
	handler := rest.NewStarHandler(mockUsecase)

	testStar := model.DBStar{Name: "Test Movie", Birthday: "0001-01-01"}
	mockUsecase.On("AddStarWithMovies", mock.Anything, testStar).Return(1, nil)

	body, _ := json.Marshal(testStar.ToAPI())
	req, _ := http.NewRequest("POST", "/stars", bytes.NewBuffer(body))
	resp := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/stars", handler.HandleStars)
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockUsecase.AssertExpectations(t)
}

func TestHandleStarsGet(t *testing.T) {
	mockUsecase := new(MockStarUsecase)
	handler := rest.NewStarHandler(mockUsecase)

	testStars := []model.APIStar{{Name: "Test Movie", Birthday: "0001-01-01"}}
	mockUsecase.On("GetAllStars", mock.Anything, mock.Anything).Return(testStars, nil)

	req, _ := http.NewRequest("GET", "/stars", nil)
	resp := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/stars", handler.HandleStars)
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockUsecase.AssertExpectations(t)
}

func TestHandleStarsPut(t *testing.T) {
	mockUsecase := new(MockStarUsecase)
	handler := rest.NewStarHandler(mockUsecase)

	updateData := map[string]interface{}{"name": "updated star"}
	mockUsecase.On("UpdateStar", mock.Anything, 1, updateData).Return(nil)

	body, _ := json.Marshal(updateData)
	req, _ := http.NewRequest("PUT", "/stars/1", bytes.NewBuffer(body))
	resp := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/stars/{id}", handler.HandleStar)
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockUsecase.AssertExpectations(t)
}

func TestHandleStarDelete(t *testing.T) {
	mockUsecase := new(MockStarUsecase)
	handler := rest.NewStarHandler(mockUsecase)

	mockUsecase.On("DeleteStar", mock.Anything, 1).Return(nil)

	req, _ := http.NewRequest("DELETE", "/star/1", nil)
	resp := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/star/{id}", handler.HandleStar)
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockUsecase.AssertExpectations(t)
}
