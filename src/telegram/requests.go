package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Recipe struct {
	Img   string `json:"img"`
	Name  string `json:"name"`
	Video string `json:"video"`
}

func GetRandomRecipe() Recipe {
	log.Println("oui")
	httpClient := http.Client{
		Timeout: time.Second * 2,
	}
	req, err := http.NewRequest(http.MethodGet, config.ApiURL+"api/v1/random", nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	recipe := Recipe{}
	err = json.Unmarshal(body, &recipe)
	if err != nil {
		log.Fatal(err)
	}

	return recipe
}