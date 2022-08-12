package usecase

import (
	"errors"
	"petadopter/domain"
	"petadopter/features/meeting/delivery"

	"github.com/go-playground/validator/v10"
)

type meetingUsecase struct {
	meetingData domain.MeetingData
	validate    *validator.Validate
}

func New(md domain.MeetingData, v *validator.Validate) domain.MeetingUsecase {
	return &meetingUsecase{
		meetingData: md,
		validate:    v,
	}
}

func (mu *meetingUsecase) AddMeeting(data domain.Meeting) (row int, err error) {
	if data.Time == "" {
		return -1, errors.New("invalid time")
	}
	if data.Date == "" {
		return -1, errors.New("invalid date")
	}
	inserted, err := mu.meetingData.Insert(data)
	return inserted, err
}

func (mu *meetingUsecase) UpdateMeeting(UpdateMeeting domain.Meeting, id int) error {
	var tmp delivery.InsertMeeting
	qry := map[string]interface{}{}
	if tmp.Time != "" {
		qry["time"] = &tmp.Time
	}
	if tmp.Date != "" {
		qry["date"] = &tmp.Date
	}
	err := mu.meetingData.Update(UpdateMeeting, id)
	return err
}

func (mu *meetingUsecase) DeleteMeeting(id int) error {
	err := mu.meetingData.Delete(id)
	return err
}

func (mu *meetingUsecase) GetMyMeeting(meetingID int) ([]domain.MeetingOwner, error) {
	res := mu.meetingData.GetMeetingID(meetingID)
	if meetingID == -1 {
		return nil, errors.New("error get data")
	}
	return res, nil
}
