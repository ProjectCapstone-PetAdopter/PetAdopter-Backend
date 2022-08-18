package domain

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

type User struct {
	ID           int
	Username     string
	Fullname     string
	Email        string
	Address      string
	City         string
	PhotoProfile string
	Password     string
	Phonenumber  string
	Role         string
	Pets         []Pets
	Adoption     []Adoption
	Meeting      []Meeting
}

type UserInfo struct {
	Email        string
	Fullname     string
	Photoprofile string
}

type Link struct {
	Url string `json:"url"`
}

type UserHandler interface {
	Login() echo.HandlerFunc
	DeleteUser() echo.HandlerFunc
	Register() echo.HandlerFunc
	Update() echo.HandlerFunc
	GetProfile() echo.HandlerFunc
	GetbyID() echo.HandlerFunc
	CallbackGoogleSignUp() echo.HandlerFunc
	CallbackGoogleLogin() echo.HandlerFunc
	SignUpGoogle() echo.HandlerFunc
	LoginGoogle() echo.HandlerFunc
}

type UserUseCase interface {
	Login(userdata User, token *oauth2.Token) (map[string]interface{}, int)
	Delete(userID int) int
	RegisterUser(newuser User, cost int, token *oauth2.Token, ui UserInfo) int
	UpdateUser(newuser User, userid, cost int) int
	GetProfile(id int) (map[string]interface{}, int)
	GetProfileID(userid int) (map[string]interface{}, int)
}

type UserData interface {
	Login(userdata User, isToken bool) User
	Delete(userID int) int
	RegisterData(newuser User) User
	UpdateUserData(newuser User) User
	CheckDuplicate(newuser User) bool
	GetPasswordData(name string) string
	GetProfile(userID int) (User, int)
	GetProfileIDData(Userid int) (User, int)
}
