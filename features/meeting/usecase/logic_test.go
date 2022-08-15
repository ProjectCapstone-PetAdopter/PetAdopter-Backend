package usecase

import (
	"errors"
	"petadopter/domain"
	"petadopter/domain/mocks"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddMeeting(t *testing.T) {
	repo := new(mocks.MockMeetingData)
	insertData := domain.Meeting{
		ID:         1,
		Time:       "09:20:10",
		Date:       "08/08/2022",
		AdoptionID: 1,
		UserID:     1,
	}

	t.Run("Success Insert Meeting", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(1, nil).Once()

		useCase := New(repo, validator.New())

		res, err := useCase.AddMeeting(insertData)
		assert.Nil(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Empty Time", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(-1, errors.New("invalid time")).Once()

		useCase := New(repo, validator.New())
		dummy := insertData
		dummy.Time = ""
		res, err := useCase.AddMeeting(dummy)
		assert.NotNil(t, err)
		assert.Equal(t, -1, res)
		assert.EqualError(t, err, errors.New("invalid time").Error())
	})

	t.Run("Empty Date", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(-1, errors.New("invalid date")).Once()

		useCase := New(repo, validator.New())
		dummy := insertData
		dummy.Date = ""
		res, err := useCase.AddMeeting(dummy)
		assert.NotNil(t, err)
		assert.Equal(t, -1, res)
		assert.EqualError(t, err, errors.New("invalid date").Error())
	})
}

func TestUpdateMeeting(t *testing.T) {
	repo := new(mocks.MockMeetingData)
	insertData := domain.Meeting{
		ID:         1,
		Time:       "09:20:10",
		Date:       "08/08/2022",
		AdoptionID: 1,
		UserID:     1,
	}

	t.Run("Succes Update Meeting", func(t *testing.T) {
		repo.On("Update", mock.Anything, mock.Anything).Return(1, nil).Once()

		useCase := New(repo, validator.New())

		res, err := useCase.UpdateMeeting(insertData, insertData.ID)
		assert.Nil(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Empty Update Time", func(t *testing.T) {
		dummy := insertData
		dummy.Time = ""
		useCase := New(repo, validator.New())

		res, err := useCase.UpdateMeeting(dummy, insertData.ID)

		assert.NotNil(t, err)
		assert.Equal(t, -1, res)
		assert.EqualError(t, err, errors.New("invalid time").Error())
		repo.AssertExpectations(t)
	})

	t.Run("Empty Update Date", func(t *testing.T) {
		dummy := insertData
		dummy.Date = ""
		useCase := New(repo, validator.New())

		res, err := useCase.UpdateMeeting(dummy, insertData.ID)

		assert.NotNil(t, err)
		assert.Equal(t, -1, res)
		assert.EqualError(t, err, errors.New("invalid date").Error())
		repo.AssertExpectations(t)
	})
}

func TestDeleteMeeting(t *testing.T) {
	repo := new(mocks.MockMeetingData)
	insertData := domain.Meeting{
		ID:         1,
		Time:       "09:20:10",
		Date:       "08/08/2022",
		AdoptionID: 1,
		UserID:     1,
	}
	t.Run("Delete Meeting Success", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(nil).Once()

		useCase := New(repo, validator.New())

		err := useCase.DeleteMeeting(insertData.ID)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetMeeting(t *testing.T) {
	repo := new(mocks.MockMeetingData)

	insertData := []domain.MeetingOwner{{
		ID:           1,
		Time:         "09:00:00",
		Date:         "21/08/2022",
		Petname:      "nunu",
		Petphoto:     "nunu.jpg",
		Seekername:   "difa",
		Fullname:     "eunwoo",
		PhotoProfile: "eunwoo.jpg",
		Address:      "surabaya",
	}}

	meetingData := domain.MeetingOwner{
		ID:           1,
		Time:         "09:00:00",
		Date:         "21/08/2022",
		Petname:      "nunu",
		Petphoto:     "nunu.jpg",
		Seekername:   "difa",
		Fullname:     "eunwoo",
		PhotoProfile: "eunwoo.jpg",
		Address:      "surabaya",
	}

	t.Run("Get Meeting Success", func(t *testing.T) {
		repo.On("GetMeetingID", mock.Anything).Return(insertData, nil).Once()

		useCase := New(repo, validator.New())

		res := useCase.GetMyMeeting(meetingData.ID)
		assert.Equal(t, insertData, res)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get Meeting", func(t *testing.T) {
		repo.On("GetMeetingID", mock.Anything).Return([]domain.MeetingOwner{}, errors.New("error get data")).Once()

		useCase := New(repo, validator.New())

		res := useCase.GetMyMeeting(meetingData.ID)
		// assert.NotNil(t, err)
		assert.Equal(t, []domain.MeetingOwner{}, res)
		// assert.EqualError(t, err, errors.New("error get data").Error())
	})
}
