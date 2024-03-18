package rest_test

import (
	"bytes"
	"context"
	"encoding/json"
	"filmLibrary/internal/rest"
	"filmLibrary/model"
	"filmLibrary/pkg/sortOptions"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockMovieUsecase struct {
	mock.Mock
}

func (m *MockMovieUsecase) AddMovie(ctx context.Context, apiMovie model.DBMovie) (int, error) {
	args := m.Called(ctx, apiMovie)
	return args.Int(0), args.Error(1)
}

func (m *MockMovieUsecase) DeleteMovie(ctx context.Context, movieId int) error {
	args := m.Called(ctx, movieId)
	return args.Error(0)
}
func (m *MockMovieUsecase) GetMovies(ctx context.Context, queryOptions map[sortOptions.SortOptionName]sortOptions.SortOptionValue) ([]model.DBMovie, error) {
	args := m.Called(ctx, queryOptions)
	return args.Get(0).([]model.DBMovie), args.Error(1)
}

func (m *MockMovieUsecase) SearchMovie(ctx context.Context, searchQuery string) ([]model.DBMovie, error) {
	args := m.Called(ctx, searchQuery)
	return args.Get(0).([]model.DBMovie), args.Error(1)
}

func (m *MockMovieUsecase) UpdateMovie(ctx context.Context, movieId int, updateData map[string]interface{}) error {
	args := m.Called(ctx, movieId, updateData)
	return args.Error(0)
}

func TestHandleMoviesPost(t *testing.T) {
	mockUsecase := new(MockMovieUsecase)
	handler := rest.NewMovieHandler(mockUsecase)

	testMovie := model.DBMovie{Name: "Test Movie", ReleaseDate: "0001-01-01"}
	mockUsecase.On("AddMovie", mock.Anything, testMovie).Return(1, nil)

	body, _ := json.Marshal(testMovie)
	req, _ := http.NewRequest("POST", "/movies", bytes.NewBuffer(body))
	resp := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/movies", handler.HandleMovies)
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockUsecase.AssertExpectations(t)
}

func TestHandleMoviesGet(t *testing.T) {
	mockUsecase := new(MockMovieUsecase)
	handler := rest.NewMovieHandler(mockUsecase)

	testMovies := []model.DBMovie{{Name: "Test Movie", ReleaseDate: "0001-01-01"}}
	mockUsecase.On("GetMovies", mock.Anything, mock.Anything).Return(testMovies, nil)

	req, _ := http.NewRequest("GET", "/movies", nil)
	resp := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/movies", handler.HandleMovies)
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockUsecase.AssertExpectations(t)
}

func TestHandleMoviePut(t *testing.T) {
	mockUsecase := new(MockMovieUsecase)
	handler := rest.NewMovieHandler(mockUsecase)

	updateData := map[string]interface{}{"Name": "Updated Movie"}
	mockUsecase.On("UpdateMovie", mock.Anything, 1, updateData).Return(nil)

	body, _ := json.Marshal(updateData)
	req, _ := http.NewRequest("PUT", "/movie/1", bytes.NewBuffer(body))
	resp := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/movie/{id}", handler.HandleMovie)
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockUsecase.AssertExpectations(t)
}

func TestHandleMovieDelete(t *testing.T) {
	mockUsecase := new(MockMovieUsecase)
	handler := rest.NewMovieHandler(mockUsecase)

	mockUsecase.On("DeleteMovie", mock.Anything, 1).Return(nil)

	req, _ := http.NewRequest("DELETE", "/movie/1", nil)
	resp := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/movie/{id}", handler.HandleMovie)
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockUsecase.AssertExpectations(t)
}

func TestHandleSearchMovies(t *testing.T) {
	mockUsecase := new(MockMovieUsecase)
	handler := rest.NewMovieHandler(mockUsecase)

	testMovies := []model.DBMovie{{Name: "Test Movie", ReleaseDate: "0001-01-01"}}
	searchQuery := "Test"
	mockUsecase.On("SearchMovie", mock.Anything, searchQuery).Return(testMovies, nil)

	req, _ := http.NewRequest("GET", "/movies/search?q="+searchQuery, nil)
	resp := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/movies/search", handler.HandleSearchMovies)
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockUsecase.AssertExpectations(t)
}
