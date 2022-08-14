package data

import (
	"petadopter/domain"
	meetingData "petadopter/features/meeting/data"

	"gorm.io/gorm"
)

type Adoption struct {
	gorm.Model
	UserID  int
	PetsID  int                   `json:"petid" form:"petid"`
	Status  string                `gorm:"default:waiting"`
	Meeting []meetingData.Meeting `gorm:"foreignKey:AdoptionID"`
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

func (a *Adoption) ToDomain() domain.Adoption {
	return domain.Adoption{
		ID:        int(a.ID),
		PetsID:    int(a.PetsID),
		UserID:    a.UserID,
		Status:    a.Status,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

func (a *AdoptionPet) ToDomainAdoptionPet() domain.AdoptionPet {
	return domain.AdoptionPet{
		ID:           int(a.ID),
		Petname:      a.Petname,
		Petphoto:     a.Petphoto,
		Fullname:     a.Fullname,
		PhotoProfile: a.PhotoProfile,
		Address:      a.Address,
		Status:       a.Status,
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

func ToLocal(data domain.Adoption) Adoption {
	var res Adoption
	res.UserID = data.UserID
	res.PetsID = data.PetsID
	res.Status = data.Status
	return res
}
