package delivery

import "petadopter/domain"

type InserFormat struct {
	Species string `json:"species" form:"species"`
}

func (i InserFormat) ToModel() domain.Species {
	return domain.Species{
		Species: i.Species,
	}
}
