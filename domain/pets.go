package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Pets struct {
	ID          int
	Name        string
	Gender      string
	Age         int
	Color       string
	Images      string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type PetsUseCase interface {
	AddPets(IDUser int, newPets Pets) (Pets, error)
	GetAllP() ([]Pets, error)
	UpPets(IDPets int, updateData Pets) (Pets, error)
	DelPets(IDPets int) (bool, error)
	GetSpecificPets(PetsID int) ([]Pets, error)
	GetmyPets(userID int) ([]Pets, error)
}

type PetsHandler interface {
	InsertPets() echo.HandlerFunc
	GetAllPets() echo.HandlerFunc
	UpdatePets() echo.HandlerFunc
	DeletePets() echo.HandlerFunc
	GetPetsID() echo.HandlerFunc
	GetmyPets() echo.HandlerFunc
}

type PetsData interface {
	Insert(insertPets Pets) Pets
	GetAll() []Pets
	Update(IDPets int, updatedPets Pets) Pets
	Delete(IDPets int) bool
	GetPetsID(PetsID int) []Pets
	GetPetsbyuser(userID int) []Pets
}
