package data

import (
	"fmt"
	"log"
	"petadopter/domain"

	"gorm.io/gorm"
)

type petsData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.PetsData {
	return &petsData{
		db: db,
	}
}

func (pd *petsData) Insert(newPets domain.Pets) domain.Pets {
	cnv := ToLocal(newPets)
	err := pd.db.Create(&cnv)
	fmt.Println("error", err.Error)
	if err.Error != nil {
		return domain.Pets{}
	}
	return cnv.ToDomain()
}

func (pd *petsData) Update(petsID int, updatedProduct domain.Pets) domain.Pets {
	cnv := ToLocal(updatedProduct)
	err := pd.db.Model(cnv).Where("ID = ?", petsID).Updates(updatedProduct)
	if err.Error != nil {
		log.Println("Cannot update data", err.Error.Error())
		return domain.Pets{}
	}
	cnv.ID = uint(petsID)
	return cnv.ToDomain()
}

func (pd *petsData) Delete(petsID int) bool {
	err := pd.db.Where("ID = ?", petsID).Delete(&Pets{})
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

func (pd *petsData) GetAll() []domain.Pets {
	var data []Pets
	err := pd.db.Find(&data)

	if err.Error != nil {
		log.Println("error on select data", err.Error.Error())
		return nil
	}

	return ParseToArr(data)
}

func (pd *petsData) GetPetsID(petsID int) []domain.Pets {
	var data []Pets
	err := pd.db.Where("ID = ?", petsID).First(&data)

	if err.Error != nil {
		log.Println("problem data", err.Error.Error())
		return nil
	}
	return ParseToArr(data)
}

func (pd *petsData) GetPetsbyuser(userID int) []domain.Pets {
	var data []Pets
	err := pd.db.Where("user_id = ?", userID).Find(&data)

	if err.Error != nil {
		log.Println("problem data", err.Error.Error())
		return nil
	}
	return ParseToArr(data)
}
