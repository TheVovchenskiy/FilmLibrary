package model

type Gender string

const (
	Male   Gender = "M"
	Female Gender = "F"
)

type APIStar struct {
	Id       int      `json:"id,omitempty"`
	Name     string   `json:"name,omitempty"`
	Gender   Gender   `json:"gender,omitempty"`
	Birthday string   `json:"birthday,omitempty"`
	Movies   []string `json:"movies,omitempty"`
	// Movies   []APIMovie `json:"movies,omitempty"`
}

func (m *APIStar) ToDB() *DBStar {
	return &DBStar{
		Id:       m.Id,
		Name:     m.Name,
		Gender:   m.Gender,
		Birthday: m.Birthday,
	}
}

type DBStar struct {
	Id       int
	Name     string
	Gender   Gender
	Birthday string
}

func (m *DBStar) ToAPI() *APIStar {
	return &APIStar{
		Id:       m.Id,
		Name:     m.Name,
		Gender:   m.Gender,
		Birthday: m.Birthday,
	}
}
