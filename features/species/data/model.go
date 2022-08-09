package data

import (
	"petadopter/domain"
	petsData "petadopter/features/pets/data"

	"gorm.io/gorm"
)

type Species struct {
	gorm.Model
	Species string
	Pets    []petsData.Pets
}

func (s *Species) ToModel() domain.Species {
	return domain.Species{
		ID:      int(s.ID),
		Species: s.Species,
	}
}

func ParseToArr(arr []Species) []domain.Species {
	var res []domain.Species

	for _, val := range arr {
		res = append(res, val.ToModel())
	}
	return res
}

func FromModel(data domain.Species) Species {
	var res Species
	res.Species = data.Species
	return res
}
