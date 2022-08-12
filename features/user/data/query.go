package data

import (
	"log"

	"petadopter/domain"

	"gorm.io/gorm"
)

type userData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.UserData {
	return &userData{
		db: db,
	}
}

// GetUserCartData implements domain.UserData
func (ud *userData) Login(userdata domain.User) domain.User {
	var user = FromModel(userdata)
	err := ud.db.First(&user, "username  = ?", userdata.Username).Error

	if err != nil {
		log.Println("Cant login data", err.Error())
		return domain.User{}
	}

	return user.ToModel()
}

func (ud *userData) Delete(userID int) int {
	err := ud.db.Where("ID = ?", userID).Delete(&User{})

	if err.Error != nil {
		log.Println("cannot delete data", err.Error.Error())
		return 500
	}
	if err.RowsAffected < 1 {
		log.Println("No data deleted", err.Error.Error())
		return 404
	}

	return 200
}

// RegisterData implements domain.UserData
func (ud *userData) RegisterData(newuser domain.User) domain.User {
	var user = FromModel(newuser)
	err := ud.db.Create(&user).Error

	if user.ID == 0 {
		log.Println("Invalid ID")
		return domain.User{}
	}

	if err != nil {
		log.Println("Cant create user object", err.Error())
		return domain.User{}
	}

	return user.ToModel()
}

// UpdateUserData implements domain.UserData
func (ud *userData) UpdateUserData(newuser domain.User) domain.User {
	var user = FromModel(newuser)
	err := ud.db.Model(&User{}).Where("ID = ?", user.ID).Updates(user)

	if err.Error != nil {
		log.Println("Cant update user object", err.Error.Error())
		return domain.User{}
	}

	if err.RowsAffected == 0 {
		log.Println("Data Not Found")
		return domain.User{}
	}

	return user.ToModel()
}

// CheckDuplicate implements domain.UserData
func (ud *userData) CheckDuplicate(newuser domain.User) bool {
	var user User
	err := ud.db.Find(&user, "username = ? OR email = ?", newuser.Username, newuser.Email)

	if err.RowsAffected == 1 {
		log.Println("Duplicated data", err.Error)
		return true
	}

	return false
}

// GetPasswordData implements domain.UserData
func (ud *userData) GetPasswordData(name string) string {
	var user User
	err := ud.db.Find(&user, "username = ?", name).Error

	if err != nil {
		log.Println("Cant retrieve user data", err.Error())
		return ""
	}

	return user.Password
}

func (ud *userData) GetProfile(userID int) (domain.User, int) {
	var tmp User

	err := ud.db.Where("ID = ?", userID).First(&tmp)
	if err.Error != nil {
		log.Println("Cant get data", err.Error)
		return domain.User{}, 500
	}

	if err.RowsAffected == 0 {
		log.Println("Rows affcted = 0", err.Error)
		return domain.User{}, 404
	}

	return tmp.ToModel(), 200
}
