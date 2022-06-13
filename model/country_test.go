package model

import (
	"fmt"
	"testing"
)


// What failing case is possible?
func TestAddCountry(t *testing.T) {
	data := &country {
		"Thailand",
		"url",
		"Asia",
		60800000,
	}

	res, err := store.InsertCountry(data)
	if err != nil {
		t.Errorf("Failed to insert into db with error %s", err)
	}

	fmt.Printf("New country id: %d", res)
}

func TestGetAllCountries(t *testing.T) {
	res, err := store.GetAllCountries()
	if err != nil {
		t.Errorf("Failed to get from db with error %s", err)
	}

	for _, ctr := range res {
		fmt.Println(ctr.Name, ctr.Flag, ctr.Region, ctr.Population)
	}
}

func TestGetCountries(t *testing.T) {

	q := map[string]any {
		"region": "Asia",
		"min":    0,
		"max":    60800000,

	}

	res, err := store.GetCountries(q)
	if err != nil {
		t.Errorf("Failed to get from db with error %s", err)
	}

	for _, ctr := range res {
		fmt.Println(ctr.Name, ctr.Flag, ctr.Region, ctr.Population)
	}
}
