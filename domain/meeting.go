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
	ID           int
	Petname      string
	Petphoto     string
	Fullname     string
	PhotoProfile string
	Address      string
	Status       string
}

type MeetingHandler interface {
	InsertMeeting() echo.HandlerFunc
	UpdateDataMeeting() echo.HandlerFunc
	DeleteDataMeeting() echo.HandlerFunc
	GetAdopt() echo.HandlerFunc
}

type MeetingUsecase interface {
	AddMeeting(data Meeting) (row int, err error)
	UpdateMeeting(UpdateMeeting Meeting, id int) error
	DeleteMeeting(id int) error
	GetMyMeeting(meetingID int) ([]MeetingOwner, error)
}

type MeetingData interface {
	Insert(data Meeting) (row int, err error)
	Update(updatedData Meeting, id int) error
	Delete(id int) error
	GetMeetingID(meetingID int) []MeetingOwner
}
