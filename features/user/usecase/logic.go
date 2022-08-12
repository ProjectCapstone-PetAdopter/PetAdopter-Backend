package usecase

import (
	"log"

	"petadopter/domain"
	"petadopter/features/common"
	"petadopter/features/user/data"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
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

func (ud *userCase) Login(userdata domain.User) (map[string]interface{}, int) {
	var arrmap = map[string]interface{}{}
	hashpw := ud.userData.GetPasswordData(userdata.Username)

	err := bcrypt.CompareHashAndPassword([]byte(hashpw), []byte(userdata.Password))
	if err != nil {
		log.Println(bcrypt.ErrMismatchedHashAndPassword, err)
		return nil, 400
	}

	login := ud.userData.Login(userdata)
	if login.ID == 0 {
		log.Println("Data login not found")
		return nil, 404
	}

	token := common.GenerateToken(login)

	arrmap["token"] = token
	arrmap["username"] = login.Username
	arrmap["role"] = login.Role

	return arrmap, 200
}

func (ud *userCase) Delete(userId int) int {
	status := ud.userData.Delete(userId)

	if status == 404 {
		log.Println("Cant delete data from query")
		return 404
	}

	if status == 500 {
		log.Println("Cant delete from query")
		return 500
	}

	return 200
}

// RegisterUser implements domain.UserUseCase
func (ud *userCase) RegisterUser(newuser domain.User, cost int, token *oauth2.Token, dataui domain.UserInfo) int {
	var user = data.FromModel(newuser)
	if token != nil {
		user.Email = dataui.Email
		user.Fullname = dataui.Fullname
		user.PhotoProfile = dataui.Photoprofile
	}

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

func (ud *userCase) GetProfile(id int) (map[string]interface{}, int) {
	var res = map[string]interface{}{}
	data, status := ud.userData.GetProfile(id)

	if status == 404 {
		log.Println("No data from query")
		return nil, 404
	}

	if status == 500 {
		log.Println("Cant get from query")
		return nil, 500
	}

	res["username"] = data.Username
	res["fullname"] = data.Fullname
	res["phonenumber"] = data.Phonenumber
	res["email"] = data.Email
	res["address"] = data.Address
	res["photoprofile"] = data.PhotoProfile
	res["city"] = data.City

	return res, 200
}
