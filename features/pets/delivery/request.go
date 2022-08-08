package delivery

import (
	"petadopter/domain"
	"time"
)

type PetsInsertRequest struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Gender      string `json:"gender" form:"gender" validate:"required"`
	Age         int    `json:"age" form:"age" validate:"required"`
	Color       string `json:"color" form:"color"`
	Description string `json:"description" form:"description"`
	Images      string `json:"images"`
}

func (pi *PetsInsertRequest) ToDomain() domain.Pets {
	return domain.Pets{
		Name:        pi.Name,
		Gender:      pi.Gender,
		Age:         pi.Age,
		Color:       pi.Color,
		Images:      pi.Images,
		Description: pi.Description,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
}
