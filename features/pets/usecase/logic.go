package usecase

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"petadopter/config"
	"petadopter/domain"
	"petadopter/utils/google"

	validator "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type petsUseCase struct {
	petsData domain.PetsData
	validate *validator.Validate
	client   *google.ClientUploader
}

func New(ud domain.PetsData, v *validator.Validate, cl *google.ClientUploader) domain.PetsUseCase {
	return &petsUseCase{
		petsData: ud,
		validate: v,
		client:   cl,
	}
}

func (pd *petsUseCase) AddPets(newPets domain.Pets, userId int, form *multipart.FileHeader) (domain.Pets, error) {
	newPets.Userid = userId

	if form != nil {
		file, err := form.Open()
		if err != nil {
			log.Println(err, "cant open file")
			return domain.Pets{}, errors.New("cant open file")
		}

		defer file.Close()
		id := uuid.New()
		filename := fmt.Sprintf("%dPost-%s.jpg", newPets.Userid, id.String())
		config.UPLOADPATH = "profile/"

		link, err := pd.client.UploadFile(file, config.UPLOADPATH, filename)
		if err != nil {
			log.Println(err, "cant upload file")
			return domain.Pets{}, errors.New("cant upload file")
		}
		newPets.Petphoto = link
	}

	res := pd.petsData.Insert(newPets)

	if res.ID == 0 {
		return domain.Pets{}, errors.New("error insert data")
	}

	return res, nil
}

func (pd *petsUseCase) GetSpecificPets(petsID int) (map[string]interface{}, error) {
	//map untuk output API agar sama dengan swagger
	var res = map[string]interface{}{}
	var petUser = domain.PetUser{}

	data := pd.petsData.GetPetsID(petsID)
	if data == nil {
		return nil, errors.New("error get Pet")
	}

	dataPetUser := pd.petsData.GetPetUser(data[0].Userid, petsID)
	if dataPetUser == petUser { //jika isinya struct kosong
		return nil, errors.New("error get Pet user")
	}

	res["petname"] = data[0].Petname
	res["petphoto"] = data[0].Petphoto
	res["species"] = dataPetUser.Species
	res["gender"] = data[0].Gender
	res["age"] = data[0].Age
	res["color"] = data[0].Color
	res["description"] = data[0].Description
	res["ownerid"] = data[0].Userid
	res["ownername"] = dataPetUser.Fullname
	res["city"] = dataPetUser.City

	//res akan ditampilkan
	return res, nil
}

func (pd *petsUseCase) GetAllP() ([]map[string]interface{}, error) {
	//membuat array of maps agar sama dengan swagger
	var arrmap = []map[string]interface{}{}

	data := pd.petsData.GetAll()
	dataPetUser := pd.petsData.GetAllPetUser()
	if len(data) == 0 {
		return nil, errors.New("no data found")
	}

	//memasukan map kedalam array sesuai dengan panjang array data yang didapat
	for i := 0; i < len(data); i++ {
		var res = map[string]interface{}{}
		res["id"] = data[i].ID
		res["petname"] = data[i].Petname
		res["petphoto"] = data[i].Petphoto
		res["species"] = dataPetUser[i].Species
		res["gender"] = data[i].Gender
		res["age"] = data[i].Age
		res["color"] = data[i].Color
		res["description"] = data[i].Description
		res["ownername"] = dataPetUser[i].Fullname
		res["city"] = dataPetUser[i].City
		arrmap = append(arrmap, res)
	}

	return arrmap, nil
}

func (pd *petsUseCase) UpPets(IDPets int, updateData domain.Pets, userID int, form *multipart.FileHeader) (domain.Pets, error) {
	updateData.Userid = userID

	if form != nil {
		file, err := form.Open()
		if err != nil {
			log.Println(err, "cant open file")
			return domain.Pets{}, errors.New("cant open file")
		}

		defer file.Close()
		id := uuid.New()
		filename := fmt.Sprintf("%dPost-%s.jpg", updateData.Userid, id.String())
		config.UPLOADPATH = "profile/"

		link, err := pd.client.UploadFile(file, config.UPLOADPATH, filename)
		if err != nil {
			log.Println(err, "cant upload file")
			return domain.Pets{}, errors.New("cant upload file")
		}
		updateData.Petphoto = link
	}

	result := pd.petsData.Update(IDPets, updateData)
	if result.ID == 0 {
		return domain.Pets{}, errors.New("error update data")
	}

	return result, nil
}

func (pd *petsUseCase) DelPets(IDPets int) (bool, error) {
	res := pd.petsData.Delete(IDPets)
	if !res {
		return false, errors.New("failed delete")
	}

	return true, nil
}

func (pd *petsUseCase) GetmyPets(userID int) ([]map[string]interface{}, error) {
	//membuat array of maps agar sama dengan swagger
	var arrmap = []map[string]interface{}{}
	data := pd.petsData.GetPetsbyuser(userID)
	dataPetUser := pd.petsData.GetPetUser(userID, 0)
	if userID < 1 {
		return nil, errors.New("error get data")
	}
	//memasukan map kedalam array sesuai dengan panjang array data yang didapat
	for i := 0; i < len(data); i++ {
		var res = map[string]interface{}{}
		res["petid"] = data[i].ID
		res["petname"] = data[i].Petname
		res["petphoto"] = data[i].Petphoto
		res["species"] = data[i].Speciesid
		res["gender"] = data[i].Gender
		res["age"] = data[i].Age
		res["color"] = data[i].Color
		res["description"] = data[i].Description
		res["ownername"] = dataPetUser.Fullname
		res["city"] = dataPetUser.City
		arrmap = append(arrmap, res)
	}

	return arrmap, nil
}
