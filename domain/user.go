package domain

import "github.com/labstack/echo/v4"

type User struct {
	ID           int
	Username     string
	Email        string
	Address      string
	PhotoProfile string
	Password     string
	Role         string
}

type UserCart struct {
	ProductName string
	Quantity    int
	Subtotal    int
}

type UserHandler interface {
	Login() echo.HandlerFunc
	DeleteUser() echo.HandlerFunc
	Register() echo.HandlerFunc
	Update() echo.HandlerFunc
	GetProfile() echo.HandlerFunc
}

type UserUseCase interface {
	Login(userdata User) (User, error)
	Delete(userID int) (bool, error)
	RegisterUser(newuser User, cost int) int
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
