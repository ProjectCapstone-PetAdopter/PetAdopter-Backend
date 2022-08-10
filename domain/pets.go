package domain

import (
	"github.com/labstack/echo/v4"
)

type Pets struct {
	ID          int
	Petname     string
	Gender      int
	Age         int
	Color       string
	Petphoto    string
	Species     string
	Description string
	Userid      int
}

type PetUser struct {
	Fullname string
	City     string
}

type PetsUseCase interface {
	AddPets(newPets Pets) (Pets, error)
	GetAllP() ([]Pets, error)
	UpPets(IDPets int, updateData Pets) (Pets, error)
	DelPets(IDPets int) (bool, error)
	GetSpecificPets(PetsID int) ([]Pets, PetUser, error)
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
	GetPetUser() PetUser
}
