package repository

import (
	"context"
	"database/sql"
	"filmLibrary/model"
	"filmLibrary/pkg/sortOptions"
	"filmLibrary/pkg/utils"
	"fmt"
	"strings"
)

type MoviesPg struct {
	db     *sql.DB
	fields []string
}

func NewMoviesPg(db *sql.DB) *MoviesPg {
	return &MoviesPg{
		db: db,
		fields: []string{
			"id",
			`"name"`,
			"description",
			"release_date",
			"rating",
		},
	}
}

func (repo *MoviesPg) mapField(field string) string {
	if field == "name" {
		return `"name"`
	}
	return field
}

func (repo *MoviesPg) validateField(field string) bool {
	return utils.In(field, repo.fields)
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

func (repo *MoviesPg) AddMovie(ctx context.Context, dbMovie model.DBMovie) (int, error) {
	query := `INSERT
				INTO
				public.movie (
					"name",
					description,
					release_date,
					rating
				)
			VALUES ($1, $2, $3, $4)
			RETURNING id`

	var insertedMovieId int

	err := repo.db.QueryRow(
		query,
		dbMovie.Name,
		dbMovie.Descriprion,
		dbMovie.ReleaseDate,
		dbMovie.Rating,
	).Scan(&insertedMovieId)

	if err == sql.ErrNoRows {
		return 0, ErrNotInserted
	}

	if err != nil {
		return 0, err
	}

	return insertedMovieId, nil
}

func (repo *MoviesPg) UpdateMovie(ctx context.Context, movieId int, updateData map[string]interface{}) error {
	setParts := []string{}
	args := []interface{}{}
	argId := 1

	for key, val := range updateData {
		key = repo.mapField(key)
		if !repo.validateField(key) {
			return ErrInvalidFieldName
		}
		setParts = append(setParts, fmt.Sprintf("%s = $%d", key, argId))
		args = append(args, val)
		argId++
	}

	args = append(args, movieId)

	query := fmt.Sprintf(`UPDATE
							public.movie
						SET
							%s
						WHERE
							id = $%d`,
		strings.Join(setParts, ", "),
		argId,
	)

	result, err := repo.db.Exec(
		query,
		args...,
	)
	if err == sql.ErrNoRows {
		return ErrNoRowsUpdated
	}
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (repo *MoviesPg) DeleteMovie(ctx context.Context, movieId int) error {
	query := `DELETE
			FROM
				public.movie m
			WHERE
				m.id = $1`

	result, err := repo.db.Exec(query, movieId)

	if err == sql.ErrNoRows {
		return ErrNoRowsDeleted
	}
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (repo *MoviesPg) SearchMovie(ctx context.Context, searchQuery string) ([]model.DBMovie, error) {
	likeSearchQuery := "%" + searchQuery + "%"
	query := `SELECT
				m.id,
				m."name",
				m.description,
				m.release_date,
				m.rating
			FROM
				public.movie m
			LEFT JOIN public.movie_star_assign msa
			ON
				m.id = msa.movie_id
			LEFT JOIN public.star s
			ON
				s.id = msa.star_id
			WHERE
				lower(m."name") LIKE lower($1)
				OR lower(s."name") LIKE lower($2)`

	rows, err := repo.db.Query(query, likeSearchQuery, likeSearchQuery)
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
