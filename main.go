package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Data struct {
	Name string `json:"name"`
	Flag string `json:"flag"`
	Region string`json:"region"`
	Population int`json:"population"`
	Currencies []struct { Name string } `json:"currencies"`
}

func main() {
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
	fmt.Println(data)
}
