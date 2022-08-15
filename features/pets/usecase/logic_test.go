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
	returnDataPetUser := []domain.PetUser{{Species: "Kucing", Fullname: "Jacob Capung", City: "Gotham"}}

	t.Run("Success Get All Pets", func(t *testing.T) {
		repo.On("GetAll").Return(returnData).Once()
		repo.On("GetAllPetUser").Return(returnDataPetUser).Once()
		useCase := New(repo, validator.New())
		res, error := useCase.GetAllP()

		assert.Nil(t, error)
		assert.NotNil(t, res)
		assert.GreaterOrEqual(t, len(res), 1)
		assert.Greater(t, len(res[0]), 0)
		repo.AssertExpectations(t)
	})
	t.Run("Data Not Found", func(t *testing.T) {
		repo.On("GetAll").Return([]domain.Pets{}).Once()
		repo.On("GetAllPetUser").Return([]domain.PetUser{}).Once()
		useCase := New(repo, validator.New())
		res, err := useCase.GetAllP()

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, 0, len(res)) //tidak ada len(res[0]) karena jika data tidak ada maka res tidak akan terbuat dan array ke 0 tidak ada
		repo.AssertExpectations(t)
	})
}

func TestGetmyPets(t *testing.T) {
	repo := new(mocks.PetsData)
	returnData := []domain.Pets{{ID: 1, Speciesid: 1, Petname: "Barky", Age: 1, Color: "blue", Petphoto: "image.jpg", Userid: 1}}
	returnDataPetUser := domain.PetUser{Species: "Kucing", Fullname: "Jacob Capung", City: "Gotham"}

	t.Run("Success Get my Pets", func(t *testing.T) {
		repo.On("GetPetsbyuser", mock.Anything).Return(returnData).Once()
		repo.On("GetPetUser", mock.Anything).Return(returnDataPetUser).Once()
		useCase := New(repo, validator.New())
		res, error := useCase.GetmyPets(1)

		assert.Nil(t, error)
		assert.GreaterOrEqual(t, len(res), 1)
		assert.Equal(t, len(res[0]), 10) // 10 adalah panjang map dari array ke 0
		repo.AssertExpectations(t)
	})

	t.Run("Data Not Found", func(t *testing.T) {
		repo.On("GetPetsbyuser", mock.Anything).Return([]domain.Pets{}).Once()
		repo.On("GetPetUser", mock.Anything).Return(domain.PetUser{}).Once()
		useCase := New(repo, validator.New())
		res, error := useCase.GetmyPets(0)

		assert.NotNil(t, error)
		assert.Equal(t, len(res), 0)
		repo.AssertExpectations(t)
	})
}

func TestAddPets(t *testing.T) {
	repo := new(mocks.PetsData)
	mockData := domain.Pets{Speciesid: 1, Petname: "Barky", Age: 1, Color: "blue", Petphoto: "image.jpg", Userid: 1, Gender: 1, Description: "this is barky"}
	returnData := mockData
	returnData.ID = 1

	t.Run("Success add", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res, err := useCase.AddPets(mockData, 1)

		assert.Nil(t, err)
		assert.Equal(t, returnData.ID, res.ID)
		assert.Equal(t, mockData.Userid, res.Userid)
	})
	t.Run("Error insert", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Pets{}).Once()
		useCase := New(repo, validator.New())
		res, err := useCase.AddPets(mockData, 0)

		assert.NotNil(t, err)
		assert.Equal(t, 0, res.ID)
		assert.Equal(t, 0, res.Userid)
	})
}

func TestGetSpecificPets(t *testing.T) {
	repo := new(mocks.PetsData)
	returnData := []domain.Pets{{ID: 1, Speciesid: 1, Petname: "Barky", Age: 1, Color: "blue", Petphoto: "image.jpg", Userid: 1}}
	returnDataPetUser := domain.PetUser{Species: "Kucing", Fullname: "Jacob Capung", City: "Gotham"}

	t.Run("Success Get Pet by Id", func(t *testing.T) {
		repo.On("GetPetsID", mock.Anything).Return(returnData).Once()
		repo.On("GetPetUser", mock.Anything).Return(returnDataPetUser).Once()
		useCase := New(repo, validator.New())
		res, error := useCase.GetSpecificPets(1)

		assert.Nil(t, error)
		assert.NotNil(t, res)
		assert.Equal(t, 9, len(res))
		repo.AssertExpectations(t)
	})

	t.Run("Data Pets Not Found", func(t *testing.T) {
		repo.On("GetPetsID", mock.Anything).Return([]domain.Pets(nil)).Once()
		useCase := New(repo, validator.New())
		res, error := useCase.GetSpecificPets(0)

		assert.NotNil(t, error)
		assert.Nil(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("Data PetUser Not Found", func(t *testing.T) {
		repo.On("GetPetsID", mock.Anything).Return(returnData).Once()
		repo.On("GetPetUser", mock.Anything).Return(domain.PetUser{}).Once()
		useCase := New(repo, validator.New())
		res, error := useCase.GetSpecificPets(0)

		assert.NotNil(t, error)
		assert.Nil(t, res)
		repo.AssertExpectations(t)
	})
}

func TestUpPets(t *testing.T) {
	repo := new(mocks.PetsData)
	mockData := domain.Pets{Speciesid: 1, Petname: "Bark", Age: 1, Color: "blue", Petphoto: "image.jpg", Userid: 1, Gender: 1, Description: "this is barky"}
	returnData := mockData
	returnData.ID = 1

	t.Run("Success Update", func(t *testing.T) {
		repo.On("Update", mock.Anything, mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res, err := useCase.UpPets(1, mockData, 1)

		assert.Nil(t, err)
		assert.Equal(t, returnData.ID, res.ID)
		assert.Equal(t, mockData.Userid, res.Userid)
	})
	t.Run("Error Update", func(t *testing.T) {
		repo.On("Update", mock.Anything, mock.Anything).Return(domain.Pets{}).Once()
		useCase := New(repo, validator.New())
		res, err := useCase.UpPets(0, mockData, 0)

		assert.NotNil(t, err)
		assert.Equal(t, 0, res.ID)
		assert.Equal(t, 0, res.Userid)
	})
}
