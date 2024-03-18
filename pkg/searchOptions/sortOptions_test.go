package searchOptions_test

import (
	"errors"
	"filmLibrary/pkg/sortOptions"
	"net/url"
	"reflect"
	"testing"
)

func TestGetSortOptions(t *testing.T) {
	tests := []struct {
		name        string
		query       url.Values
		expected    map[sortOptions.SortOptionName]sortOptions.SortOptionValue
		expectedErr error
	}{
		{
			"valid 2 params",
			url.Values{"order": []string{"descending"}, "field": []string{"name"}},
			map[sortOptions.SortOptionName]sortOptions.SortOptionValue{
				sortOptions.SortOrder: sortOptions.Descending,
				sortOptions.SortFiled: sortOptions.Name,
			},
			nil,
		},
		{
			"valid order param",
			url.Values{"order": []string{"ascending"}},
			map[sortOptions.SortOptionName]sortOptions.SortOptionValue{
				sortOptions.SortOrder: sortOptions.Ascending,
				sortOptions.SortFiled: sortOptions.Rating,
			},
			nil,
		},
		{
			"valid field param",
			url.Values{"field": []string{"name"}},
			map[sortOptions.SortOptionName]sortOptions.SortOptionValue{
				sortOptions.SortOrder: sortOptions.Descending,
				sortOptions.SortFiled: sortOptions.Name,
			},
			nil,
		},
		{
			"empty params",
			url.Values{},
			map[sortOptions.SortOptionName]sortOptions.SortOptionValue{
				sortOptions.SortOrder: sortOptions.Descending,
				sortOptions.SortFiled: sortOptions.Rating,
			},
			nil,
		},
		{
			"3 params",
			url.Values{"order": []string{"descending"}, "field": []string{"name"}, "sdjf": []string{"fsdfsd"}},
			map[sortOptions.SortOptionName]sortOptions.SortOptionValue{
				sortOptions.SortOrder: sortOptions.Descending,
				sortOptions.SortFiled: sortOptions.Name,
			},
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, actualErr := sortOptions.GetSortOptions(tt.query)

			if !reflect.DeepEqual(actual, tt.expected) || !errors.Is(actualErr, tt.expectedErr) {
				t.Errorf("GetSearchQuery()\n\tgot: %v, %v,\n\twant %v, %v", actual, actualErr, tt.expected, tt.expectedErr)
			}
		})
	}
}
