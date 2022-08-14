package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Adoption struct {
	ID        int
	PetsID    int
	UserID    int
	Status    string
	Meeting   []Meeting
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AdoptionPet struct {
	ID           int
	Petname      string
	Petphoto     string
	Fullname     string
	PhotoProfile string
	Address      string
	Status       string
}

type AdoptionUseCase interface {
	AddAdoption(IDUser int, newAdops Adoption) (Adoption, error)
	GetAllAP(userID int) ([]AdoptionPet, error)
	UpAdoption(IDAdoption int, updateData Adoption) (Adoption, error)
	DelAdoption(IDAdoption int) (bool, error)
	GetSpecificAdoption(AdoptionID int) ([]AdoptionPet, error)
	GetmyAdoption(userID int) ([]AdoptionPet, error)
}

type AdoptionHandler interface {
	InsertAdoption() echo.HandlerFunc
	GetAllAdoption() echo.HandlerFunc
	UpdateAdoption() echo.HandlerFunc
	DeleteAdoption() echo.HandlerFunc
	GetAdoptionID() echo.HandlerFunc
	GetMYAdopt() echo.HandlerFunc
}

type AdoptionData interface {
	Insert(insertAdoption Adoption) Adoption
	GetAll(userID int) []AdoptionPet
	Update(IDAdoption int, updatedAdoption Adoption) Adoption
	Delete(IDAdoption int) bool
	GetAdoptionID(AdoptionID int) []AdoptionPet
	GetAdoptionbyuser(userID int) []AdoptionPet
}
