package data

import (
	"petadopter/domain"

	"gorm.io/gorm"
)

type Meeting struct {
	gorm.Model
	Time        string `json:"time" form:"time"`
	Date        string `json:"date" form:"date"`
	Adoption_id int
	Adoption    Adoption
	UserID      int
}

type Adoption struct {
	gorm.Model
	UserID   int
	PetsID   int       `json:"petid" form:"petid"`
	Status   string    `gorm:"default:waiting"`
	Petphoto string    `json:"petphoto" form:"petphoto"`
	Meeting  []Meeting `gorm:"foreignKey:Adoption_id"`
}

type MeetingOwner struct {
	ID           int
	Petname      string
	Petphoto     string
	Fullname     string
	PhotoProfile string
	Address      string
	Status       string
}

func (m *Meeting) ToModel() domain.Meeting {
	return domain.Meeting{
		ID:         int(m.ID),
		Time:       m.Time,
		Date:       m.Date,
		AdoptionID: m.Adoption_id,
		UserID:     m.UserID,
	}
}

func (m *MeetingOwner) ToModelMeeting() domain.MeetingOwner {
	return domain.MeetingOwner{
		ID:           m.ID,
		Petname:      m.Petname,
		Petphoto:     m.Petphoto,
		Fullname:     m.Fullname,
		PhotoProfile: m.PhotoProfile,
		Address:      m.Address,
		Status:       m.Status,
	}
}

func ParseToArr(arr []Meeting) []domain.Meeting {
	var res []domain.Meeting

	for _, val := range arr {
		res = append(res, val.ToModel())
	}
	return res
}

func ParseToArrMeeting(arr []MeetingOwner) []domain.MeetingOwner {
	var res []domain.MeetingOwner

	for _, val := range arr {
		res = append(res, val.ToModelMeeting())
	}
	return res
}

func FromModel(data domain.Meeting) Meeting {
	var res Meeting
	res.Time = data.Time
	res.Date = data.Date
	res.Adoption_id = data.AdoptionID
	res.UserID = data.UserID
	return res
}
