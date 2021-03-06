package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Ingredient struct {
	Ingredient string `json:"ingredient"`
	Mesure     string `json:"mesures"`
}

type IngredientCart struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Url   string  `json:"url"`
}

type Cart struct {
	Wallmart []IngredientCart `json:"wallmart"`
}

// Recipe structures is used to load recipes data from JSON request
type Recipe struct {
	Ingredients  []Ingredient `json:"ingredients"`
	Img          string       `json:"img"`
	Name         string       `json:"name"`
	Video        string       `json:"video"`
	Instructions string       `json:"instructions"`
	Cart         Cart         `json:"cart"`
	Error        bool         `json:"error"`
}

// GetRandomRecipe request the Reciprice Api and return a Recipe object with a random recipe.
func GetRandomRecipe() Recipe {
	httpClient := http.Client{
		Timeout: time.Second * 60,
	}
	req, err := http.NewRequest(http.MethodGet, config.APIURL+"api/v1/random", nil)
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

func GetRecipe(name string) Recipe {
	httpClient := http.Client{
		Timeout: time.Second * 60,
	}
	req, err := http.NewRequest(http.MethodGet, config.APIURL+"api/v1/search/"+name, nil)
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
