package usecase

import (
	"errors"
	"fmt"
	"petadopter/domain"
	"petadopter/domain/mocks"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestAddSpecies(t *testing.T) {
	repo := new(mocks.SpeciesData)
	insertData := domain.Species{
		ID:      1,
		Species: "Kucing",
	}

	t.Run("Success Insert", func(t *testing.T) {
		repo.On("InsertSpecies", mock.Anything).Return(1, nil).Once()

		useCase := New(repo, validator.New())

		res, err := useCase.AddSpecies(insertData)
		assert.Nil(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Insert Empty Species", func(t *testing.T) {
		repo.On("InsertSpecies", mock.Anything, mock.Anything).Return(-1, errors.New("invalid species")).Once()
		useCase := New(repo, validator.New())
		dummy := insertData
		dummy.Species = ""

		res, err := useCase.AddSpecies(dummy)
		assert.NotNil(t, err)
		assert.Equal(t, -1, res)
		assert.EqualError(t, err, errors.New("invalid species").Error())
		// repo.AssertExpectations(t)
	})
}
func TestDeleteSpecies(t *testing.T) {
	repo := new(mocks.SpeciesData)
	insertDate := domain.Species{
		ID:      1,
		Species: "Kucing",
	}

	t.Run("Success delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(1, nil).Once()

		usecase := New(repo, validator.New())

		res, err := usecase.DeleteSpecies(insertDate.ID)
		assert.Nil(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Failed delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(0, fmt.Errorf("failed to delete species")).Once()

		usecase := New(repo, validator.New())

		_, err := usecase.DeleteSpecies(1)
		assert.NotNil(t, err)
		assert.Equal(t, err, fmt.Errorf("failed to delete species"))
		repo.AssertExpectations(t)
	})

	t.Run("Data Not Found", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(0, gorm.ErrRecordNotFound).Once()

		usecase := New(repo, validator.New())

		_, err := usecase.DeleteSpecies(1)
		assert.NotNil(t, err)
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

func TestUpdateSpecies(t *testing.T) {
	repo := new(mocks.SpeciesData)
	insertData := domain.Species{
		ID:      1,
		Species: "Kucing",
	}

	t.Run("Success Update Species", func(t *testing.T) {
		repo.On("Update", mock.Anything, mock.Anything).Return(1, nil).Once()

		useCase := New(repo, validator.New())

		res, err := useCase.UpdateSpecies(insertData.ID, insertData)
		assert.Nil(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Empty Update Species", func(t *testing.T) {
		dummy := insertData
		dummy.Species = ""
		useCase := New(repo, validator.New())

		res, err := useCase.UpdateSpecies(insertData.ID, dummy)

		assert.NotNil(t, err)
		assert.Equal(t, -1, res)
		assert.EqualError(t, err, errors.New("invalid species").Error())
		repo.AssertExpectations(t)
	})
}
