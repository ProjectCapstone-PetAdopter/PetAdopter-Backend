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
		Status:   "waiting",
	}}

	// returnempty := []domain.AdoptionPet{{
	// 	ID:       0,
	// 	Petname:  "",
	// 	Fullname: "",
	// 	Status:   "",
	// }}

	t.Run("Success Get All", func(t *testing.T) {
		repo.On("GetAll", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res, error := useCase.GetAllAP(1)

		assert.Equal(t, error, nil)
		assert.GreaterOrEqual(t, len(res), 1)
		repo.AssertExpectations(t)
	})
	// t.Run("Data Not Found", func(t *testing.T) {
	// 	repo.On("GetAll", mock.Anything).Return(returnData).Once()
	// 	useCase := New(repo, validator.New())
	// 	result, error := useCase.GetAllAP(0)

	// 	assert.Equal(t, error, nil)
	// 	assert.Equal(t, "Somrthing wrong here, please fix me asap", result)
	// 	repo.AssertExpectations(t)
	// })
}
