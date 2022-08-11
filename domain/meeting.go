package domain

import "github.com/labstack/echo/v4"

type Meeting struct {
	ID         int
	Time       string
	Date       string
	AdoptionID int
	UserID     int
}

type MeetingHandler interface {
	InsertMeeting() echo.HandlerFunc
	UpdateDataMeeting() echo.HandlerFunc
	DeleteDataMeeting() echo.HandlerFunc
}

type MeetingUsecase interface {
	AddMeeting(data Meeting) error
	UpdateMeeting(UpdateMeeting Meeting, id int) error
	DeleteMeeting(id int) error
}

type MeetingData interface {
	Insert(data Meeting) error
	Update(updatedData Meeting, id int) error
	Delete(id int) error
}
