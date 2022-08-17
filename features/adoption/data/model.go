package data

import (
	"petadopter/domain"
	meetingData "petadopter/features/meeting/data"

	"gorm.io/gorm"
)

type Adoption struct {
	gorm.Model
	UserID  int
	PetsID  int `json:"petid" form:"petid"`
	Status  string
	Meeting []meetingData.Meeting `gorm:"foreignKey:AdoptionID"`
}

type AdoptionPet struct {
	ID           int
	PetsID       int
	Petname      string
	Petphoto     string
	Fullname     string
	PhotoProfile string
	City         string
	Status       string
}

type ApplierPet struct {
	ID           int
	Petname      string
	Ownername    string
	UserID       int
	Fullname     string
	PhotoProfile string
	Status       string
}

func (a *Adoption) ToDomain() domain.Adoption {
	return domain.Adoption{
		ID:     int(a.ID),
		PetsID: a.PetsID,
		UserID: a.UserID,
		Status: a.Status,
	}
}

func (a *AdoptionPet) ToDomainAdoptionPet() domain.AdoptionPet {
	return domain.AdoptionPet{
		ID:           int(a.ID),
		PetsID:       a.PetsID,
		Petname:      a.Petname,
		Petphoto:     a.Petphoto,
		Fullname:     a.Fullname,
		PhotoProfile: a.PhotoProfile,
		City:         a.City,
		Status:       a.Status,
	}
}

func (a *ApplierPet) ToDomainApplierPet() domain.ApplierPet {
	return domain.ApplierPet{
		UserID:       a.UserID,
		Fullname:     a.Fullname,
		PhotoProfile: a.PhotoProfile,
	}
}

func ParseToArr(arr []Adoption) []domain.Adoption {
	var res []domain.Adoption

	for _, val := range arr {
		res = append(res, val.ToDomain())
	}
	return res
}

func ParseToArrAdoptionPet(arr []AdoptionPet) []domain.AdoptionPet {
	var res []domain.AdoptionPet

	for _, val := range arr {
		res = append(res, val.ToDomainAdoptionPet())
	}
	return res
}

func ParseToArrApplierPet(arr []ApplierPet) []domain.ApplierPet {
	var res []domain.ApplierPet

	for _, val := range arr {
		res = append(res, val.ToDomainApplierPet())
	}
	return res
}

func ToLocal(data domain.Adoption) Adoption {
	var res Adoption
	res.UserID = data.UserID
	res.PetsID = data.PetsID
	res.Status = data.Status
	return res
}
