package model

import "time"

type APIMovie struct {
	Id          int       `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Descriprion string    `json:"description,omitempty"`
	ReleaseDate string    `json:"releaseDate,omitempty"`
	Rating      float32   `json:"rating,omitempty"`
	// Stars       []APIStar `json:"stars,omitempty"`
}

func (m *APIMovie) ToDB() *DBMovie {
	res := &DBMovie{
		Id:          m.Id,
		Name:        m.Name,
		Descriprion: m.Descriprion,
		Rating:      m.Rating,
	}

	releaseDate, err := time.Parse(time.RFC3339, m.ReleaseDate)
	if err != nil {
		releaseDateStr := releaseDate.Format(time.DateOnly)
		res.ReleaseDate = releaseDateStr
	}

	return res
}

type DBMovie struct {
	Id          int
	Name        string
	Descriprion string
	ReleaseDate string
	Rating      float32
}

func (m *DBMovie) ToAPI() *APIMovie {
	return &APIMovie{
		Id:          m.Id,
		Name:        m.Name,
		Descriprion: m.Descriprion,
		ReleaseDate: m.ReleaseDate,
		Rating:      m.Rating,
	}
}
