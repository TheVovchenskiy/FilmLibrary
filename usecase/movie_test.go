package usecase_test

import (
	"context"
	"filmLibrary/model"
	"filmLibrary/pkg/sortOptions"
	"filmLibrary/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockMovieStorage struct {
	mock.Mock
}

func (m *MockMovieStorage) GetMovies(ctx context.Context, queryOptions map[sortOptions.SortOptionName]sortOptions.SortOptionValue) ([]model.DBMovie, error) {
	args := m.Called(ctx, queryOptions)
	return args.Get(0).([]model.DBMovie), args.Error(1)
}

func (m *MockMovieStorage) AddMovie(ctx context.Context, dbMovie model.DBMovie) (int, error) {
	args := m.Called(ctx, dbMovie)
	return args.Int(0), args.Error(1)
}

func (m *MockMovieStorage) UpdateMovie(ctx context.Context, movieId int, updateData map[string]interface{}) error {
	args := m.Called(ctx, movieId, updateData)
	return args.Error(0)
}

func (m *MockMovieStorage) DeleteMovie(ctx context.Context, movieId int) error {
	args := m.Called(ctx, movieId)
	return args.Error(0)
}

func (m *MockMovieStorage) SearchMovie(ctx context.Context, searchQuery string) ([]model.DBMovie, error) {
	args := m.Called(ctx, searchQuery)
	return args.Get(0).([]model.DBMovie), args.Error(1)
}

func TestAddMovie(t *testing.T) {
	mockStorage := new(MockMovieStorage)
	usecase := usecase.NewMovieUsecase(mockStorage)

	testMovie := model.DBMovie{Name: "Test Movie", ReleaseDate: "0001-01-01"}
	mockStorage.On("AddMovie", mock.Anything, testMovie).Return(1, nil)

	id, err := usecase.AddMovie(context.Background(), *testMovie.ToAPI())

	assert.NoError(t, err)
	assert.Equal(t, 1, id)
	mockStorage.AssertExpectations(t)
}

func TestGetAllMovies(t *testing.T) {
	mockStorage := new(MockMovieStorage)
	usecase := usecase.NewMovieUsecase(mockStorage)

	testMovies := []model.DBMovie{{Name: "Test Movie"}}
	mockStorage.On("GetMovies", mock.Anything, mock.Anything).Return(testMovies, nil)

	queryOptions := make(map[sortOptions.SortOptionName]sortOptions.SortOptionValue)
	movies, err := usecase.GetAllMovies(context.Background(), queryOptions)

	assert.NoError(t, err)
	assert.Len(t, movies, 1)
	assert.Equal(t, "Test Movie", movies[0].Name)
	mockStorage.AssertExpectations(t)
}

func TestUpdateMovie(t *testing.T) {
	mockStorage := new(MockMovieStorage)
	usecase := usecase.NewMovieUsecase(mockStorage)

	updateData := map[string]interface{}{"Name": "Updated Movie"}
	mockStorage.On("UpdateMovie", mock.Anything, 1, updateData).Return(nil)

	err := usecase.UpdateMovie(context.Background(), 1, updateData)

	assert.NoError(t, err)
	mockStorage.AssertExpectations(t)
}

func TestDeleteMovie(t *testing.T) {
	mockStorage := new(MockMovieStorage)
	usecase := usecase.NewMovieUsecase(mockStorage)

	mockStorage.On("DeleteMovie", mock.Anything, 1).Return(nil)

	err := usecase.DeleteMovie(context.Background(), 1)

	assert.NoError(t, err)
	mockStorage.AssertExpectations(t)
}
func TestSearchMovies(t *testing.T) {
	mockStorage := new(MockMovieStorage)
	usecase := usecase.NewMovieUsecase(mockStorage)

	testMovies := []model.DBMovie{{Name: "Test Movie"}}
	searchQuery := "Test"
	mockStorage.On("SearchMovie", mock.Anything, searchQuery).Return(testMovies, nil)

	movies, err := usecase.SearchMovies(context.Background(), searchQuery)

	assert.NoError(t, err)
	assert.Len(t, movies, 1)
	assert.Equal(t, "Test Movie", movies[0].Name)
	mockStorage.AssertExpectations(t)
}
