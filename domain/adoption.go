package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Adoption struct {
	ID        int
	IDPets    int
	IDUser    int
	Status    string
	Petphoto  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AdoptionUseCase interface {
	AddAdoption(IDUser int, newAdops Adoption) (Adoption, error)
	GetAllAP() ([]Adoption, error)
	UpAdoption(IDAdoption int, updateData Adoption) (Adoption, error)
	DelAdoption(IDAdoption int) (bool, error)
	GetSpecificAdoption(AdoptionID int) ([]Adoption, error)
}

type AdoptionHandler interface {
	InsertAdoption() echo.HandlerFunc
	GetAllAdoption() echo.HandlerFunc
	UpdateAdoption() echo.HandlerFunc
	DeleteAdoption() echo.HandlerFunc
	GetAdoptionID() echo.HandlerFunc
}

type AdoptionData interface {
	Insert(insertAdoption Adoption) Adoption
	GetAll() []Adoption
	Update(IDAdoption int, updatedAdoption Adoption) Adoption
	Delete(IDAdoption int) bool
	GetAdoptionID(AdoptionID int) []Adoption
}
