package main

import (
	"log"

	"github.com/desafio-estagio/database"
	"github.com/desafio-estagio/src/controllers"
	"github.com/gin-gonic/gin"
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
	r := gin.Default()

	r.GET("/player", controllers.GetPlayers)
	r.GET("/player/:id", controllers.GetPlayer)
	r.POST("/player", controllers.CreatePlayer)
	r.PUT("/player/:id", controllers.UpdatePlayer)
	r.DELETE("/player/:id", controllers.DeletePlayer)
	r.POST("/player/quest/join/", controllers.JoinQuestController)

	r.GET("/quest", controllers.GetQuests)
	r.GET("/quest/:id", controllers.GetQuest)
	r.POST("/quest", controllers.CreateQuest)
	r.PUT("/quest/:id", controllers.UpdateQuest)
	r.DELETE("/quest/:id", controllers.DeleteQuest)

	r.Run()
}
