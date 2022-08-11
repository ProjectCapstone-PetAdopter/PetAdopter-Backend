package data

import (
	"petadopter/domain"
	meetingData "petadopter/features/meeting/data"

	"gorm.io/gorm"
)

type Adoption struct {
	gorm.Model
	UserID   uint                  `json:"user_id" form:"user_id"`
	PetsID   uint                  `json:"pets_id" form:"pets_id"`
	Status   string                `json:"status" form:"status"   gorm:"default:waiting"`
	Petphoto string                `json:"petphoto" form:"petphoto"`
	Meeting  []meetingData.Meeting `gorm:"foreignKey:AdoptionID"`
}

// type Pets struct {
// 	gorm.Model
// 	Name        string `json:"name" form:"name" validate:"required"`
// 	Gender      string `json:"gender" form:"gender" validate:"required"`
// 	Age         int    `json:"age" form:"age" validate:"required"`
// 	Color       string `json:"color" form:"color"`
// 	Description string `json:"description" form:"description"`
// 	Images      string `json:"images"`
// }

// type User struct {
// 	gorm.Model
// 	Username string `json:"username" form:"username" validate:"required"`
// 	Email    string `gorm:"unique" json:"email" form:"email" validate:"required"`
// 	Password string `json:"password" form:"password" validate:"required"`
// 	FullName string `json:"fullname" form:"fullname" validate:"required"`
// 	Role     string `json:"role" form:"role" gorm:"default:users"`
// 	Photo    string `json:"image_url"`
// }

func (a *Adoption) ToDomain() domain.Adoption {
	return domain.Adoption{
		ID:        int(a.ID),
		PetsID:    int(a.PetsID),
		IDUser:    int(a.UserID),
		Petphoto:  a.Petphoto,
		Status:    a.Status,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

func ParseToArr(arr []Adoption) []domain.Adoption {
	var res []domain.Adoption

	for _, val := range arr {
		res = append(res, val.ToDomain())
	}
	return res
}

func ToLocal(data domain.Adoption) Adoption {
	var res Adoption
	res.UserID = uint(data.IDUser)
	res.PetsID = uint(data.PetsID)
	res.Petphoto = data.Petphoto
	res.Status = data.Status
	return res
}
