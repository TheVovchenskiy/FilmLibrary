package app

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"os"
)

func GetPostgres() (*sql.DB, error) {
	pgConnStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", os.Getenv("PG_USER"), os.Getenv("PG_DBNAME"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_HOST"), os.Getenv("PG_PORT"))
	conn, err := sql.Open("pgx", pgConnStr)
	if err != nil {
		return nil, err
	}
	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	return conn, nil
}
