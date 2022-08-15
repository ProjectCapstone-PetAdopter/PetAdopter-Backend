package usecase

import (
	"errors"
	"petadopter/domain"

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

func (mu *meetingUsecase) AddMeeting(data domain.Meeting) (idMeet int, err error) {
	if data.Time == "" {
		return -1, errors.New("invalid time")
	}
	if data.Date == "" {
		return -1, errors.New("invalid date")
	}
	inserted, err := mu.meetingData.Insert(data)
	return inserted, err
}

func (mu *meetingUsecase) UpdateMeeting(UpdateMeeting domain.Meeting, id int) (idMeet int, err error) {
	// var tmp delivery.InsertMeeting
	if UpdateMeeting.Time == "" {
		return -1, errors.New("invalid time")
	}
	if UpdateMeeting.Date == "" {
		return -1, errors.New("invalid date")
	}
	inserted, err := mu.meetingData.Update(UpdateMeeting, id)
	return inserted, err
}

func (mu *meetingUsecase) DeleteMeeting(id int) error {
	err := mu.meetingData.Delete(id)
	return err
}

func (mu *meetingUsecase) GetMyMeeting(meetingID int) ([]domain.MeetingOwner, error) {

	data := mu.meetingData.GetMeetingID(meetingID)
	if meetingID == -1 {
		return []domain.MeetingOwner{}, errors.New("error get data")

	}
	return data, nil
}
