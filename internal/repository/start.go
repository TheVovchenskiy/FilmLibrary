package repository

import (
	"context"
	"database/sql"
	"filmLibrary/model"
)

type StarsPg struct {
	db *sql.DB
}

func NewStarsPg(db *sql.DB) *StarsPg {
	return &StarsPg{
		db: db,
	}
}

func (repo *StarsPg) GetMoviesStars(ctx context.Context, movieIds []int) ([]model.DBStar, error) {
	return nil, nil
}
