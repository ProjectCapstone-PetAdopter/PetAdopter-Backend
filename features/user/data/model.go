package data

import (
	"petadopter/domain"
	adoptiondata "petadopter/features/adoption/data"
	meetingData "petadopter/features/meeting/data"
	"petadopter/features/pets/data"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `json:"username" form:"username" validate:"required"`
	Fullname     string `json:"fullname" form:"fullname" validate:"required"`
	Email        string `json:"email" form:"email" validate:"required"`
	Address      string `json:"address" form:"address" validate:"required"`
	City         string `json:"city" form:"city" validate:"required"`
	PhotoProfile string `json:"photoprofile" form:"photoprofile"`
	Password     string `json:"password" form:"password" validate:"required"`
	Phonenumber  string `json:"phonenumber" form:"phonenumber" validate:"required"`
	Role         string
	Pets         []data.Pets             `gorm:"foreignKey:Userid"`
	Adoption     []adoptiondata.Adoption `gorm:"foreignKey:UserID"`
	Meeting      []meetingData.Meeting   `gorm:"foreignKey:UserID"`
}

type UserInfo struct {
	Email        string `json:"email"`
	Fullname     string `json:"name"`
	Photoprofile string `json:"picture"`
}

func (u *User) ToModel() domain.User {
	return domain.User{
		ID:           int(u.ID),
		Username:     u.Username,
		Fullname:     u.Fullname,
		Email:        u.Email,
		Address:      u.Address,
		City:         u.City,
		PhotoProfile: u.PhotoProfile,
		Password:     u.Password,
		Phonenumber:  u.Phonenumber,
		Role:         u.Role,
	}
}

func (u *UserInfo) ToModelUserInfo() domain.UserInfo {
	return domain.UserInfo{
		Email:        u.Email,
		Fullname:     u.Fullname,
		Photoprofile: u.Photoprofile,
	}
}

func ParseToArr(arr []User) []domain.User {
	var res []domain.User

	for _, val := range arr {
		res = append(res, val.ToModel())
	}
	return res
}

func FromModel(data domain.User) User {
	var res User
	res.ID = uint(data.ID)
	res.Username = data.Username
	res.Fullname = data.Fullname
	res.Email = data.Email
	res.Address = data.Address
	res.City = data.City
	res.PhotoProfile = data.PhotoProfile
	res.Password = data.Password
	res.Phonenumber = data.Phonenumber
	res.Role = data.Role
	return res
}
