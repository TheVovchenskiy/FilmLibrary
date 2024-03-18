package repository_test

import (
	"filmLibrary/internal/repository"
	"filmLibrary/model"
	"filmLibrary/pkg/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)


func TestGetAllStars(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	repo := repository.NewStarsPg(db)

	rows := sqlmock.NewRows([]string{"id", "name", "gender", "birthday", "movie"}).
		AddRow(1, "Star 1", "Male", "1990-01-01", "Movie 1").
		AddRow(1, "Star 1", "Male", "1990-01-01", "Movie 2")

	mock.ExpectQuery("SELECT s.id, s.\"name\", s.gender, s.birthday, m.\"name\" FROM public.star s JOIN public.movie_star_assign msa ON s.id = msa.star_id JOIN public.movie m ON msa.movie_id = m.id").
		WillReturnRows(rows)


	stars, err := repo.GetAllStars(utils.СtxWithLogger)
	assert.NoError(t, err)
	assert.Len(t, stars, 1)
	assert.Len(t, stars[0].Movies, 2)
}

func TestAddStarWithMovies(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	repo := repository.NewStarsPg(db)

	mock.ExpectQuery("INSERT INTO public.star (.+) VALUES (.+) RETURNING id").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	star := model.DBStar{
		Name:     "New Star",
		Gender:   "Female",
		Birthday: "1995-05-05",
	}

	insertedStarId, err := repo.AddStarWithMovies(utils.СtxWithLogger, star)
	assert.NoError(t, err)
	assert.Equal(t, 1, insertedStarId)
}

func TestUpdateStar(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	repo := repository.NewStarsPg(db)

	mock.ExpectExec("UPDATE public.star SET (.+) WHERE id = (.+)").
		WillReturnResult(sqlmock.NewResult(0, 1))

	starId := 1
	updateData := map[string]interface{}{
		"name":     "Updated Star",
		"gender":   "Male",
		"birthday": "1990-01-01",
	}

	err = repo.UpdateStar(utils.СtxWithLogger, starId, updateData)
	assert.NoError(t, err)
}

func TestDeleteStar(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	repo := repository.NewStarsPg(db)

	mock.ExpectExec("DELETE FROM public.star m WHERE m.id = (.+)").
		WillReturnResult(sqlmock.NewResult(0, 1))

	starId := 1

	err = repo.DeleteStar(utils.СtxWithLogger, starId)
	assert.NoError(t, err)
}
