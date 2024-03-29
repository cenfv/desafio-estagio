package main

import (
	"log"

	"github.com/desafio-estagio/database"
	"github.com/desafio-estagio/src/models"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.SetupDatabase()
}

func main() {
	database.DB.AutoMigrate(&models.Player{})
	database.DB.AutoMigrate(&models.Guild{})
	database.DB.AutoMigrate(&models.Quest{})
}
