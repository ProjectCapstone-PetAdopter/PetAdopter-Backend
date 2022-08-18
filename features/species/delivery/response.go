package delivery

import "petadopter/domain"

type SpeciesResponse struct {
	ID      int    `json:"id"`
	Species string `json:"species"`
}

func FromModel(data domain.Species) SpeciesResponse {
	var res SpeciesResponse
	res.ID = data.ID
	res.Species = data.Species
	return res
}

func FromModelToList(data []domain.Species) []SpeciesResponse {
	result := []SpeciesResponse{}
	for val := range data {
		result = append(result, FromModel(data[val]))
	}
	return result
}
