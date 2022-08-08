package delivery

import "petadopter/domain"

type UserFormat struct {
	Username     string `json:"username" form:"username" validate:"required"`
	Email        string `json:"email" form:"email" validate:"required,email"`
	Address      string `json:"address" form:"address" validate:"required"`
	Password     string `json:"password" form:"password" validate:"required"`
	Photoprofile string `json:"photoprofile" form:"photoprofile"`
}

func (i *UserFormat) ToModel() domain.User {
	return domain.User{
		Username:     i.Username,
		Email:        i.Email,
		Address:      i.Address,
		Password:     i.Password,
		PhotoProfile: i.Photoprofile,
	}
}

type LoginFormat struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (i *LoginFormat) ToModelLogin() domain.User {
	return domain.User{
		Username: i.Username,
		Password: i.Password,
	}
}
