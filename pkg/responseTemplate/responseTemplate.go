package responseTemplate

import (
	"encoding/json"
	"filmLibrary/pkg/serverErrors"
	"net/http"
)

type ErrToSend struct {
	Message string `json:"message"`
}

func MarshalAndSend(w http.ResponseWriter, data any) {
	js, err := json.Marshal(data)
	if err != nil {
		ServeJsonError(w, serverErrors.ErrInternal)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func MarshalResponseError(errMsg string) []byte {
	data, _ := json.Marshal(ErrToSend{Message: errMsg})
	return data
}

func ServeJsonError(w http.ResponseWriter, err error) {
	msg, status := serverErrors.MapHTTPError(err)

	w.Header().Set("Content-Type", "application/json;")
	w.WriteHeader(status)
	w.Write(MarshalResponseError(msg))
}
