package serverErrors

import "net/http"

var HTTPErrors = map[error]int{
	ErrMethodNotAllowed: http.StatusMethodNotAllowed,
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
