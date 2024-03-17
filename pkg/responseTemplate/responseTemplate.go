package responseTemplate

import (
	"encoding/json"
	"filmLibrary/pkg/serverErrors"
	"net/http"
)

func MarshalResponseError(errMsg string) []byte {
	data, _ := json.Marshal(map[string]string{"message": errMsg})
	return data
}

func ServeJsonError(w http.ResponseWriter, err error) {
	msg, status := serverErrors.MapHTTPError(err)

	w.Header().Set("Content-Type", "application/json;")
	w.WriteHeader(status)
	w.Write(MarshalResponseError(msg))
}
