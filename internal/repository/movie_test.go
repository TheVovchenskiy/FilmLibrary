package repository_test

import (
	"filmLibrary/internal/repository"
	"filmLibrary/model"
	"filmLibrary/pkg/sortOptions"
	"filmLibrary/pkg/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

// func TestGetMovies(t *testing.T) {
// 	tests := []struct {
// 		expected    []model.DBMovie
// 		expectedErr error
// 	}{
// 		{
// 			[]model.DBMovie{model.DBMovie{Id: 1, Name: "name", Descriprion: "description", ReleaseDate: "2000-01-01", Rating: 10}},
// 			nil,
// 		},
// 	}
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	repo := repository.NewMoviesPg(db)

// 	for _, tt := range tests {
// 		rows := sqlmock.NewRows([]string{
// 			"id",
// 			`"name`,
// 			"description",
// 			"release_date",
// 			"rating",
// 		})

// 		for _, item := range tt.expected {
// 			rows = rows.AddRow(
// 				item.Id,
// 				item.Name,
// 				item.Descriprion,
// 				item.ReleaseDate,
// 				item.Rating,
// 			)
// 		}
// 		mock.
// 			ExpectQuery("SELECT(.|\n)+FROM(.|\n)+").
// 			WillReturnRows(rows)

// 		actual, err := repo.GetMovies(utils.СtxWithLogger, map[sortOptions.SortOptionName]sortOptions.SortOptionValue{})
// 		if err != nil {
// 			t.Errorf("unexpected err: %s", err)
// 			return
// 		}
// 		if err := mock.ExpectationsWereMet(); err != nil {
// 			t.Errorf("there were unfulfilled expectations: %s", err)
// 			return
// 		}
// 		if !reflect.DeepEqual(actual, tt.expected) {
// 			t.Errorf("results not match, want %v, have %v", tt.expected, actual)
// 			return
// 		}
// 	}

// }

// func TestAddMovie(t *testing.T) {
// 	tests := []struct {
// 		movie       model.DBMovie
// 		expectedId  int
// 		expectedErr error
// 	}{
// 		{
// 			model.DBMovie{Id: 1, Name: "name", Descriprion: "description", ReleaseDate: "2000-01-01", Rating: 10},
// 			1,
// 			nil,
// 		},
// 	}
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	repo := repository.NewMoviesPg(db)

// 	for _, tt := range tests {
// 		rows := sqlmock.NewRows([]string{
// 			"id",
// 		}).
// 			AddRow(tt.expectedId)

// 		mock.
// 			ExpectQuery(utils.InsertQuery).
// 			WithArgs(
// 				tt.movie.Name,
// 				tt.movie.Descriprion,
// 				tt.movie.ReleaseDate,
// 				tt.movie.Rating,
// 			).
// 			WillReturnRows(rows)

// 		actual, err := repo.AddMovie(utils.СtxWithLogger, tt.movie)
// 		if err != nil {
// 			t.Errorf("unexpected err: %s", err)
// 			return
// 		}
// 		if err := mock.ExpectationsWereMet(); err != nil {
// 			t.Errorf("there were unfulfilled expectations: %s", err)
// 			return
// 		}
// 		if actual != tt.expectedId {
// 			t.Errorf("results not match, want %v, have %v", tt.expectedId, actual)
// 			return
// 		}
// 	}

// }

func TestGetMovies(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	repo := repository.NewMoviesPg(db)

	rows := sqlmock.NewRows([]string{"id", "name", "description", "release_date", "rating"}).
		AddRow(1, "Movie 1", "Description 1", "2022-01-01", 5).
		AddRow(2, "Movie 2", "Description 2", "2022-02-01", 4)

	mock.ExpectQuery("SELECT(.| )+").
		WillReturnRows(rows)

	queryOptions := map[sortOptions.SortOptionName]sortOptions.SortOptionValue{
		sortOptions.SortFiled: "name",
		sortOptions.SortOrder: sortOptions.AscendingSQL,
	}

	movies, err := repo.GetMovies(utils.СtxWithLogger, queryOptions)
	assert.NoError(t, err)
	assert.Len(t, movies, 2)
}

func TestAddMovie(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	repo := repository.NewMoviesPg(db)

	mock.ExpectQuery("INSERT(.| )+").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	// ctx := context.Background()
	dbMovie := model.DBMovie{
		Name:        "New Movie",
		Descriprion: "New Description",
		ReleaseDate: "2022-03-01",
		Rating:      4,
	}

	insertedMovieId, err := repo.AddMovie(utils.СtxWithLogger, dbMovie)
	assert.NoError(t, err)
	assert.Equal(t, 1, insertedMovieId)
}

func TestUpdateMovie(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	repo := repository.NewMoviesPg(db)

	mock.ExpectExec("UPDATE (.| )+").
		WillReturnResult(sqlmock.NewResult(0, 1))

	movieId := 1
	updateData := map[string]interface{}{
		"name": "Updated Movie",
		"description": "Updated Description",
		"release_date": "2022-04-01",
		"rating": 5,
	}

	err = repo.UpdateMovie(utils.СtxWithLogger, movieId, updateData)
	assert.NoError(t, err)
}

func TestDeleteMovie(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	repo := repository.NewMoviesPg(db)

	mock.ExpectExec("DELETE (.| )+").
		WillReturnResult(sqlmock.NewResult(0, 1))

	movieId := 1

	err = repo.DeleteMovie(utils.СtxWithLogger, movieId)
	assert.NoError(t, err)
}

func TestSearchMovie(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	repo := repository.NewMoviesPg(db)

	rows := sqlmock.NewRows([]string{"id", "name", "description", "release_date", "rating"}).
		AddRow(1, "Search Movie 1", "Description 1", "2022-01-01", 4).
		AddRow(2, "Search Movie 2", "Description 2", "2022-02-01", 3)

	mock.ExpectQuery("SELECT (.| )+").
		WillReturnRows(rows)

	searchQuery := "Search"

	movies, err := repo.SearchMovie(utils.СtxWithLogger, searchQuery)
	assert.NoError(t, err)
	assert.Len(t, movies, 2)
}
