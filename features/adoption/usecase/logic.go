package usecase

import (
	"errors"
	"petadopter/domain"

	validator "github.com/go-playground/validator/v10"
)

type adoptionUseCase struct {
	adoptionData domain.AdoptionData
	validate     *validator.Validate
}

func New(ud domain.AdoptionData, v *validator.Validate) domain.AdoptionUseCase {
	return &adoptionUseCase{
		adoptionData: ud,
		validate:     v,
	}
}

func (au *adoptionUseCase) AddAdoption(IDUser int, newAdoption domain.Adoption) (domain.Adoption, error) {
	if IDUser < 1 {
		return domain.Adoption{}, errors.New("invalid user")
	}

	newAdoption.UserID = IDUser

	res := au.adoptionData.Insert(newAdoption)
	if res.ID == 0 {
		return domain.Adoption{}, errors.New("error insert")
	}

	return res, nil
}

func (au *adoptionUseCase) GetSpecificAdoption(adoptionID int) (map[string]interface{}, error) {
	var res = map[string]interface{}{}
	data := au.adoptionData.GetAdoptionID(adoptionID)

	if adoptionID < 1 {
		return nil, errors.New("error get Data")
	}

	res["adoptionid"] = data[0].ID
	res["petname"] = data[0].Petname
	res["petphoto"] = data[0].Petphoto
	res["ownername"] = data[0].Fullname
	res["ownerphoto"] = data[0].PhotoProfile
	res["ownercity"] = data[0].City
	res["status"] = data[0].Status

	return res, nil
}

func (au *adoptionUseCase) GetAllAP(userid int) ([]map[string]interface{}, error) {
	var arrmap = []map[string]interface{}{}
	data, dataSeeker := au.adoptionData.GetAll(userid)

	if data == nil {
		return nil, errors.New("no data found")
	}

	for i := 0; i < len(data); i++ {
		var res = map[string]interface{}{}
		res["adoptionid"] = data[i].ID
		res["petname"] = data[i].Petname
		res["ownername"] = data[i].Fullname
		res["seekerid"] = dataSeeker[i].UserID
		res["seekername"] = dataSeeker[i].Fullname
		res["seekerphoto"] = dataSeeker[i].PhotoProfile
		res["status"] = data[i].Status

		arrmap = append(arrmap, res)
	}

	return arrmap, nil
}

func (au *adoptionUseCase) UpAdoption(IDAdoption int, updateData domain.Adoption, userID int) (domain.Adoption, error) {

	if IDAdoption < 1 {
		return domain.Adoption{}, errors.New("invalid data")
	}
	updateData.UserID = userID
	result := au.adoptionData.Update(IDAdoption, updateData)

	if result.ID == 0 {
		return domain.Adoption{}, errors.New("error update")
	}
	return result, nil
}

func (au *adoptionUseCase) DelAdoption(IDAdoption int) (bool, error) {
	res := au.adoptionData.Delete(IDAdoption)

	if !res {
		return false, errors.New("failed delete")
	}

	return true, nil
}

func (au *adoptionUseCase) GetmyAdoption(userID int) ([]map[string]interface{}, error) {
	var arrmap = []map[string]interface{}{}
	data := au.adoptionData.GetAdoptionbyuser(userID)

	if userID < 1 {
		return nil, errors.New("error get data")
	}

	for i := 0; i < len(data); i++ {
		var res = map[string]interface{}{}
		res["adoptionid"] = data[i].ID
		res["petid"] = data[i].PetsID
		res["petname"] = data[i].Petname
		res["petphoto"] = data[i].Petphoto
		res["ownername"] = data[i].Fullname
		res["ownerphoto"] = data[i].PhotoProfile
		res["ownercity"] = data[i].City
		res["status"] = data[i].Status

		arrmap = append(arrmap, res)
	}

	return arrmap, nil
}
