package usecase

import (
	"petadopter/domain"
	"petadopter/domain/mocks"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteSpecies(t *testing.T) {
	repo := new(mocks.SpeciesData)

	t.Run("Success delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(200, nil).Once()
		usecase := New(repo, validator.New())
		status, _ := usecase.DeleteSpecies(1)

		assert.Equal(t, 200, status)
		repo.AssertExpectations(t)
	})

	t.Run("Failed delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(404, nil).Once()
		usecase := New(repo, validator.New())
		delete, _ := usecase.DeleteSpecies(1)

		assert.NotNil(t, delete, 404)
		repo.AssertExpectations(t)
	})
}

func TestGetAllSpecies(t *testing.T) {
	repo := new(mocks.SpeciesData)

	returnData := []domain.Species{{
		ID:      1,
		Species: "Kucing",
	}}

	t.Run("Success Get All", func(t *testing.T) {
		repo.On("GetAll", mock.Anything).Return(returnData, nil).Once()
		useCase := New(repo, validator.New())
		res, error := useCase.GetAllSpecies()

		assert.Equal(t, error, nil)
		assert.GreaterOrEqual(t, len(res), 1)
		repo.AssertExpectations(t)
	})
	t.Run("Data Not Found", func(t *testing.T) {
		repo.On("GetAll", mock.Anything).Return(nil, nil).Once()
		useCase := New(repo, validator.New())
		result, _ := useCase.GetAllSpecies()
		assert.Equal(t, []domain.Species(nil), result)
		repo.AssertExpectations(t)
	})
}

func TestAddSpecies(t *testing.T) {
	//
}

func TestUpdateSpecies(t *testing.T) {
	//
}