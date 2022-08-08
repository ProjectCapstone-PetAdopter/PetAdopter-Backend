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
func (ud *userData) GetUserCartData(userid int) []domain.UserCart {
	var usercart []domain.UserCart

	err := ud.db.Model(&User{}).Select("products.product_name, carts.quantity, carts.subtotal").Joins("join carts on carts.userid = users.id").
		Joins("join products on carts.productid = products.id").Where("carts.userid = ?", userid).Find(&usercart)

	if err.Error != nil {
		log.Println("Cant get data", err.Error)
		return nil
	}

	return usercart
}

func (ud *userData) Login(userdata domain.User) domain.User {
	var user = FromModel(userdata)
	err := ud.db.First(&user, "username  = ?", userdata.Username).Error

	if err != nil {
		log.Println("Cant login data", err.Error())
		return domain.User{}
	}

	return user.ToModel()
}

func (ud *userData) Delete(userID int) bool {
	err := ud.db.Where("ID = ?", userID).Delete(&User{})

	if err.Error != nil {
		log.Println("cannot delete data", err.Error.Error())
		return false
	}
	if err.RowsAffected < 1 {
		log.Println("No content deleted", err.Error.Error())
		return false
	}

	return true
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

func (ud *userData) GetProfile(userID int) (domain.User, error) {
	var tmp User
	err := ud.db.Where("ID = ?", userID).First(&tmp).Error
	if err != nil {
		log.Println("There is a problem with data", err.Error())
		return domain.User{}, err
	}

	return tmp.ToModel(), nil
}
