package model

import (
	"database/sql"
	"fmt"
	"strings"
)

type country struct {
	Name       string `json:"name"`
	Flag       string `json:"flag"`
	Region     string `json:"region"`
	Population int    `json:"population"`
}

func NewCountry(name string, flag string, region string, population int) *country {
	return &country{
		Name:       name,
		Flag:       flag,
		Region:     region,
		Population: population,
	}
}

type CountryStore interface {
	InsertCountry(*country) (int64, error)
	GetAllCountries() ([]*country, error)
	GetCountries(map[string]any) ([]*country, error)
}

func NewCountryStore(db *sql.DB) *countryStore {
	return &countryStore{
		db: db,
	}
}

type countryStore struct {
	db *sql.DB
}

func (s *countryStore) InsertCountry(ctr *country) (int64, error) {

	var id int64

	query := fmt.Sprintf(
		`INSERT INTO country (
	name,
	flag,
	region,
	population
	) VALUES (
	?, ?, ?, ?
	)`)

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return id, err
	}

	res, err := stmt.Exec(ctr.Name, ctr.Flag, ctr.Region, ctr.Population)
	if err != nil {
		return id, err
	}

	id, err = res.LastInsertId()

	return id, err
}

func (s countryStore) GetAllCountries() ([]*country, error) {
	var ctrs []*country

	query := "SELECT name, flag, region, population FROM country"

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		ctr := &country{}
		if err := rows.Scan(
			&ctr.Name,
			&ctr.Flag,
			&ctr.Region,
			&ctr.Population,
		); err != nil {
			return nil, err
		}
		ctrs = append(ctrs, ctr)
	}

	return ctrs, nil
}

func (s countryStore) GetCountries(q map[string]any) ([]*country, error) {

	// q := map[string]interface{}{
	// 	"region": "Asia",
	// 	"min":    0,
	// 	"max":    10,
	// }
	//

	var ctrs []*country
	var val []any
	var where []string

	for k, v := range q {
		switch k {
		case "min":
			where = append(where, fmt.Sprintf(`%s >= ?`, "population"))
		case "max":
			where = append(where, fmt.Sprintf(`%s <= ?`, "population"))
		default:
			where = append(where, fmt.Sprintf(`%s = ?`, k))
		}
		val = append(val, v)
	}

	query := ("SELECT name, flag, region, population FROM Country WHERE " + strings.Join(where, " AND "))

	rows, err := s.db.Query(query, val...)
	if err != nil {
		return ctrs, err
	}
	defer rows.Close()

	for rows.Next() {
		ctr := &country{}
		if err := rows.Scan(
			&ctr.Name,
			&ctr.Flag,
			&ctr.Region,
			&ctr.Population,
		); err != nil {
			return nil, err
		}
		ctrs = append(ctrs, ctr)
	}

	return ctrs, nil
}
