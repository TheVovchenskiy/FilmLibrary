package responseTemplate_test

import (
	"errors"
	"filmLibrary/pkg/responseTemplate"
	"filmLibrary/pkg/serverErrors"
	"net/http/httptest"
	"testing"
)

func TestServeJSONError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want string
	}{
		{"valid error", serverErrors.ErrInternal, `{"message":"internal server error"}`},
		{"empty error", errors.New(""), `{"message":"internal server error"}`},
		{"nil error", nil, `{"message":"internal server error"}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			responseTemplate.ServeJsonError(rr, tt.err)

			if rr.Body.String() != tt.want {
				t.Errorf("ServeJSONError() = %v, want %v", rr.Body.String(), tt.want)
			}
		})
	}
}
