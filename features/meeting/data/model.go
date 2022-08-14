package data

import (
	"petadopter/domain"

	"gorm.io/gorm"
)

type Meeting struct {
	gorm.Model
	Time       string `json:"time" form:"time"`
	Date       string `json:"date" form:"date"`
	AdoptionID int
	UserID     int
}

type MeetingOwner struct {
	ID           int
	Time         string
	Date         string
	Petname      string
	Petphoto     string
	Seekername   string
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
		AdoptionID: m.AdoptionID,
		UserID:     m.UserID,
	}
}

func (m *MeetingOwner) ToModelMeeting() domain.MeetingOwner {
	return domain.MeetingOwner{
		ID:           m.ID,
		Time:         m.Time,
		Date:         m.Date,
		Petname:      m.Petname,
		Petphoto:     m.Petphoto,
		Seekername:   m.Seekername,
		Fullname:     m.Fullname,
		PhotoProfile: m.PhotoProfile,
		Address:      m.Address,
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
	res.AdoptionID = data.AdoptionID
	res.UserID = data.UserID
	return res
}
