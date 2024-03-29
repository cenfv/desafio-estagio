package controllers

import (
	"net/http"
	"strconv"

	"github.com/desafio-estagio/src/services"
	"github.com/gin-gonic/gin"
)

func GetPlayers(context *gin.Context) {
	players := services.GetPlayers()

	context.JSON(http.StatusOK, players)
}

func GetPlayer(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.Status(http.StatusBadRequest)
		return
	}
	res := services.GetPlayer(id)
	context.JSON(http.StatusOK, res)
}

func CreatePlayer(context *gin.Context) {
	var player struct {
		Name  string
		Email string
		Class string
	}
	context.Bind(&player)

	res, error := services.CreatePlayer(player.Name, player.Email, player.Class)

	if error != nil {
		context.Status(400)
		return
	}

	context.JSON(http.StatusOK, res)
}
func UpdatePlayer(context *gin.Context) {
	var player struct {
		Name    string
		Email   string
		Class   string
		Level   int
		GuildID int
	}
	context.Bind(&player)
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.Status(http.StatusBadRequest)
		return
	}

	err = services.UpdatePlayer(id, player.Name, player.Email, player.Class, player.Level, player.GuildID)
	if err != nil {
		context.Status(400)
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Player updated successfully"})
}

func DeletePlayer(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.Status(http.StatusBadRequest)
		return
	}

	err = services.DeletePlayer(id)
	if err != nil {
		context.Status(400)
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Player deleted successfully"})
}

func JoinQuestController(context *gin.Context) {

	playerIDStr := context.Query("player_id")
	questIDStr := context.Query("quest_id")

	playerID, err := strconv.Atoi(playerIDStr)
	if err != nil {
		context.Status(http.StatusBadRequest)
		return
	}

	questID, err := strconv.Atoi(questIDStr)
	if err != nil {
		context.Status(http.StatusBadRequest)
		return
	}

	err = services.JoinQuest(playerID, questID)
	if err != nil {
		context.Status(400)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "player associated with the quest successfully"})
}
