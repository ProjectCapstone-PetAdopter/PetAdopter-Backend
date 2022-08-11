package delivery

import "petadopter/domain"

type InsertMeeting struct {
	Time       string `json:"time" form:"time"`
	Date       string `json:"date" form:"date"`
	AdoptionID uint   `json:"adoptionid" form:"adoptionid"`
}

func (i InsertMeeting) ToModel() domain.Meeting {
	return domain.Meeting{
		Time:       i.Time,
		Date:       i.Date,
		AdoptionID: int(i.AdoptionID),
	}
}
