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
	if IDUser == -1 {
		return domain.Adoption{}, errors.New("invalid user")
	}

	newAdoption.UserID = IDUser

	res := au.adoptionData.Insert(newAdoption)
	if res.ID == 0 {
		return domain.Adoption{}, errors.New("error insert")
	}

	return res, nil
}

func (au *adoptionUseCase) GetSpecificAdoption(adoptionID int) ([]domain.AdoptionPet, error) {
	res := au.adoptionData.GetAdoptionID(adoptionID)
	if adoptionID == -1 {
		return nil, errors.New("error get Data")
	}

	return res, nil
}

func (au *adoptionUseCase) GetAllAP(userid int) ([]domain.AdoptionPet, error) {
	res := au.adoptionData.GetAll(userid)

	if len(res) == 0 {
		return nil, errors.New("no data found")
	}

	return res, nil
}

func (au *adoptionUseCase) UpAdoption(IDAdoption int, updateData domain.Adoption) (domain.Adoption, error) {

	if IDAdoption == -1 {
		return domain.Adoption{}, errors.New("invalid data")
	}
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

func (au *adoptionUseCase) GetmyAdoption(userID int) ([]domain.AdoptionPet, error) {
	res := au.adoptionData.GetAdoptionbyuser(userID)
	if userID == -1 {
		return nil, errors.New("error get data")
	}

	return res, nil
}
