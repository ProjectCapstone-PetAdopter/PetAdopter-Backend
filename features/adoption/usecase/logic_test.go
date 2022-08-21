package usecase

import (
	"petadopter/domain"
	"petadopter/domain/mocks"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteAdoption(t *testing.T) {
	repo := new(mocks.AdoptionData)

	t.Run("Success delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(true).Once()
		usecase := New(repo, validator.New())
		delete, err := usecase.DelAdoption(1)

		assert.Nil(t, err)
		assert.Equal(t, true, delete)
		repo.AssertExpectations(t)
	})

	t.Run("Failed delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(false).Once()
		usecase := New(repo, validator.New())
		delete, err := usecase.DelAdoption(1)

		assert.NotNil(t, err)
		assert.Equal(t, false, delete)
		repo.AssertExpectations(t)
	})
}

func TestGetAllAdoption(t *testing.T) {
	repo := new(mocks.AdoptionData)

	returnData := []domain.AdoptionPet{{
		ID:       1,
		Petname:  "Barky",
		Fullname: "Barky Alots",
		Status:   "Requested",
	}}

	returnDataSeeker := []domain.ApplierPet{{
		Fullname:     "John",
		UserID:       2,
		PhotoProfile: "John.jpg",
	}}

	t.Run("Success Get All", func(t *testing.T) {
		repo.On("GetAll", mock.Anything).Return(returnData, returnDataSeeker).Once()
		useCase := New(repo, validator.New())
		res, error := useCase.GetAllAP(1)

		assert.Equal(t, error, nil)
		assert.GreaterOrEqual(t, len(res), 1)
		repo.AssertExpectations(t)
	})
	t.Run("Data Not Found", func(t *testing.T) {
		useCase := New(repo, validator.New())
		result, error := useCase.GetAllAP(0)
		assert.NotNil(t, error)
		assert.Equal(t, []map[string]interface{}(nil), result)
		repo.AssertExpectations(t)
	})
}

func TestGetSpecificAdoption(t *testing.T) {
	repo := new(mocks.AdoptionData)

	returnData := []domain.AdoptionPet{{
		ID:           1,
		Petname:      "Barky",
		Petphoto:     "Barky.jpg",
		Fullname:     "Barky Alots",
		PhotoProfile: "Alots.jps",
		City:         "Gotham",
		Status:       "Requested",
	}}

	t.Run("Success Get Data", func(t *testing.T) {
		repo.On("GetAdoptionID", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res, error := useCase.GetSpecificAdoption(1)

		assert.Equal(t, error, nil)
		assert.GreaterOrEqual(t, len(res), 1)
		repo.AssertExpectations(t)
	})

	t.Run("Data Not Found", func(t *testing.T) {
		repo.On("GetAdoptionID", mock.Anything).Return(nil).Once()
		useCase := New(repo, validator.New())
		result, error := useCase.GetSpecificAdoption(0)

		assert.NotNil(t, error)
		assert.Equal(t, map[string]interface{}(nil), result)
		repo.AssertExpectations(t)
	})
}

func TestGetmyAdoption(t *testing.T) {
	repo := new(mocks.AdoptionData)

	returnData := []domain.AdoptionPet{{
		ID:           1,
		PetsID:       1,
		Petname:      "Barky",
		Petphoto:     "Barky.jpg",
		Fullname:     "Barky Alots",
		PhotoProfile: "Alots.jps",
		City:         "Gotham",
		Status:       "Requested",
	}}

	t.Run("Success Get Data", func(t *testing.T) {
		repo.On("GetAdoptionbyuser", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res, error := useCase.GetmyAdoption(1)

		assert.Equal(t, error, nil)
		assert.GreaterOrEqual(t, len(res), 1)
		repo.AssertExpectations(t)
	})
	t.Run("Data Not Found", func(t *testing.T) {
		useCase := New(repo, validator.New())
		result, error := useCase.GetmyAdoption(0)

		assert.NotNil(t, error)
		assert.Equal(t, []map[string]interface{}(nil), result)
		repo.AssertExpectations(t)
	})
}

func TestAddAdoption(t *testing.T) {
	repo := new(mocks.AdoptionData)

	mockData := domain.Adoption{PetsID: 1}

	returnData := mockData
	returnData.ID = 1
	returnData.UserID = 1
	returnData.Status = "Requested"

	t.Run("Success apply", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res, err := useCase.AddAdoption(1, mockData)

		assert.Nil(t, err)
		assert.Equal(t, returnData, res)
		repo.AssertExpectations(t)
	})

	t.Run("Invalid user", func(t *testing.T) {
		useCase := New(repo, validator.New())
		res, err := useCase.AddAdoption(0, mockData)

		assert.NotNil(t, err)
		assert.Equal(t, domain.Adoption{}, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error insert", func(t *testing.T) {
		returnData.ID = 0
		repo.On("Insert", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res, err := useCase.AddAdoption(1, mockData)

		assert.NotNil(t, err)
		assert.Equal(t, domain.Adoption{}, res)
		repo.AssertExpectations(t)
	})
}

func TestUpAdoption(t *testing.T) {
	repo := new(mocks.AdoptionData)

	mockData := domain.Adoption{Status: "Adopted"}

	returnData := mockData
	returnData.ID = 1
	returnData.PetsID = 1
	returnData.UserID = 1

	t.Run("Success Update", func(t *testing.T) {
		repo.On("Update", mock.Anything, mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res, err := useCase.UpAdoption(1, mockData)

		assert.Nil(t, err)
		assert.Equal(t, returnData, res)
		repo.AssertExpectations(t)
	})

	t.Run("Invalid user", func(t *testing.T) {
		useCase := New(repo, validator.New())
		res, err := useCase.UpAdoption(0, mockData)

		assert.NotNil(t, err)
		assert.Equal(t, domain.Adoption{}, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error insert", func(t *testing.T) {
		returnData.ID = 0
		repo.On("Update", mock.Anything, mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res, err := useCase.UpAdoption(1, mockData)

		assert.NotNil(t, err)
		assert.Equal(t, domain.Adoption{}, res)
		repo.AssertExpectations(t)
	})
}
