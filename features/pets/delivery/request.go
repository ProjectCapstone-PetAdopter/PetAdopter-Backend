package delivery

import (
	"petadopter/domain"
)

type PetsInsertRequest struct {
	Petname     string `json:"petname" form:"petname" validate:"required"`
	Gender      int    `json:"gender" form:"gender" validate:"required"`
	Age         int    `json:"age" form:"age" validate:"required"`
	Color       string `json:"color" form:"color"`
	Description string `json:"description" form:"description"`
	Petphoto    string `json:"petphoto"`
	Speciesid   int
	Userid      int
}

func (pi *PetsInsertRequest) ToDomain() domain.Pets {
	return domain.Pets{
		Petname:     pi.Petname,
		Gender:      pi.Gender,
		Age:         pi.Age,
		Color:       pi.Color,
		Petphoto:    pi.Petphoto,
		Description: pi.Description,
		Speciesid:   pi.Speciesid,
		Userid:      pi.Userid,
	}
}
