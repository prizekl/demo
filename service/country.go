package service

import (
	"demo/model"
	"strconv"
)

type CountryService interface {
	GetAllCountries() ([]*model.country, error)
	GetCountries(map[string]string) ([]*model.country, error)
}

type countryService struct {
	store model.CountryStore
}

func NewCountryService(store model.CountryStore) *countryService {
	return &countryService {
		store: store,
	}
}

func (s *countryService) GetAllCountries() ([]*model.country, error) {
	res, err := s.store.GetAllCountries()

	//TODO: Error handling?

	return res, err
}

func (s *countryService) GetCountries(q map[string]string) ([]*model.country, error) {

	newq := make(map[string]any)

	// Parsing types
	for k, v := range q {
		if (k == "min" || k == "max") {
			i, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			newq[k] = i
		} else {
			newq[k] = v
		}
	}

	return s.store.GetCountries(newq)
}
