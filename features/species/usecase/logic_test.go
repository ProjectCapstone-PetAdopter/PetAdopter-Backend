package usecase

import (
	"petadopter/domain"
	"petadopter/domain/mocks"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddSpecies(t *testing.T) {
	repo := new(mocks.SpeciesData)
	insertData := domain.Species{
		ID:      1,
		Species: "Kucing",
	}

	t.Run("Success Insert", func(t *testing.T) {
		repo.On("Insert Species", mock.Anything).Return(1, nil).Once()

		useCase := New(repo, validator.New())

		res, err := useCase.AddSpecies(insertData)
		assert.Nil(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})
}
