package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Configuration is the model for loading JSON configuration file
type Configuration struct {
	Token  string `json:"token"`
	APIURL string `json:"apiURL"`
}

// Load is used to.. load the configuration file.
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
