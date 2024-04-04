package main

import (
	"log"

	"github.com/desafio-estagio/database"
	"github.com/desafio-estagio/docs"
	"github.com/desafio-estagio/src/controllers"
	"github.com/desafio-estagio/src/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.SetupDatabase()
	services.CreateInitialUser()
}

// @contact.name   API Support
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	r := gin.Default()
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Version = "1.0"

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/login", controllers.Login)

	r.GET("/player", controllers.AuthMiddleware(), controllers.GetPlayers)
	r.GET("/player/:id", controllers.AuthMiddleware(), controllers.GetPlayer)
	r.POST("/player", controllers.AuthMiddleware(), controllers.CreatePlayer)
	r.PUT("/player/:id", controllers.AuthMiddleware(), controllers.UpdatePlayer)
	r.DELETE("/player/:id", controllers.AuthMiddleware(), controllers.DeletePlayer)
	r.POST("/player/quest/join/", controllers.AuthMiddleware(), controllers.JoinQuestController)
	r.POST("/player/guild/join/", controllers.AuthMiddleware(), controllers.JoinGuildController)

	r.GET("/quest", controllers.AuthMiddleware(), controllers.GetQuests)
	r.GET("/quest/:id", controllers.AuthMiddleware(), controllers.GetQuest)
	r.POST("/quest", controllers.AuthMiddleware(), controllers.CreateQuest)
	r.PUT("/quest/:id", controllers.AuthMiddleware(), controllers.UpdateQuest)
	r.DELETE("/quest/:id", controllers.AuthMiddleware(), controllers.DeleteQuest)

	r.GET("/guild", controllers.AuthMiddleware(), controllers.GetGuilds)
	r.GET("/guild/:id", controllers.AuthMiddleware(), controllers.GetGuild)
	r.POST("/guild", controllers.AuthMiddleware(), controllers.CreateGuild)
	r.PUT("/guild/:id", controllers.AuthMiddleware(), controllers.UpdateGuild)
	r.DELETE("/guild/:id", controllers.AuthMiddleware(), controllers.DeleteGuild)
	r.DELETE("/guild/kick", controllers.AuthMiddleware(), controllers.KickPlayerController)

	r.Run()
}
