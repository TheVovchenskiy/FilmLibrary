package serverErrors

import (
	"filmLibrary/internal/repository"
	"filmLibrary/pkg/searchOptions"
	"filmLibrary/pkg/sortOptions"
	"net/http"
)

var HTTPErrors = map[error]int{
	ErrMethodNotAllowed:                 http.StatusMethodNotAllowed,
	sortOptions.ErrInvalidQueryParam:    http.StatusBadRequest,
	ErrInvalidRequest:                   http.StatusBadRequest,
	repository.ErrNotInserted:           http.StatusInternalServerError,
	repository.ErrNoRowsDeleted:         http.StatusInternalServerError,
	repository.ErrNoRowsUpdated:         http.StatusInternalServerError,
	repository.ErrInvalidFieldName:      http.StatusBadRequest,
	repository.ErrEmptyIds:              http.StatusBadRequest,
	searchOptions.ErrInvalidSearchQuery: http.StatusBadRequest,
}

func MapHTTPError(err error) (msg string, status int) {
	if err == nil {
		err = ErrInternal
	}

	status, ok := HTTPErrors[err]
	if !ok {
		status = http.StatusInternalServerError
		err = ErrInternal
	}

	msg = err.Error()

	return
}
