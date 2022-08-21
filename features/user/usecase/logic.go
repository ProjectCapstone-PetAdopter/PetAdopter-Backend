package usecase

import (
	"fmt"
	"log"
	"mime/multipart"

	"petadopter/config"
	"petadopter/domain"
	"petadopter/features/common"
	"petadopter/features/user/data"
	"petadopter/utils/google"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
)

type userCase struct {
	userData domain.UserData
	valid    *validator.Validate
	client   *google.ClientUploader
}

func New(ud domain.UserData, val *validator.Validate, cl *google.ClientUploader) domain.UserUseCase {
	return &userCase{
		userData: ud,
		valid:    val,
		client:   cl,
	}
}

// GetProfileID implements domain.UserUseCase
func (ud *userCase) GetProfileID(userid int) (map[string]interface{}, int) {
	var res = map[string]interface{}{}
	data, status := ud.userData.GetProfileIDData(userid)

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

func (ud *userCase) Login(userdata domain.User, authtoken *oauth2.Token) (map[string]interface{}, int) {
	var resMap = map[string]interface{}{}
	var tokenOauth string

	isToken := true
	if authtoken == nil {
		hashPw := ud.userData.GetPasswordData(userdata.Username)
		err := bcrypt.CompareHashAndPassword([]byte(hashPw), []byte(userdata.Password))
		if err != nil {
			log.Println(bcrypt.ErrMismatchedHashAndPassword, err)
			return nil, 400
		}
		isToken = false
		tokenOauth = ""
	}

	if authtoken != nil {
		tokenOauth = authtoken.AccessToken
	}

	login := ud.userData.Login(userdata, isToken)
	if login.ID == 0 {
		log.Println("Data login not found")
		return nil, 404
	}

	tokenjwt := common.GenerateToken(login)
	resMap["tokenoauth"] = tokenOauth
	resMap["token"] = tokenjwt
	resMap["username"] = login.Username
	resMap["role"] = login.Role

	return resMap, 200
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
func (ud *userCase) RegisterUser(newuser domain.User, cost int, token *oauth2.Token, dataui domain.UserInfo) (map[string]interface{}, int) {
	var resMap = map[string]interface{}{}
	var user = data.FromModel(newuser)

	if token != nil {
		user.Email = dataui.Email
		user.Fullname = dataui.Fullname
		user.PhotoProfile = dataui.Photoprofile
		user.Username = dataui.Fullname
		tokenOauth := token.AccessToken
		resMap["tokenoauth"] = tokenOauth
	}

	if token == nil {
		validError := ud.valid.Struct(user)
		if validError != nil {
			log.Println("Validation errror : ", validError)
			return nil, 400
		}
	}

	duplicate := ud.userData.CheckDuplicate(user.ToModel())
	if duplicate {
		log.Println("Duplicate Data")
		return nil, 409
	}

	hashed, hasherr := bcrypt.GenerateFromPassword([]byte(user.Password), cost)
	if hasherr != nil {
		log.Println("Cant encrypt: ", hasherr)
		return nil, 500
	}

	user.Password = string(hashed)
	user.Role = "user"
	insert := ud.userData.RegisterData(user.ToModel())

	if insert.ID == 0 {
		log.Println("Empty data")
		return nil, 404
	}

	tokenjwt := common.GenerateToken(insert)

	resMap["token"] = tokenjwt
	resMap["role"] = insert.Role

	return resMap, 200
}

// UpdateUser implements domain.UserUseCase
func (ud *userCase) UpdateUser(newuser domain.User, userid, cost int, form *multipart.FileHeader) int {
	var user = data.FromModel(newuser)
	user.ID = uint(userid)

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

	if form != nil {
		file, err := form.Open()
		if err != nil {
			log.Println(err, "cant open file")
			return 500
		}

		defer file.Close()
		id := uuid.New()
		filename := fmt.Sprintf("%sPP-%s.jpg", newuser.Username, id.String())
		config.UPLOADPATH = "profile/"

		link, err := ud.client.UploadFile(file, config.UPLOADPATH, filename)
		if err != nil {
			log.Println(err, "cant upload file")
			return 500
		}
		user.PhotoProfile = link
	}

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
