package usecase

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"os"
	"petadopter/domain"
	"petadopter/domain/mocks"
	"petadopter/utils/google"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeletePets(t *testing.T) {
	repo := new(mocks.PetsData)

	t.Run("Succes delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(true).Once()
		usecase := New(repo, validator.New(), google.InitStorage("pet-adopter-358806-9e20643cb88d.json", "be10-petdopter", "pet-adopter-358806"))
		delete, err := usecase.DelPets(1)

		assert.Nil(t, err)
		assert.Equal(t, true, delete)
		repo.AssertExpectations(t)
	})

	t.Run("Failed delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(false).Once()
		usecase := New(repo, validator.New(), google.InitStorage("pet-adopter-358806-9e20643cb88d.json", "be10-petdopter", "pet-adopter-358806"))
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
		useCase := New(repo, validator.New(), google.InitStorage("pet-adopter-358806-9e20643cb88d.json", "be10-petdopter", "pet-adopter-358806"))
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
		useCase := New(repo, validator.New(), google.InitStorage("pet-adopter-358806-9e20643cb88d.json", "be10-petdopter", "pet-adopter-358806"))
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
		repo.On("GetPetUser", mock.Anything, mock.Anything).Return(returnDataPetUser).Once()
		useCase := New(repo, validator.New(), google.InitStorage("pet-adopter-358806-9e20643cb88d.json", "be10-petdopter", "pet-adopter-358806"))
		res, error := useCase.GetmyPets(1)

		assert.Nil(t, error)
		assert.GreaterOrEqual(t, len(res), 1)
		assert.Equal(t, len(res[0]), 10) // 10 adalah panjang map dari array ke 0
		repo.AssertExpectations(t)
	})

	t.Run("Data Not Found", func(t *testing.T) {
		repo.On("GetPetsbyuser", mock.Anything).Return([]domain.Pets{}).Once()
		repo.On("GetPetUser", mock.Anything, mock.Anything).Return(domain.PetUser{}).Once()
		useCase := New(repo, validator.New(), google.InitStorage("pet-adopter-358806-9e20643cb88d.json", "be10-petdopter", "pet-adopter-358806"))
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

	fileContents, _ := os.ReadFile("./aki.jpg")
	body := new(bytes.Buffer)

	_, _ = body.Read(fileContents)
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("petphoto", "aki.jpg")

	_, _ = part.Write(fileContents)
	_ = writer.WriteField("petphoto", string(fileContents))

	err := writer.Close()
	if err != nil {
		fmt.Println(err)
	}

	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/pets", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	form, err := c.FormFile("petphoto")
	if err != nil {
		fmt.Println(err)
	}

	// header := make(map[string][]string)

	// header["Content-Disposition"] = []string{` form-data; name="file"; filename="karthus.jpg"`}

	// header["Content-Type"] = []string{"image/jpeg"}
	// FileHeader := &multipart.FileHeader{

	// 	Filename: "karthus.jpg",

	// 	Header: header,

	// 	Size: 1289231,
	// }

	FileHeader := &multipart.FileHeader{
		Filename: "b.JPG",
		Header:   textproto.MIMEHeader{"Content-Disposition": {"form-data", "name=photoprofile", "filename=b.JPG"}, "Content-Type": {"image/jpeg"}},
		Size:     0,
	}

	t.Run("Success add", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New(), google.InitStorage("pet-adopter-358806-9e20643cb88d.json", "be10-petdopter", "pet-adopter-358806"))
		res, err := useCase.AddPets(mockData, 1, form)

		assert.Nil(t, err)
		assert.Equal(t, returnData.ID, res.ID)
		assert.Equal(t, mockData.Userid, res.Userid)
	})
	t.Run("Error insert", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Pets{}).Once()
		useCase := New(repo, validator.New(), google.InitStorage("pet-adopter-358806-9e20643cb88d.json", "be10-petdopter", "pet-adopter-358806"))
		res, err := useCase.AddPets(mockData, 0, form)

		assert.NotNil(t, err)
		assert.Equal(t, 0, res.ID)
		assert.Equal(t, 0, res.Userid)
	})

	t.Run("Cant Upload File", func(t *testing.T) {
		useCase := New(repo, validator.New(), google.InitStorage("pet-adopter-358806-9e20643cb88d.json", "", ""))
		res, err := useCase.AddPets(mockData, 0, form)

		assert.NotNil(t, err)
		assert.Equal(t, 0, res.ID)
		assert.Equal(t, 0, res.Userid)
	})
	t.Run("Cant Open File", func(t *testing.T) {
		form.Size = 1
		useCase := New(repo, validator.New(), google.InitStorage("pet-adopter-358806-9e20643cb88d.json", "be10-petdopter", "pet-adopter-358806"))
		res, err := useCase.AddPets(mockData, 0, FileHeader)

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
		repo.On("GetPetUser", mock.Anything, mock.Anything).Return(returnDataPetUser).Once()
		useCase := New(repo, validator.New(), google.InitStorage("pet-adopter-358806-9e20643cb88d.json", "be10-petdopter", "pet-adopter-358806"))
		res, error := useCase.GetSpecificPets(1)

		assert.Nil(t, error)
		assert.NotNil(t, res)
		assert.Equal(t, 11, len(res))
		repo.AssertExpectations(t)
	})

	t.Run("Data Pets Not Found", func(t *testing.T) {
		repo.On("GetPetsID", mock.Anything).Return([]domain.Pets(nil)).Once()
		useCase := New(repo, validator.New(), google.InitStorage("pet-adopter-358806-9e20643cb88d.json", "be10-petdopter", "pet-adopter-358806"))
		res, error := useCase.GetSpecificPets(0)

		assert.NotNil(t, error)
		assert.Nil(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("Data PetUser Not Found", func(t *testing.T) {
		repo.On("GetPetsID", mock.Anything).Return(returnData).Once()
		repo.On("GetPetUser", mock.Anything, mock.Anything).Return(domain.PetUser{}).Once()
		useCase := New(repo, validator.New(), google.InitStorage("pet-adopter-358806-9e20643cb88d.json", "be10-petdopter", "pet-adopter-358806"))
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

	fileContents, _ := os.ReadFile("./aki.jpg")
	body := new(bytes.Buffer)

	_, _ = body.Read(fileContents)
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("petphoto", "aki.jpg")

	_, _ = part.Write(fileContents)
	_ = writer.WriteField("petphoto", string(fileContents))

	err := writer.Close()
	if err != nil {
		fmt.Println(err)
	}

	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/pets", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	form, err := c.FormFile("petphoto")
	if err != nil {
		fmt.Println(err)
	}

	// header := make(map[string][]string)

	// header["Content-Disposition"] = []string{` form-data; name="file"; filename="karthus.jpg"`}

	// header["Content-Type"] = []string{"image/jpeg"}
	// FileHeader := &multipart.FileHeader{

	// 	Filename: "karthus.jpg",

	// 	Header: header,

	// 	Size: 1289231,
	// }

	FileHeader := &multipart.FileHeader{
		Filename: "b.JPG",
		Header:   textproto.MIMEHeader{"Content-Disposition": {"form-data", "name=photoprofile", "filename=b.JPG"}, "Content-Type": {"image/jpeg"}},
		Size:     0,
	}

	t.Run("Success Update", func(t *testing.T) {
		repo.On("Update", mock.Anything, mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New(), google.InitStorage("pet-adopter-358806-9e20643cb88d.json", "be10-petdopter", "pet-adopter-358806"))
		res, err := useCase.UpPets(1, mockData, 1, form)

		assert.Nil(t, err)
		assert.Equal(t, returnData.ID, res.ID)
		assert.Equal(t, mockData.Userid, res.Userid)
	})
	t.Run("Error Update", func(t *testing.T) {
		repo.On("Update", mock.Anything, mock.Anything).Return(domain.Pets{}).Once()
		useCase := New(repo, validator.New(), google.InitStorage("pet-adopter-358806-9e20643cb88d.json", "be10-petdopter", "pet-adopter-358806"))
		res, err := useCase.UpPets(0, mockData, 0, form)

		assert.NotNil(t, err)
		assert.Equal(t, 0, res.ID)
		assert.Equal(t, 0, res.Userid)
	})

	t.Run("Cant Upload File", func(t *testing.T) {
		useCase := New(repo, validator.New(), google.InitStorage("pet-adopter-358806-9e20643cb88d.json", "", ""))
		res, err := useCase.UpPets(1, mockData, 0, form)

		assert.NotNil(t, err)
		assert.Equal(t, 0, res.ID)
		assert.Equal(t, 0, res.Userid)
	})
	t.Run("Cant Open File", func(t *testing.T) {
		form.Size = 1
		useCase := New(repo, validator.New(), google.InitStorage("pet-adopter-358806-9e20643cb88d.json", "be10-petdopter", "pet-adopter-358806"))
		res, err := useCase.UpPets(1, mockData, 0, FileHeader)

		assert.NotNil(t, err)
		assert.Equal(t, 0, res.ID)
		assert.Equal(t, 0, res.Userid)
	})
}
