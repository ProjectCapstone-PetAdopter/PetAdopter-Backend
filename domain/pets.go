package domain

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Pets struct {
	ID          int
	Petname     string
	Gender      int
	Age         int
	Color       string
	Petphoto    string
	Speciesid   int
	Description string
	Userid      int
	Adoption    []Adoption
}

type PetUser struct {
	Species  string
	Fullname string
	City     string
}

type PetsUseCase interface {
	AddPets(newPets Pets, userId int, form *multipart.FileHeader) (Pets, error)
	GetAllP() ([]map[string]interface{}, error)
	UpPets(IDPets int, updateData Pets, userID int, form *multipart.FileHeader) (Pets, error)
	DelPets(IDPets int) (bool, error)
	GetSpecificPets(PetsID int) (map[string]interface{}, error)
	GetmyPets(userID int) ([]map[string]interface{}, error)
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
	GetPetUser(userID int) PetUser
	GetAllPetUser() []PetUser
}
