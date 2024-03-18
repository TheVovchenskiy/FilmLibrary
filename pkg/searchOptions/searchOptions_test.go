package searchOptions_test

import (
	"errors"
	"filmLibrary/pkg/searchOptions"
	"net/url"
	"testing"
)

func TestGetSearchQuery(t *testing.T) {
	tests := []struct {
		name          string
		query         url.Values
		expectedQuery string
		expectedErr   error
	}{
		{"valid values", url.Values{"q": []string{"query"}}, "query", nil},
		{"empty", url.Values{"": []string{}}, "", searchOptions.ErrInvalidSearchQuery},
		{"invalid query param", url.Values{"w": []string{"query"}}, "", searchOptions.ErrInvalidSearchQuery},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualQuery, actualErr := searchOptions.GetSearchQuery(tt.query)

			if actualQuery != tt.expectedQuery || !errors.Is(actualErr, tt.expectedErr) {
				t.Errorf("GetSearchQuery() = %v, %v, want %v, %v", actualQuery, actualErr, tt.expectedQuery, tt.expectedErr)
			}
		})
	}
}
