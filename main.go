package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/DeveloperBeau/CoffeeTime-Backend/api"
	"github.com/DeveloperBeau/CoffeeTime-Backend/db"
)

type configuration struct {
	ServerAddress string `json:"webserver"`
	isProduction  bool   `json:"isProduction"`
}

func main() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	config := new(configuration)
	json.NewDecoder(file).Decode(config)
	connectionString := "user=postgres dbname=coffeetime sslmode=disable"
	handler, dbErr := db.MakeDatabaseHandler(connectionString, config.isProduction)
	if dbErr != nil {
		log.Fatal(err)
	}
	log.Println("Starting web server on address ", config.ServerAddress)
	api.Run(config.ServerAddress, handler)
}
