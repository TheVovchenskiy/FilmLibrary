package usecase

import (
	"context"
	"filmLibrary/model"
)

type StarUsecase struct {
	starStorage StarStorage
}

func NewStarUsecase(starStorage StarStorage) *StarUsecase {
	return &StarUsecase{
		starStorage: starStorage,
	}
}

func (u *StarUsecase) GetAllStars(ctx context.Context) ([]model.APIStar, error) {
	stars, err := u.starStorage.GetAllStars(ctx)
	if err != nil {
		return nil, err
	}
	return stars, nil
}

func (u *StarUsecase) AddStar(ctx context.Context, apiStar model.APIStar) (int, error) {
	insertedId, err := u.starStorage.AddStarWithMovies(ctx, *apiStar.ToDB())
	if err != nil {
		return 0, err
	}
	return insertedId, nil
}

func (u *StarUsecase) UpdateStar(ctx context.Context, starId int, updateData map[string]interface{}) error {
	return u.starStorage.UpdateStar(ctx, starId, updateData)
}

func (u *StarUsecase) DeleteStar(ctx context.Context, starId int) error {
	return u.starStorage.DeleteStar(ctx, starId)
}
