package serverErrors

import "errors"

var (
	ErrInternal         = errors.New("internal server error")
	ErrMethodNotAllowed = errors.New("method not allowed")
	ErrInvalidRequest   = errors.New("invalid request")
)
