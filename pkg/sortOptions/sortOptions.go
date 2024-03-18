package sortOptions

import (
	"filmLibrary/pkg/utils"
	"net/url"
)

type SortOptionName string

const (
	SortOrder SortOptionName = "order"
	SortFiled SortOptionName = "field"
)

type SortOptionValue string

const (
	Descending    SortOptionValue = "descending"
	DescendingSQL SortOptionValue = "DESC"
	Ascending     SortOptionValue = "ascending"
	AscendingSQL  SortOptionValue = "ASC"
	Name          SortOptionValue = "name"
	Rating        SortOptionValue = "rating"
	ReleaseDate   SortOptionValue = "release_date"
)

func GetSortOptions(query url.Values) (map[SortOptionName]SortOptionValue, error) {
	res := map[SortOptionName]SortOptionValue{}

	for key, values := range query {
		switch key {
		case string(SortOrder):
			if _, ok := res[SortOptionName(key)]; !ok && utils.In(values[0], []string{string(Descending), string(Ascending)}) {
				res[SortOptionName(key)] = SortOptionValue(values[0])
			} else {
				return nil, ErrInvalidQueryParam
			}
		case string(SortFiled):
			if _, ok := res[SortOptionName(key)]; !ok && utils.In(values[0], []string{string(Name), string(Rating), string(ReleaseDate)}) {
				res[SortOptionName(key)] = SortOptionValue(values[0])
			} else {
				return nil, ErrInvalidQueryParam
			}
		}
	}
	switch len(res) {
	case 0:
		res[SortOrder] = Descending
		res[SortFiled] = Rating
	case 1:
		if _, ok := res[SortOrder]; ok {
			res[SortFiled] = Rating
		} else {
			res[SortOrder] = Descending
		}
	}
	return res, nil
}

func MapSortOrderSQL(order SortOptionValue) SortOptionValue {
	return map[SortOptionValue]SortOptionValue{
		Descending: DescendingSQL,
		Ascending:  AscendingSQL,
	}[order]
}
