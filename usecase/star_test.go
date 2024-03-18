package usecase_test

import (
	"context"
	"filmLibrary/model"
	"filmLibrary/usecase"
	"testing"

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

func TestAddStar(t *testing.T) {
	mockStorage := new(MockStarUsecase)
	usecase := usecase.NewStarUsecase(mockStorage)

	testStar := model.DBStar{Name: "Test Movie", Birthday: "0001-01-01"}
	mockStorage.On("AddStarWithMovies", mock.Anything, testStar).Return(1, nil)

	id, err := usecase.AddStar(context.Background(), *testStar.ToAPI())

	assert.NoError(t, err)
	assert.Equal(t, 1, id)
	mockStorage.AssertExpectations(t)
}

func TestGetAllStars(t *testing.T) {
	mockStorage := new(MockStarUsecase)
	usecase := usecase.NewStarUsecase(mockStorage)

	testStars := []model.APIStar{{Name: "Test star"}}
	mockStorage.On("GetAllStars", mock.Anything).Return(testStars, nil)

	stars, err := usecase.GetAllStars(context.Background())

	assert.NoError(t, err)
	assert.Len(t, stars, 1)
	assert.Equal(t, "Test star", stars[0].Name)
	mockStorage.AssertExpectations(t)
}

func TestUpdateStar(t *testing.T) {
	mockStorage := new(MockStarUsecase)
	usecase := usecase.NewStarUsecase(mockStorage)

	updateData := map[string]interface{}{"Name": "Updated Movie"}
	mockStorage.On("UpdateStar", mock.Anything, 1, updateData).Return(nil)

	err := usecase.UpdateStar(context.Background(), 1, updateData)

	assert.NoError(t, err)
	mockStorage.AssertExpectations(t)
}

func TestDeleteStar(t *testing.T) {
	mockStorage := new(MockStarUsecase)
	usecase := usecase.NewStarUsecase(mockStorage)

	mockStorage.On("DeleteStar", mock.Anything, 1).Return(nil)

	err := usecase.DeleteStar(context.Background(), 1)

	assert.NoError(t, err)
	mockStorage.AssertExpectations(t)
}
