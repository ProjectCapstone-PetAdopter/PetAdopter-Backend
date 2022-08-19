package delivery

import "petadopter/domain"

type MeetingResponse struct {
	ID   int    `json:"meetingid"`
	Time string `json:"time"`
	Date string `json:"date"`
}

func FromModel(data domain.Meeting) MeetingResponse {
	var res MeetingResponse
	res.ID = data.ID
	res.Time = data.Time
	res.Date = data.Date
	return res
}
