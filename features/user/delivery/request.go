package delivery

import "petadopter/domain"

type UserFormat struct {
	Username     string `json:"username" form:"username" validate:"required"`
	Email        string `json:"email" form:"email" validate:"required,email"`
	Address      string `json:"address" form:"address" validate:"required"`
	City         string `json:"city" form:"city" validate:"required"`
	Password     string `json:"password" form:"password" validate:"required"`
	Fullname     string `json:"fullname" form:"fullname" validate:"required"`
	Phonenumber  string `json:"phonenumber" form:"phonenumber" validate:"required"`
	Photoprofile string `json:"photoprofile" form:"photoprofile"`
}

func (i *UserFormat) ToModel() domain.User {
	return domain.User{
		Username:     i.Username,
		Email:        i.Email,
		Address:      i.Address,
		City:         i.City,
		Password:     i.Password,
		Fullname:     i.Fullname,
		Phonenumber:  i.Phonenumber,
		PhotoProfile: i.Photoprofile,
	}
}

type LoginFormat struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func (i *LoginFormat) ToModelLogin() domain.User {
	return domain.User{
		Username: i.Username,
		Password: i.Password,
	}
}

type UpdateFormat struct {
	Username     string `json:"username" form:"username" validate:"required"`
	Email        string `json:"email" form:"email" validate:"required,email"`
	Address      string `json:"address" form:"address" validate:"required"`
	City         string `json:"city" form:"city" validate:"required"`
	Password     string `json:"password" form:"password" validate:"required"`
	Photoprofile string `json:"photoprofile" form:"photoprofile"`
	Phonenumber  string `json:"phonenumber" form:"phonenumber" validate:"required"`
	Fullname     string `json:"fullname" form:"fullname" validate:"required"`
}

func (i *UpdateFormat) ToModelUpdate() domain.User {
	return domain.User{
		Username:     i.Username,
		Email:        i.Email,
		Address:      i.Address,
		City:         i.City,
		Password:     i.Password,
		PhotoProfile: i.Photoprofile,
		Fullname:     i.Fullname,
		Phonenumber:  i.Phonenumber,
	}
}

type UserInfoFormat struct {
	Email        string `json:"email"`
	Fullname     string `json:"name"`
	Photoprofile string `json:"picture"`
}

func (i *UserInfoFormat) ToModelUserInfoFormat() domain.User {
	return domain.User{
		Email:        i.Email,
		Fullname:     i.Fullname,
		PhotoProfile: i.Photoprofile,
	}
}
