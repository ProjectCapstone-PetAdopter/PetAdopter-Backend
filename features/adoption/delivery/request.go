package delivery

import (
	"petadopter/domain"
	"time"
)

type AdoptionInsertRequest struct {
	UserID   uint   `json:"user_id" form:"user_id"`
	PetsID   uint   `json:"pets_id" form:"pets_id"`
	Status   string `json:"status" form:"status"   gorm:"default:waiting"`
	Petphoto string `json:"petphoto" form:"petphoto"`
}

func (ai *AdoptionInsertRequest) ToDomain() domain.Adoption {
	return domain.Adoption{
		PetsID:    int(ai.PetsID),
		IDUser:    int(ai.UserID),
		Petphoto:  ai.Petphoto,
		Status:    ai.Status,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}
