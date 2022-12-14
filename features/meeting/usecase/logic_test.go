package usecase

import (
	"errors"
	"petadopter/domain"
	"petadopter/domain/mocks"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestAddMeeting(t *testing.T) {
	repo := new(mocks.MeetingData)
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
	repo := new(mocks.MeetingData)
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
	repo := new(mocks.MeetingData)
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
	repo := new(mocks.MeetingData)

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

		res, _ := useCase.GetOwnerMeeting(meetingData.ID)
		// assert.Nil(t, err)
		assert.GreaterOrEqual(t, 1, len(res))
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get Meeting", func(t *testing.T) {
		repo.On("GetMeetingID", mock.Anything).Return([]domain.MeetingOwner{}, errors.New("error get data")).Once()

		useCase := New(repo, validator.New())

		res, err := useCase.GetOwnerMeeting(-1)
		// assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, errors.New("error get data").Error())
	})
}

func TestGetMyMeeting(t *testing.T) {
	repo := new(mocks.MeetingData)

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

	t.Run("Get MyMeeting Success", func(t *testing.T) {
		repo.On("GetMyMeetingID", mock.Anything).Return(insertData, nil).Once()

		useCase := New(repo, validator.New())

		res, _ := useCase.GetSeekMeeting(meetingData.ID)
		// assert.Nil(t, err)
		assert.GreaterOrEqual(t, 1, len(res))
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get Meeting", func(t *testing.T) {
		repo.On("GetMyMeetingID", mock.Anything).Return([]domain.MeetingOwner{}, errors.New("error get data")).Once()

		useCase := New(repo, validator.New())

		res, err := useCase.GetSeekMeeting(-1)
		// assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, errors.New("error get data").Error())
	})
}

func TestGetEmail(t *testing.T) {
	repo := new(mocks.MeetingData)

	returnDataOwner := domain.Ownerdata{Email: "lukman@gmail.com", Address: "jln.cijantung", City: "Jakarta"}
	returnDataSeeker := domain.Seekerdata{Email: "jacob@gmail.com"}

	t.Run("success get email", func(t *testing.T) {
		repo.On("GetEmailData", mock.Anything, mock.Anything).Return(returnDataOwner, returnDataSeeker, 200).Once()
		useCase := New(repo, validator.New())
		resOwner, resSeeker := useCase.GetEmail(2, 1)

		assert.Equal(t, returnDataOwner, resOwner)
		assert.Equal(t, returnDataSeeker, resSeeker)
		repo.AssertExpectations(t)
	})

	t.Run("Data not found", func(t *testing.T) {
		repo.On("GetEmailData", mock.Anything, mock.Anything).Return(domain.Ownerdata{}, domain.Seekerdata{}, 404).Once()
		useCase := New(repo, validator.New())
		resOwner, resSeeker := useCase.GetEmail(2, 1)

		assert.Equal(t, domain.Ownerdata{}, resOwner)
		assert.Equal(t, domain.Seekerdata{}, resSeeker)
		repo.AssertExpectations(t)
	})

	t.Run("Internal server error", func(t *testing.T) {
		repo.On("GetEmailData", mock.Anything, mock.Anything).Return(domain.Ownerdata{}, domain.Seekerdata{}, 500).Once()
		useCase := New(repo, validator.New())
		resOwner, resSeeker := useCase.GetEmail(2, 1)

		assert.Equal(t, domain.Ownerdata{}, resOwner)
		assert.Equal(t, domain.Seekerdata{}, resSeeker)
		repo.AssertExpectations(t)
	})
}

func TestGetPetMeeting(t *testing.T) {
	repo := new(mocks.MeetingData)

	insertData := domain.Meeting{
		ID:   1,
		Time: "09:00:00",
		Date: "21/08/2022",
	}
	outputData := domain.Meeting{
		ID:   1,
		Time: "09:00:00",
		Date: "21/08/2022",
	}
	t.Run("Get Pet Meeting Success", func(t *testing.T) {
		repo.On("GetMyMeetingPets", mock.Anything).Return(insertData, nil).Once()

		useCase := New(repo, validator.New())

		res, err := useCase.GetPetMeeting(insertData.ID)
		assert.Nil(t, err)
		assert.Equal(t, outputData, res)
		repo.AssertExpectations(t)
	})
	t.Run("Get Pet Meeting Failed", func(t *testing.T) {
		repo.On("GetMyMeetingPets", mock.Anything).Return(domain.Meeting{}, gorm.ErrRecordNotFound).Once()

		useCase := New(repo, validator.New())

		res, err := useCase.GetPetMeeting(insertData.ID)
		assert.NotNil(t, err)
		assert.Equal(t, domain.Meeting{}, res)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Server Error", func(t *testing.T) {
		repo.On("GetMyMeetingPets", mock.Anything).Return(domain.Meeting{}, errors.New("server error")).Once()

		useCase := New(repo, validator.New())

		_, err := useCase.GetPetMeeting(insertData.ID)
		assert.NotNil(t, err)
		assert.EqualError(t, err, errors.New("server error").Error())
	})
}
