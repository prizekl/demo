package service

import (
	"database/sql"
	"demo/model"
	"fmt"
	"os"
	"testing"
	_ "github.com/go-sql-driver/mysql"
)

var service *countryService

func TestMain(m *testing.M) {
	code, err := setup(m)
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}

func setup(m *testing.M) (code int, err error) {
	dns := "root@tcp(localhost:3306)/country"
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return -1, fmt.Errorf("could not connect to database: %w", err)
	}

	defer db.Close()

	// Assign countryStore to interface CountryStore
	var store model.CountryStore = model.NewCountryStore(db)
	service = NewCountryService(store)

	return m.Run(), nil
}
