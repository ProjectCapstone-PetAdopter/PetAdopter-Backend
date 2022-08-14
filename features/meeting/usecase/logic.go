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
	var tmp delivery.InsertMeeting
	if tmp.Time != "" {
		return -1, errors.New("invalid time")
	}
	if tmp.Date != "" {
		return -1, errors.New("invalid date")
	}
	inserted, err := mu.meetingData.Update(UpdateMeeting, id)
	return inserted, nil
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
