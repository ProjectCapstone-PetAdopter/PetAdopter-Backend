package usecase

import (
	"errors"
	"petadopter/domain"

	validator "github.com/go-playground/validator/v10"
)

type petsUseCase struct {
	petsData domain.PetsData
	validate *validator.Validate
}

func New(ud domain.PetsData, v *validator.Validate) domain.PetsUseCase {
	return &petsUseCase{
		petsData: ud,
		validate: v,
	}
}

func (pd *petsUseCase) AddPets(newPets domain.Pets) (domain.Pets, error) {
	res := pd.petsData.Insert(newPets)
	if res.ID == 0 {
		return domain.Pets{}, errors.New("error insert data")
	}

	return res, nil
}

func (pd *petsUseCase) GetSpecificPets(petsID int) ([]domain.Pets, domain.PetUser, error) {
	petuser := pd.petsData.GetPetUser()
	res := pd.petsData.GetPetsID(petsID)

	if len(res) == 0 {
		return nil, domain.PetUser{}, errors.New("error get Pet")
	}

	return res, petuser, nil
}

func (pd *petsUseCase) GetAllP() ([]domain.Pets, error) {
	res := pd.petsData.GetAll()
	if len(res) == 0 {
		return nil, errors.New("no data found")
	}

	return res, nil
}

func (pd *petsUseCase) UpPets(IDPets int, updateData domain.Pets) (domain.Pets, error) {
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

func (pd *petsUseCase) GetmyPets(userID int) ([]domain.Pets, error) {
	res := pd.petsData.GetPetsbyuser(userID)
	if userID == -1 {
		return nil, errors.New("error get data")
	}

	return res, nil
}
