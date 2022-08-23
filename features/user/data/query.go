package data

import (
	"log"
	"time"

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

// GetProfileIDData implements domain.UserData
func (ud *userData) GetProfileIDData(userID int) (domain.User, int) {
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

// GetUserCartData implements domain.UserData
func (ud *userData) Login(dataLogin domain.User, isToken bool) domain.User {
	log.Println(isToken)
	var user = FromModel(dataLogin)
	var err error
	if isToken {
		err = ud.db.First(&user, "email  = ?", dataLogin.Email).Error
	} else {
		err = ud.db.First(&user, "username  = ?", dataLogin.Username).Error
	}
	if err != nil {
		log.Println("Cant login data", err.Error())
		return domain.User{}
	}

	return user.ToModel()
}

func (ud *userData) Delete(userID int) int {
	time := time.Now()

	err := ud.db.Where("ID = ?", userID).Delete(&User{})

	if err.Error != nil {
		log.Println("cannot delete data", err.Error)
		return 500
	}
	if err.RowsAffected < 1 {
		log.Println("No data deleted", err.Error)
		return 404
	}

	deleteData := ud.db.Exec("UPDATE pets SET pets.deleted_at=? WHERE pets.userid = ?;", time, userID)

	if deleteData.Error != nil {
		log.Println("cannot delete data", deleteData.Error)
		return 500
	}

	if deleteData.RowsAffected < 1 {
		log.Println("No data pets deleted", deleteData.Error)
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
	var user = FromModel(newuser)
	var check User
	log.Println(user.ID)
	if user.ID > 0 {
		err := ud.db.Where("username = ? AND ID != ?", user.Username, user.ID).Find(&check)
		if err.RowsAffected == 1 {
			log.Println("Duplicated data", err.Error)
			return true
		}

		err = ud.db.Where("email = ? AND ID != ?", user.Email, user.ID).Find(&check)
		if err.RowsAffected == 1 {
			log.Println("Duplicated data", err.Error)
			return true
		}

	} else {
		err := ud.db.Find(&user, "username = ? OR email = ?", user.Username, user.Email)
		if err.RowsAffected == 1 {
			log.Println("Duplicated data", err.Error)
			return true
		}
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
