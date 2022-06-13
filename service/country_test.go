package service

import (
	"fmt"
	"testing"
)

func TestGetAllCountries(t *testing.T) {
	res, err := service.GetAllCountries()
	if err != nil {
		t.Errorf("GetAllCountries service failed with error: %s", err)
	}

	for _, ctr := range res {
		fmt.Println(ctr)
	}
}

func TestGetCountries(t *testing.T) {

	q := map[string]string {
		"region": "Asia",
		"min":    "0",
		"max":    "60800000",
	}

	res, err := service.GetCountries(q)
	if err != nil {
		t.Errorf("GetCountries service failed with error: %s", err)
	}

	for _ , ctr := range res {
		fmt.Println(ctr)
	}
}
