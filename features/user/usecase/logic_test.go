package usecase

import (
	"petadopter/domain"
	"petadopter/domain/mocks"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/oauth2"
)

func TestUpdateUser(t *testing.T) {
	repo := new(mocks.UserData)
	cost := 10

	mockData := domain.User{Username: "batman", Email: "brucewayne@gmail.com", Address: "jakarta", Password: "polar", PhotoProfile: "wayne.jpg"}

	returnData := mockData
	returnData.ID = 1

	t.Run("Success Update", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		repo.On("UpdateUserData", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res := useCase.UpdateUser(mockData, 1, cost)

		assert.Equal(t, 200, res)
		repo.AssertExpectations(t)
	})

	t.Run("Data Not Found", func(t *testing.T) {
		useCase := New(repo, validator.New())
		res := useCase.UpdateUser(mockData, 0, cost)

		assert.Equal(t, 404, res)
		repo.AssertExpectations(t)
	})

	t.Run("Generate Hash Error", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		useCase := New(repo, validator.New())
		res := useCase.UpdateUser(mockData, 1, 40)

		assert.Equal(t, 500, res)
		repo.AssertExpectations(t)
	})

	t.Run("Duplicate Data", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(true).Once()
		useCase := New(repo, validator.New())
		res := useCase.RegisterUser(mockData, cost, &oauth2.Token{})

		assert.Equal(t, 409, res)
		repo.AssertExpectations(t)
	})
}

func TestLoginUser(t *testing.T) {
	repo := new(mocks.UserData)
	mockData := domain.User{Username: "batman", Password: "polar"}
	returnData := domain.User{ID: 1}
	notfound := mockData
	notfound.ID = 0
	token := "sjage2w62vsdgaqsgh"

	t.Run("Succes Login", func(t *testing.T) {
		repo.On("GetPasswordData", mock.Anything).Return("$2a$10$SrMvwwY/QnQ4nZunBvGOuOm2U1w8wcAENOoAMI7l8xH7C1Vmt5mru")
		repo.On("Login", mock.Anything).Return(returnData, token).Once()
		userUseCase := New(repo, validator.New())
		res, err := userUseCase.Login(mockData)

		assert.Nil(t, err)
		assert.Greater(t, res.ID, 0)
		repo.AssertExpectations(t)
	})

	t.Run("Data not found", func(t *testing.T) {
		repo.On("GetPasswordData", mock.Anything).Return("$2a$10$SrMvwwY/QnQ4nZunBvGOuOm2U1w8wcAENOoAMI7l8xH7C1Vmt5mru")
		repo.On("Login", mock.Anything).Return(notfound, token).Once()
		userUseCase := New(repo, validator.New())
		res, err := userUseCase.Login(mockData)

		assert.NotNil(t, err)
		assert.Equal(t, res.ID, 0)
		repo.AssertExpectations(t)
	})

	t.Run("Wrong input", func(t *testing.T) {
		mockData.Password = ""
		repo.On("GetPasswordData", mock.Anything).Return("$2a$10$SrMvwwY/QnQ4nZunBvGOuOm2U1w8w")
		userUseCase := New(repo, validator.New())
		res, err := userUseCase.Login(mockData)

		assert.NotNil(t, err)
		assert.Equal(t, 0, res.ID)
		repo.AssertExpectations(t)
	})
}

func TestDeleteUser(t *testing.T) {
	repo := new(mocks.UserData)

	t.Run("Succes delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(true).Once()
		usecase := New(repo, validator.New())
		delete, err := usecase.Delete(1)

		assert.Nil(t, err)
		assert.Equal(t, true, delete)
		repo.AssertExpectations(t)
	})

	t.Run("Failed delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(false).Once()
		usecase := New(repo, validator.New())
		delete, err := usecase.Delete(1)

		assert.NotNil(t, err)
		assert.Equal(t, false, delete)
		repo.AssertExpectations(t)
	})
}
