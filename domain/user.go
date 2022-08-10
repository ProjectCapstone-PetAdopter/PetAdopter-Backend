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
}

type UserHandler interface {
	Login() echo.HandlerFunc
	DeleteUser() echo.HandlerFunc
	Register() echo.HandlerFunc
	Update() echo.HandlerFunc
	GetProfile() echo.HandlerFunc
	CallbackGoogle() echo.HandlerFunc
	SignUpGoogle() echo.HandlerFunc
}

type UserUseCase interface {
	Login(userdata User) (User, error)
	Delete(userID int) (bool, error)
	RegisterUser(newuser User, cost int, token *oauth2.Token) int
	UpdateUser(newuser User, userid, cost int) int
	GetProfile(id int) (User, error)
}

type UserData interface {
	Login(userdata User) User
	Delete(userID int) bool
	RegisterData(newuser User) User
	UpdateUserData(newuser User) User
	CheckDuplicate(newuser User) bool
	GetPasswordData(name string) string
	GetProfile(userID int) (User, error)
}
