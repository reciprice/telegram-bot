package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Token  string `json:"token"`
	ApiURL string `json:"apiURL"`
}

func Load(filename string) Configuration {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Print("Error")
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}
