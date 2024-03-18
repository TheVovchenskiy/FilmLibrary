package repository

import (
	"context"
	"database/sql"
	"filmLibrary/model"
	"filmLibrary/pkg/sortOptions"
	"fmt"
)

type MoviesPg struct {
	db *sql.DB
}

func NewMoviesPg(db *sql.DB) *MoviesPg {
	return &MoviesPg{
		db: db,
	}
}

func (repo *MoviesPg) GetMovies(ctx context.Context, queryOptions map[sortOptions.SortOptionName]sortOptions.SortOptionValue) ([]model.DBMovie, error) {
	query := fmt.Sprintf(`SELECT
							id,
							"name",
							description,
							release_date,
							rating
						FROM
							public.movie m
						ORDER BY
							%s %s`,
		queryOptions[sortOptions.SortFiled],
		sortOptions.MapSortOrderSQL(queryOptions[sortOptions.SortOrder]),
	)
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	moviesToReturn := []model.DBMovie{}

	for rows.Next() {
		movie := model.DBMovie{}
		err := rows.Scan(
			&movie.Id,
			&movie.Name,
			&movie.Descriprion,
			&movie.ReleaseDate,
			&movie.Rating,
		)
		if err != nil {
			return nil, err
		}
		moviesToReturn = append(moviesToReturn, movie)
	}

	return moviesToReturn, nil
}
