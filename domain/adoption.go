package domain

import (
	"github.com/labstack/echo/v4"
)

type Adoption struct {
	ID     int
	PetsID int
	UserID int
	Status string
}

type AdoptionPet struct {
	ID           int
	PetsID       int
	Petname      string
	Petphoto     string
	Userid       int
	Fullname     string
	PhotoProfile string
	City         string
	Status       string
}

type ApplierPet struct {
	Fullname     string
	UserID       int
	PhotoProfile string
}

type AdoptionUseCase interface {
	AddAdoption(IDUser int, newAdops Adoption) (Adoption, error)
	GetAllAP(userID int) ([]map[string]interface{}, error)
	UpAdoption(IDAdoption int, updateData Adoption, userID int) (Adoption, error)
	DelAdoption(IDAdoption int) (bool, error)
	GetSpecificAdoption(AdoptionID int) (map[string]interface{}, error)
	GetmyAdoption(userID int) ([]map[string]interface{}, error)
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
	GetAll(userID int) ([]AdoptionPet, []ApplierPet)
	Update(IDAdoption int, updatedAdoption Adoption) Adoption
	Delete(IDAdoption int) bool
	GetAdoptionID(AdoptionID int) []AdoptionPet
	GetAdoptionbyuser(userID int) []AdoptionPet
}
