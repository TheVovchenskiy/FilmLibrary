package repository

import (
	"context"
	"database/sql"
	"filmLibrary/model"
	"filmLibrary/pkg/utils"
	"fmt"
	"strings"
)

type StarsPg struct {
	db     *sql.DB
	fields []string
}

func NewStarsPg(db *sql.DB) *StarsPg {
	return &StarsPg{
		db: db,
		fields: []string{
			"id",
			`"name"`,
			"gender",
			"birthday",
		},
	}
}

func (repo *StarsPg) mapField(field string) string {
	if field == "name" {
		return `"name"`
	}
	return field
}

func (repo *StarsPg) validateField(field string) bool {
	return utils.In(field, repo.fields)
}

func (repo *StarsPg) GetAllStars(ctx context.Context) ([]model.APIStar, error) {
	query := `SELECT
				s.id,
				s."name",
				s.gender,
				s.birthday,
				m."name"
			FROM
				public.star s
			JOIN public.movie_star_assign msa
			ON
				s.id = msa.star_id
			JOIN public.movie m
			ON
				msa.movie_id = m.id`
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	starsMap := make(map[int]*model.APIStar)
	for rows.Next() {
		star := model.APIStar{}
		var movie string
		err := rows.Scan(
			&star.Id,
			&star.Name,
			&star.Gender,
			&star.Birthday,
			&movie,
		)
		if err != nil {
			return nil, err
		}

		if existingStar, exists := starsMap[star.Id]; exists {
			existingStar.Movies = append(existingStar.Movies, movie)
		} else {
			star.Movies = append(star.Movies, movie)
			starsMap[star.Id] = &star
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	stars := make([]model.APIStar, 0, len(starsMap))
	for _, star := range starsMap {
		stars = append(stars, *star)
	}

	return stars, nil
}

func (repo *StarsPg) AddStarWithMovies(ctx context.Context, star model.DBStar) (int, error) {
	query := `INSERT
				INTO
				public.star (
					"name",
					gender,
					birthday
				)
			VALUES ($1, $2, $3)
			RETURNING id`

	var insertedStarId int

	err := repo.db.QueryRow(
		query,
		star.Name,
		star.Gender,
		star.Birthday,
	).Scan(&insertedStarId)

	if err == sql.ErrNoRows {
		return 0, ErrNotInserted
	}

	if err != nil {
		return 0, err
	}

	return insertedStarId, nil
}

func (repo *StarsPg) UpdateStar(ctx context.Context, starId int, updateData map[string]interface{}) error {
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

	args = append(args, starId)

	query := fmt.Sprintf(`UPDATE
							public.star
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

func (repo *StarsPg) DeleteStar(ctx context.Context, starId int) error {
	query := `DELETE
			FROM
				public.star m
			WHERE
				m.id = $1`

	result, err := repo.db.Exec(query, starId)

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
