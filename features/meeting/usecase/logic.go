package usecase

import (
	"errors"
	"log"
	"petadopter/domain"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
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

func (mh *meetingUsecase) GetPetMeeting(id int) (domain.Meeting, error) {
	data, err := mh.meetingData.GetMyMeetingPets(id)

	if err != nil {
		log.Println("use case", err.Error())
		if err == gorm.ErrRecordNotFound {
			return domain.Meeting{}, errors.New("data not found")
		} else {
			return domain.Meeting{}, errors.New("server error")
		}
	}
	return data, nil
}

// GetEmail implements domain.MeetingUsecase
func (mh *meetingUsecase) GetEmail(userID, meetingID int) (domain.Ownerdata, domain.Seekerdata) {
	owner, seeker, status := mh.meetingData.GetEmailData(userID, meetingID)

	if status == 404 {
		log.Println("data not found")
		return domain.Ownerdata{}, domain.Seekerdata{}
	}

	if status == 500 {
		log.Println("error in query")
		return domain.Ownerdata{}, domain.Seekerdata{}
	}

	return owner, seeker
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

func (mu *meetingUsecase) GetMyMeeting(meetingID int) (getMyData []domain.MeetingOwner, err error) {

	data := mu.meetingData.GetMeetingID(meetingID)

	if meetingID < 1 {
		return []domain.MeetingOwner{}, errors.New("error get data")
	}
	return data, err
}
