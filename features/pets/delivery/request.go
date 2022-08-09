package delivery

import (
	"petadopter/domain"
	"time"
)

type PetsInsertRequest struct {
	Petname     string `json:"petname" form:"petname" validate:"required"`
	Gender      string `json:"gender" form:"gender" validate:"required"`
	Species     string `json:"species" form:"species" validate:"required"`
	Age         int    `json:"age" form:"age" validate:"required"`
	Color       string `json:"color" form:"color"`
	Description string `json:"description" form:"description"`
	Petphoto    string `json:"petphoto"`
}

func (pi *PetsInsertRequest) ToDomain() domain.Pets {
	return domain.Pets{
		Petname:     pi.Petname,
		Species:     pi.Species,
		Gender:      pi.Gender,
		Age:         pi.Age,
		Color:       pi.Color,
		Petphoto:    pi.Petphoto,
		Description: pi.Description,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
}
