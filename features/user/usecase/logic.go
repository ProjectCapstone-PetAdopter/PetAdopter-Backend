package usecase

import (
	"errors"
	"log"

	"petadopter/domain"
	"petadopter/features/user/data"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type userCase struct {
	userData domain.UserData
	valid    *validator.Validate
}

func New(ud domain.UserData, val *validator.Validate) domain.UserUseCase {
	return &userCase{
		userData: ud,
		valid:    val,
	}
}

func (ud *userCase) Login(userdata domain.User) (domain.User, error) {

	hashpw := ud.userData.GetPasswordData(userdata.Username)

	err := bcrypt.CompareHashAndPassword([]byte(hashpw), []byte(userdata.Password))

	if err != nil {
		log.Println(bcrypt.ErrMismatchedHashAndPassword, err)
		return domain.User{}, err
	}

	login := ud.userData.Login(userdata)

	if login.ID == 0 {
		return domain.User{}, errors.New("no data")
	}

	return login, nil
}

func (ud *userCase) Delete(userId int) (bool, error) {
	res := ud.userData.Delete(userId)

	if !res {
		return false, errors.New("failed to delete user")
	}
	return true, nil
}

// RegisterUser implements domain.UserUseCase
func (ud *userCase) RegisterUser(newuser domain.User, cost int, token *oauth2.Token) int {
	var user = data.FromModel(newuser)

	if token == nil {
		validError := ud.valid.Struct(user)
		if validError != nil {
			log.Println("Validation errror : ", validError)
			return 400
		}
	}

	duplicate := ud.userData.CheckDuplicate(user.ToModel())
	if duplicate {
		log.Println("Duplicate Data")
		return 409
	}

	hashed, hasherr := bcrypt.GenerateFromPassword([]byte(user.Password), cost)
	if hasherr != nil {
		log.Println("Cant encrypt: ", hasherr)
		return 500
	}

	user.Password = string(hashed)
	user.Role = "user"
	insert := ud.userData.RegisterData(user.ToModel())

	if insert.ID == 0 {
		log.Println("Empty data")
		return 404
	}

	return 200
}

// UpdateUser implements domain.UserUseCase
func (ud *userCase) UpdateUser(newuser domain.User, userid int, cost int) int {
	var user = data.FromModel(newuser)

	if userid == 0 {
		log.Println("Data not found")
		return 404
	}

	duplicate := ud.userData.CheckDuplicate(user.ToModel())

	if duplicate {
		log.Println("Duplicate Data")
		return 409
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), cost)

	if err != nil {
		log.Println("Error encrypt password", err)
		return 500
	}

	user.ID = uint(userid)
	user.Password = string(hashed)

	ud.userData.UpdateUserData(user.ToModel())

	return 200
}

func (ud *userCase) GetProfile(id int) (domain.User, error) {
	data, err := ud.userData.GetProfile(id)

	if err != nil {
		log.Println("Use case", err.Error())
		if err == gorm.ErrRecordNotFound {
			return domain.User{}, errors.New("data not found")
		} else {
			return domain.User{}, errors.New("server error")
		}
	}

	return data, nil
}
