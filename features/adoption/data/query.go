package data

import (
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
	var ownerid int

	getownerid := ad.db.Model(&Adoption{}).Select("pets.userid").Joins("join pets on adoptions.pets_id = pets.id").
		Where("adoptions.pets_id = ?", newAdopt.PetsID).Scan(&ownerid)

	if getownerid.Error != nil {
		log.Println("Cannot get pet id", getownerid.Error.Error())
		return domain.Adoption{}
	}
	//memverifikasi user yang mengakses
	if newAdopt.UserID == ownerid {
		log.Println("Cant adopt owned pet")
		return domain.Adoption{}
	}

	err := ad.db.Create(&cnv)
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

func (ad *adoptionData) GetAll(userid int) ([]domain.AdoptionPet, []domain.ApplierPet) {
	var data []AdoptionPet
	var dataSeeker []ApplierPet

	//mengambil data seeker untuk mendapatkan seekername dan seekerphoto
	getSeeker := ad.db.Model(&Adoption{}).Select("adoptions.user_id, users.fullname, users.photo_profile").
		Joins("join pets on adoptions.pets_id = pets.id").Joins("join users on adoptions.user_id = users.id").Where("pets.userid = ?", userid).Scan(&dataSeeker)

	if getSeeker.Error != nil {
		log.Println("problem data seeker", getSeeker.Error.Error())
		return nil, nil
	}
	//mengambil data owner dan pet
	err := ad.db.Model(&Adoption{}).Select("adoptions.id, pets.petname, pets.userid, users.fullname, adoptions.status").
		Joins("join pets on adoptions.pets_id = pets.id").Joins("join users on pets.userid = users.id").Where("pets.userid = ?", userid).Scan(&data)

	if err.Error != nil {
		log.Println("error on select data owner", err.Error.Error())
		return nil, nil
	}

	return ParseToArrAdoptionPet(data), ParseToArrApplierPet(dataSeeker)
}

func (ad *adoptionData) GetAdoptionID(adoptID int) []domain.AdoptionPet {
	var data []AdoptionPet

	//mengambil data owner dan pet
	err := ad.db.Model(&Adoption{}).Select("adoptions.id, pets.petname, pets.petphoto, users.fullname, users.photo_profile, users.city, adoptions.status").
		Joins("join pets on adoptions.pets_id = pets.id").Joins("join users on pets.userid = users.id").Where("adoptions.id = ?", adoptID).Scan(&data)

	if err.Error != nil {
		log.Println("problem data", err.Error.Error())
		return nil
	}

	return ParseToArrAdoptionPet(data)
}

func (ad *adoptionData) GetAdoptionbyuser(userID int) []domain.AdoptionPet {
	var data []AdoptionPet

	// mengambil data owner dan pet
	err := ad.db.Model(&Adoption{}).Select("adoptions.id, adoptions.pets_id, pets.petname, pets.petphoto, users.fullname, users.photo_profile, users.city, adoptions.status").
		Joins("join pets on adoptions.pets_id = pets.id").Joins("join users on pets.userid = users.id").Where("adoptions.user_id = ?", userID).Scan(&data)

	if err.Error != nil {
		log.Println("problem data", err.Error.Error())
		return nil
	}

	return ParseToArrAdoptionPet(data)
}
