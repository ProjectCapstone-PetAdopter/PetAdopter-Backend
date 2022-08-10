package domain

import "github.com/labstack/echo/v4"

type Species struct {
	ID      int
	Species string
	Pet     []Pets
}

type SpeciesHandler interface {
	AddSpecies() echo.HandlerFunc
	GetSpecies() echo.HandlerFunc
	UpdateDataSpecies() echo.HandlerFunc
	DeleteDataSpecies() echo.HandlerFunc
}

type SpeciesUsecase interface {
	AddSpecies(newSpecies Species) (row int, err error)
	// GetUser(id uint) error
	GetAllSpecies() ([]Species, error)
	UpdateSpecies(id int, UpdateSpecies Species) (row int, err error)
	DeleteSpecies(id int) (row int, err error)
}

type SpeciesData interface {
	InsertSpecies(newSpecies Species) (row int, err error)
	// GetUser(id uint) error
	GetAll() ([]Species, error)
	Update(id int, updatedData Species) (row int, err error)
	Delete(id int) (row int, err error)
}
