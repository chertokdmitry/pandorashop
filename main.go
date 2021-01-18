package main

import (
	"github.com/joho/godotenv"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/app"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	app.StartApplication()
}
