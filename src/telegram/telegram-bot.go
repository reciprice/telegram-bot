package main

import "fmt"

func main() {
	config := Load("config/config.json")
	fmt.Println(config.Token)
}
