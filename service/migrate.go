package service

import (
	// "demo/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Data struct {
	name string `json:"name"`
	flag string `json:"flag"`
	region string`json:"region"`
	population int`json:"population"`
	currencies []struct { Name string } `json:"currencies"`
}

func migrate() {
	resp, err := http.Get("https://restcountries.com/v2/all")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var data []Data
	if err = json.Unmarshal(body, &data); err != nil {
		log.Fatalln(err)
	}
}

func insertData(data *Data) {
}

func insertCountry() {
}
