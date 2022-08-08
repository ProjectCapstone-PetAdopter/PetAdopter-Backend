package data

import (
	"petadopter/domain"

	"gorm.io/gorm"
)

type Pets struct {
	gorm.Model
	Name        string `json:"name" form:"name" validate:"required"`
	Gender      string `json:"gender" form:"gender" validate:"required"`
	Age         int    `json:"age" form:"age" validate:"required"`
	Color       string `json:"color" form:"color"`
	Description string `json:"description" form:"description"`
	Images      string `json:"images"`
}

func (p *Pets) ToDomain() domain.Pets {
	return domain.Pets{
		ID:          int(p.ID),
		Name:        p.Name,
		Gender:      p.Gender,
		Age:         p.Age,
		Color:       p.Color,
		Images:      p.Images,
		Description: p.Description,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func ParseToArr(arr []Pets) []domain.Pets {
	var res []domain.Pets

	for _, val := range arr {
		res = append(res, val.ToDomain())
	}
	return res
}

func ToLocal(data domain.Pets) Pets {
	var res Pets
	res.Name = data.Name
	res.Gender = data.Gender
	res.Age = data.Age
	res.Color = data.Color
	res.Images = data.Images
	res.Description = data.Description
	return res
}
