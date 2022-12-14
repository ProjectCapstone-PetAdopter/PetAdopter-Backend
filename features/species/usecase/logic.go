package usecase

import (
	"errors"
	"fmt"
	"petadopter/domain"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type speciesUseCase struct {
	speciesData domain.SpeciesData
	validate    *validator.Validate
}

func New(sd domain.SpeciesData, v *validator.Validate) domain.SpeciesUsecase {
	return &speciesUseCase{
		speciesData: sd,
		validate:    v,
	}
}

func (su *speciesUseCase) AddSpecies(newSpecies domain.Species) (idSpecies int, err error) {
	if newSpecies.Species == "" {
		return -1, errors.New("invalid species")
	}

	fmt.Println(newSpecies.Species)
	inserted, err := su.speciesData.InsertSpecies(newSpecies)
	return inserted, err
}

func (su *speciesUseCase) GetAllSpecies() ([]domain.Species, error) {
	data, err := su.speciesData.GetAll()
	return data, err
}

func (su *speciesUseCase) UpdateSpecies(id int, UpdateSpecies domain.Species) (idSpecies int, err error) {

	if UpdateSpecies.Species == "" {
		return -1, errors.New("invalid species")
	}

	data, err := su.speciesData.Update(id, UpdateSpecies)
	return data, err
}

func (su *speciesUseCase) DeleteSpecies(id int) (idSpecies int, err error) {
	idSpecies, err = su.speciesData.Delete(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return idSpecies, errors.New("data not found")
		} else {
			return idSpecies, errors.New("failed to delete species")
		}
	}
	return idSpecies, nil
}
