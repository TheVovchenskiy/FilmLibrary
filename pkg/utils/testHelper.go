package utils

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	ErrQuery = fmt.Errorf("some query error")

	location, _ = time.LoadLocation("Local")
	Created_at  = time.Date(2023, 11, 1, 0, 0, 0, 0, location)
	Updated_at  = time.Date(2023, 11, 2, 0, 0, 0, 0, location)

	СtxWithLogger = context.WithValue(context.Background(), LOGGER_KEY, InitCtxLogger())
)

const (
	SelectQuery      = "SELECT(.|\n)+FROM(.|\n)+"
	SelectExistQuery = "SELECT EXISTS(.|\n)+"
	InsertQuery      = "INSERT(.|\n)+INTO(.|\n)+"
	UpdateQuery      = "UPDATE(.|\n)+SET(.|\n)+WHERE(.|\n)+"
	DeleteQuery      = "DELETE(.|\n)+FROM(.|\n)+"
)

func InitCtxLogger() *logrus.Entry {
	logger := &logrus.Entry{
		Logger: &logrus.Logger{
			Out: io.Discard,
		},
	}
	return logger
}
