package data

import (
	"fmt"
	"log"
	"petadopter/domain"

	"gorm.io/gorm"
)

type adoptionData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.AdoptionData {
	return &adoptionData{
		db: db,
	}
}

func (ad *adoptionData) Insert(newAdopt domain.Adoption) domain.Adoption {
	cnv := ToLocal(newAdopt)
	err := ad.db.Create(&cnv)
	fmt.Println("error", err.Error)
	if err.Error != nil {
		return domain.Adoption{}
	}
	return cnv.ToDomain()
}
func (ad *adoptionData) Update(adoptID int, updatedAdopt domain.Adoption) domain.Adoption {
	cnv := ToLocal(updatedAdopt)
	err := ad.db.Model(cnv).Where("ID = ?", adoptID).Updates(updatedAdopt)
	if err.Error != nil {
		log.Println("Cannot update data", err.Error.Error())
		return domain.Adoption{}
	}
	cnv.ID = uint(adoptID)
	return cnv.ToDomain()
}

func (ad *adoptionData) Delete(adoptID int) bool {
	err := ad.db.Where("ID = ?", adoptID).Delete(&Adoption{})
	if err.Error != nil {
		log.Println("Cannot delete data", err.Error.Error())
		return false
	}
	if err.RowsAffected < 1 {
		log.Println("No data deleted", err.Error.Error())
		return false
	}
	return true
}

func (ad *adoptionData) GetAll() []domain.Adoption {
	var data []Adoption
	err := ad.db.Find(&data)

	if err.Error != nil {
		log.Println("error on select data", err.Error.Error())
		return nil
	}

	return ParseToArr(data)
}

func (ad *adoptionData) GetAdoptionID(adoptID int) []domain.Adoption {
	var data []Adoption
	err := ad.db.Where("ID = ?", adoptID).First(&data)

	if err.Error != nil {
		log.Println("problem data", err.Error.Error())
		return nil
	}
	return ParseToArr(data)
}

func (ad *adoptionData) GetAdoptionbyuser(userID int) []domain.Adoption {
	var data []Adoption
	err := ad.db.Where("user_id = ?", userID).Find(&data)

	if err.Error != nil {
		log.Println("problem data", err.Error.Error())
		return nil
	}
	return ParseToArr(data)
}
