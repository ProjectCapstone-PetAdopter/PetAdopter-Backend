package delivery

import "petadopter/domain"

type AdoptionResponse struct {
	ID       int    `json:"id"`
	UserID   uint   `json:"user_id"`
	PetsID   uint   `json:"pets_id"`
	Status   string `json:"status" gorm:"default:waiting"`
	Petphoto string `json:"petphoto"`
}

func FromDomain(data domain.Adoption) AdoptionResponse {
	var res AdoptionResponse
	res.ID = int(data.ID)
	res.UserID = uint(data.IDUser)
	res.PetsID = uint(data.PetsID)
	res.Status = data.Status
	res.Petphoto = data.Petphoto
	return res
}
