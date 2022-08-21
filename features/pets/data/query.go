package data

import (
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

// GetAllPetUser implements domain.PetsData
func (pd *petsData) GetAllPetUser() []domain.PetUser {
	var petuser []PetUser

	err := pd.db.Model(&Pets{}).Select("users.fullname, users.city, species.species").Joins("join users on pets.userid = users.id").Joins("join species on pets.speciesid = species.id").
		Scan(&petuser)

	if err.Error != nil {
		log.Println("cant get petuser data", err.Error.Error())
		return nil
	}

	return ParseToArrPetUser(petuser)
}

// GetPetUser implements domain.PetsData
func (pd *petsData) GetPetUser(userID, petID int) domain.PetUser {
	var petuser PetUser

	if petID != 0 {
		err := pd.db.Model(&Pets{}).Select("users.fullname, users.city, species.species").Joins("join users on pets.userid = users.id").Joins("join species on pets.speciesid = species.id").
			Where("pets.userid = ? and pets.id = ?", userID, petID).Limit(1).Scan(&petuser)

		if err.Error != nil {
			log.Println("cant get petuser data", err.Error.Error())
			return domain.PetUser{}
		}
	} else {
		err := pd.db.Model(&Pets{}).Select("users.fullname, users.city, species.species").Joins("join users on pets.userid = users.id").Joins("join species on pets.speciesid = species.id").
			Where("pets.userid = ?", userID).Limit(1).Scan(&petuser)

		if err.Error != nil {
			log.Println("cant get petuser data", err.Error.Error())
			return domain.PetUser{}
		}
	}

	return petuser.ToDomainPetUser()
}

func (pd *petsData) Insert(newPets domain.Pets) domain.Pets {
	cnv := ToLocal(newPets)

	err := pd.db.Create(&cnv)
	if err.Error != nil {
		return domain.Pets{}
	}

	return cnv.ToDomain()
}

func (pd *petsData) Update(petsID int, updatedProduct domain.Pets) domain.Pets {
	cnv := ToLocal(updatedProduct)

	log.Println(cnv)
	err := pd.db.Model(cnv).Where("ID = ? AND userid = ?", petsID, cnv.Userid).Updates(cnv)
	if err.Error != nil {
		log.Println("Cannot update data in db", err.Error.Error())
		return domain.Pets{}
	}

	if err.RowsAffected == 0 {
		log.Println("Data not found")
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

	err := pd.db.Model(&Pets{}).Select("DISTINCT pets.id, pets.petname, pets.gender, pets.age, pets.color, pets.description, pets.petphoto, pets.speciesid, pets.userid").Joins("left join adoptions on pets.id = adoptions.pets_id").
		Where("adoptions.status is null or adoptions.status = 'Requested' or adoptions.status = 'Rejected'").Scan(&data)
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

	err := pd.db.Where("userid = ?", userID).Find(&data)
	if err.Error != nil {
		log.Println("problem data", err.Error.Error())
		return nil
	}

	return ParseToArr(data)
}
