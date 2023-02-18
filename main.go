package main

import (
	_ "github.com/joho/godotenv/autoload"
	"log"

	"Test_MNC_2/config"
)

func main() {
	err := config.Routers.Run()
	if err != nil {
		log.Fatal(err)
	}
}
