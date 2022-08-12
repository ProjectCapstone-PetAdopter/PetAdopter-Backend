package usecase

import (
	"petadopter/domain"
	"petadopter/domain/mocks"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeletePets(t *testing.T) {
	repo := new(mocks.PetsData)

	t.Run("Succes delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(true).Once()
		usecase := New(repo, validator.New())
		delete, err := usecase.DelPets(1)

		assert.Nil(t, err)
		assert.Equal(t, true, delete)
		repo.AssertExpectations(t)
	})

	t.Run("Failed delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(false).Once()
		usecase := New(repo, validator.New())
		delete, err := usecase.DelPets(1)

		assert.NotNil(t, err)
		assert.Equal(t, false, delete)
		repo.AssertExpectations(t)
	})
}

func TestGetAllPets(t *testing.T) {
	repo := new(mocks.PetsData)
	returnData := []domain.Pets{{ID: 1, Speciesid: 1, Petname: "Barky", Age: 1, Color: "blue", Petphoto: "image.jpg"}}

	t.Run("Success Get All Pets", func(t *testing.T) {
		repo.On("GetAll", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res, error := useCase.GetAllP()

		assert.Equal(t, error, nil)
		assert.GreaterOrEqual(t, len(res), 1)
		assert.Greater(t, res[0].ID, 0)
		repo.AssertExpectations(t)
	})
	t.Run("Data Not Found", func(t *testing.T) {
		repo.On("GetAll", mock.Anything).Return([]domain.Pets{}).Once()
		useCase := New(repo, validator.New())
		res, _ := useCase.GetAllP()

		assert.Equal(t, len(res), 0)
		assert.Equal(t, []domain.Pets([]domain.Pets(nil)), res)
		assert.Equal(t, []domain.Pets(nil), res)
		repo.AssertExpectations(t)
	})
}

func TestGetmyPets(t *testing.T) {
	repo := new(mocks.PetsData)
	returnData := []domain.Pets{{ID: 1, Speciesid: 1, Petname: "Barky", Age: 1, Color: "blue", Petphoto: "image.jpg", Userid: 1}}

	t.Run("Success Get my Pets", func(t *testing.T) {
		repo.On("GetPetsbyuser", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res, error := useCase.GetmyPets(1)

		assert.Equal(t, error, nil)
		assert.GreaterOrEqual(t, len(res), 1)
		assert.Equal(t, res[0].ID, 1)
		repo.AssertExpectations(t)
	})

	t.Run("Data Not Found", func(t *testing.T) {
		repo.On("GetPetsbyuser", mock.Anything).Return([]domain.Pets{}).Once()
		useCase := New(repo, validator.New())
		res, error := useCase.GetmyPets(0)

		assert.Equal(t, len(res), 0)
		assert.Equal(t, error, nil)
		repo.AssertExpectations(t)
	})
}
