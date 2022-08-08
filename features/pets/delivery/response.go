package delivery

import "petadopter/domain"

type PetsResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name" form:"name"`
	Gender      string `json:"gender" form:"gender"`
	Age         int    `json:"age" form:"age"`
	Color       string `json:"color" form:"color"`
	Description string `json:"description" form:"description"`
	Images      string `json:"images"`
}

func FromDomain(data domain.Pets) PetsResponse {
	var res PetsResponse
	res.ID = int(data.ID)
	res.Name = data.Name
	res.Gender = data.Gender
	res.Age = data.Age
	res.Color = data.Color
	res.Description = data.Description
	res.Images = data.Images
	return res
}
