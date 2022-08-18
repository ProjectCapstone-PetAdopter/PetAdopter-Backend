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
	AdoptionID   int
	UserID       int
	Time         string
	Date         string
	Petname      string
	Petphoto     string
	Seekerid     int
	Seekername   string
	Fullname     string
	PhotoProfile string
	Address      string
}

type Ownerdata struct {
	Email   string
	Address string
	City    string
}

type Seekerdata struct {
	UserID   int
	Fullname string
	Email    string
}

type MeetingHandler interface {
	InsertMeeting() echo.HandlerFunc
	UpdateDataMeeting() echo.HandlerFunc
	DeleteDataMeeting() echo.HandlerFunc
	GetMeeting() echo.HandlerFunc
	GetMyMeeting() echo.HandlerFunc
}

type MeetingUsecase interface {
	AddMeeting(data Meeting) (idMeet int, err error)
	UpdateMeeting(UpdateMeeting Meeting, id int) (idMeet int, err error)
	DeleteMeeting(id int) error
	GetOwnerMeeting(userID int) (arrmap []map[string]interface{}, err error)
	GetSeekMeeting(userID int) (arrmap []map[string]interface{}, err error)
	GetEmail(userID, meetingID int) (Ownerdata, Seekerdata)
}

type MeetingData interface {
	Insert(data Meeting) (idMeet int, err error)
	Update(updatedData Meeting, id int) (idMeet int, err error)
	Delete(id int) error
	GetMeetingID(userID int) []MeetingOwner
	GetMyMeetingID(userID int) []MeetingOwner
	GetEmailData(userID, meetingID int) (Ownerdata, Seekerdata, int)
}
