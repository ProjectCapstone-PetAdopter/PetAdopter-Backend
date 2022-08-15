package domain

import "github.com/labstack/echo/v4"

type Meeting struct {
	ID         int
	Time       string
	Date       string
	AdoptionID int
	UserID     int
}

type MeetingOwner struct {
	ID           int    `json:"id"`
	Time         string `json:"time"`
	Date         string `json:"date"`
	Petname      string `json:"petname"`
	Petphoto     string `json:"petphoto"`
	Seekername   string `json:"seekername"`
	Fullname     string `json:"fullname"`
	PhotoProfile string `json:"photoprofile"`
	Address      string `json:"address"`
}

type MeetingHandler interface {
	InsertMeeting() echo.HandlerFunc
	UpdateDataMeeting() echo.HandlerFunc
	DeleteDataMeeting() echo.HandlerFunc
	GetAdopt() echo.HandlerFunc
}

type MeetingUsecase interface {
	AddMeeting(data Meeting) (idMeet int, err error)
	UpdateMeeting(UpdateMeeting Meeting, id int) (idMeet int, err error)
	DeleteMeeting(id int) error
	GetMyMeeting(meetingID int) []MeetingOwner
}

type MeetingData interface {
	Insert(data Meeting) (idMeet int, err error)
	Update(updatedData Meeting, id int) (idMeet int, err error)
	Delete(id int) error
	GetMeetingID(meetingID int) []MeetingOwner
}
