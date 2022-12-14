package delivery

import "petadopter/domain"

type PetsResponse struct {
	ID          int    `json:"id"`
	Petname     string `json:"petname" form:"petname" validate:"required"`
	Gender      int    `json:"gender" form:"gender" validate:"required"`
	Age         int    `json:"age" form:"age" validate:"required"`
	Color       string `json:"color" form:"color"`
	Description string `json:"description" form:"description"`
	Petphoto    string `json:"petphoto"`
}

func FromDomain(data domain.Pets) PetsResponse {
	var res PetsResponse
	res.ID = int(data.ID)
	res.Petname = data.Petname
	res.Gender = data.Gender
	res.Age = data.Age
	res.Color = data.Color
	res.Description = data.Description
	res.Petphoto = data.Petphoto
	return res
}
