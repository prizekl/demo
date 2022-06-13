package model

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
	_ "github.com/go-sql-driver/mysql"
)

var store *countryStore

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

	store = NewCountryStore(db)

	// defer func() {
	// 	for _, t := range []string{"country", "currency", "country_currency"} {
	// 		_, _ = db.Exec(fmt.Sprintf("DELETE FROM %s", t))
	// 	}
	// 	db.Close()
	// 	}()

	return m.Run(), nil
}
