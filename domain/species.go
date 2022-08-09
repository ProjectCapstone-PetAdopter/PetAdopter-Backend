package domain

type Species struct {
	ID      int
	Species string
}

type SpeciesHandler interface {
}

type SpeciesUsecase interface {
	AddSpeciesUseCase(newSpecies Species) (row int, err error)
	GetUser(id uint) error
	GetAllSpecies() ([]Species, error)
	UpdateSpecies(id int, UpdateSpecies Species) (row int, err error)
}

type SpeciesData interface {
	InsertSpeciesQuery(newSpecies Species) (row int, err error)
	GetUser(id uint) error
	GetAll() ([]Species, error)
	Update(id int, updatedData Species) (row int, err error)
}
