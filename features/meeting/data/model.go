package data

import (
	"petadopter/domain"

	"gorm.io/gorm"
)

type Meeting struct {
	gorm.Model
	Time       string
	Date       string
	AdoptionID int
	UserID     int
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

func ParseToArr(arr []Meeting) []domain.Meeting {
	var res []domain.Meeting

	for _, val := range arr {
		res = append(res, val.ToModel())
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
