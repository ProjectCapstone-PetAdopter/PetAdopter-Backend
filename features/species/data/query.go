package data

import (
	"errors"
	"fmt"
	"log"
	"petadopter/domain"

	"gorm.io/gorm"
)

type speciesData struct {
	db *gorm.DB
}

func New(DB *gorm.DB) domain.SpeciesData {
	return &speciesData{
		db: DB,
	}
}

func (sd *speciesData) InsertSpecies(newSpecies domain.Species) (idSpecies int, err error) {
	cnv := FromModel(newSpecies)
	// fmt.Println(cnv)
	result := sd.db.Table("species").Create(&cnv)
	if result.Error != nil {
		log.Println("Cannot create object", errors.New("error from db"))
		return -1, errors.New("species already exist")
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("failed insert data")
	}
	return int(result.RowsAffected), nil
}

// func (sd *speciesData) GetUser(id uint) error {
// 	var user domain.User
// 	err := sd.db.Where("id = ? and role = `admin`", id).First(&user).Error
// 	if err != nil {
// 		return errors.New("you are not admin")
// 	}
// 	return nil
// }

func (sd *speciesData) GetAll() ([]domain.Species, error) {
	var speciesData []Species
	err := sd.db.Find(&speciesData).Error
	if err != nil {
		return []domain.Species{}, err
	}
	var convert []domain.Species
	for i := 0; i < len(speciesData); i++ {
		convert = append(convert, speciesData[i].ToModel())
	}
	return convert, nil
}

func (sd *speciesData) Update(id int, updatedData domain.Species) (idSpecies int, err error) {
	var cnv = FromModel(updatedData)
	cnv.ID = uint(id)
	result := sd.db.Model(&Species{}).Where("ID = ?", id).Updates(cnv)
	if result.Error != nil {
		log.Println("Cannot update data", errors.New("error db"))
		return -1, errors.New("species already exist")
	}
	if result.RowsAffected == 0 {
		return -1, errors.New("failed update data")
	}

	return int(result.RowsAffected), nil
}

func (sd *speciesData) Delete(id int) (idSpecies int, err error) {
	res := sd.db.Delete(&Species{}, id)
	if res.Error != nil {
		log.Println("cannot delete data", res.Error.Error())
		return 0, res.Error
	}
	if res.RowsAffected < 1 {
		log.Println("no data deleted", res.Error.Error())
		return 0, fmt.Errorf("failed to delete species")
	}
	return int(res.RowsAffected), nil
}
