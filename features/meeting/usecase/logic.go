package usecase

import (
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

func (mu *meetingUsecase) AddMeeting(data domain.Meeting) error {
	err := mu.meetingData.Insert(data)
	return err
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
