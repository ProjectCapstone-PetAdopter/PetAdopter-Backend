package delivery

import "petadopter/domain"

type InsertMeeting struct {
	Time       string `json:"time" form:"time"`
	Date       string `json:"date" form:"date"`
	Token      string `json:"token" form:"token"`
	AdoptionID int    `json:"adoptionid" form:"adoptionid"`
	Userid     int
}

func (i InsertMeeting) ToModel() domain.Meeting {
	return domain.Meeting{
		Time:       i.Time,
		Date:       i.Date,
		UserID:     i.Userid,
		AdoptionID: i.AdoptionID,
	}
}
