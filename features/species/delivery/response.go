package delivery

import "petadopter/domain"

type SpeciesResponse struct {
	Species []string
}

func FromModel(data domain.Species) SpeciesResponse {
	var res SpeciesResponse
	return res
}

func FromModelToList(data []domain.Species) []string {
	result := []string{}
	for val := range data {
		result = append(result, data[val].Species)
	}
	return result
}
