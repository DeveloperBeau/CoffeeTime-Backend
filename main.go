package main

import (
	"CoffeeTime-Go/api"
	"CoffeeTime-Go/db"
	"encoding/json"
	"log"
	"os"
)

type configuration struct {
	ServerAddress string `json:"webserver"`
}

func main() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	connectionString := "user=postgres dbname=coffeetime sslmode=disable"
	handler, dbErr := db.MakeDatabaseHandler(connectionString)
	if dbErr != nil {
		log.Fatal(err)
	}
	config := new(configuration)
	json.NewDecoder(file).Decode(config)
	log.Println("Starting web server on address ", config.ServerAddress)
	api.Run(config.ServerAddress, handler)
}
