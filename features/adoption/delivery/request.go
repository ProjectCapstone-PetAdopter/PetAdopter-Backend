package delivery

import (
	"petadopter/domain"
	"time"
)

type AdoptionInsertRequest struct {
	UserID int
	PetsID int    `json:"petid" form:"petid"`
	Status string `gorm:"default:waiting"`
}

func (ai *AdoptionInsertRequest) ToDomain() domain.Adoption {
	return domain.Adoption{
		PetsID:    ai.PetsID,
		UserID:    ai.UserID,
		Status:    ai.Status,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}
