package repository

import "errors"

var (
	ErrNotInserted      = errors.New("data not inserted")
	ErrNoRowsUpdated    = errors.New("rows are not updated")
	ErrNoRowsDeleted    = errors.New("rows are not deleted")
	ErrInvalidFieldName = errors.New("invalid field name")
	ErrEmptyIds         = errors.New("empty slice of ids")
)
