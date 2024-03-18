package searchOptions

import "net/url"

type SearchOptionName string

const (
	Query SearchOptionName = "q"
)

func GetSearchQuery(query url.Values) (string, error) {
	queryOptionName := query.Get(string(Query))
	if queryOptionName == "" {
		return "", ErrInvalidSearchQuery
	}
	return queryOptionName, nil
}
