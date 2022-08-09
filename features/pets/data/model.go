package data

import (
	"petadopter/domain"

	"gorm.io/gorm"
)

type Pets struct {
	gorm.Model
	Petname     string `json:"petname" form:"petname" validate:"required"`
	Gender      string `json:"gender" form:"gender" validate:"required"`
	Species     string `json:"species" form:"species" validate:"required"`
	Age         int    `json:"age" form:"age" validate:"required"`
	Color       string `json:"color" form:"color"`
	Description string `json:"description" form:"description"`
	Petphoto    string `json:"petphoto"`
}

type User struct {
	gorm.Model
	Username string `json:"username" form:"username" validate:"required"`
	Email    string `gorm:"unique" json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	FullName string `json:"fullname" form:"fullname" validate:"required"`
	Role     string `json:"role" form:"role" gorm:"default:users"`
	Photo    string `json:"image_url"`
}

func (p *Pets) ToDomain() domain.Pets {
	return domain.Pets{
		ID:          int(p.ID),
		Petname:     p.Petname,
		Species:     p.Species,
		Gender:      p.Gender,
		Age:         p.Age,
		Color:       p.Color,
		Petphoto:    p.Petphoto,
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
	res.Petname = data.Petname
	res.Species = data.Species
	res.Gender = data.Gender
	res.Age = data.Age
	res.Color = data.Color
	res.Petphoto = data.Petphoto
	res.Description = data.Description
	return res
}
