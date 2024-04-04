package controllers

import (
	"net/http"
	"strconv"

	"github.com/desafio-estagio/src/services"
	"github.com/gin-gonic/gin"
)

// @Summary Get all players
// @Description Get a list of all players
// @Tags players
// @Accept json
// @Produce json
// @Success 200
// @Router /player [get]
// @Security Bearer
func GetPlayers(context *gin.Context) {
	players := services.GetPlayers()

	context.JSON(http.StatusOK, players)
}

// @Summary Get a player by ID
// @Description Get a player by its ID
// @Tags players
// @Accept json
// @Produce json
// @Param id path int true "Player ID"
// @Success 200
// @Router /player/{id} [get]
// @Security Bearer
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

type Player struct {
	Name     string
	Email    string
	Password string
	Class    string
}

// @Summary Create a new player
// @Description Create a new player
// @Tags players
// @Accept json
// @Produce json
// @Success 201
// @Router /player [post]
// @Security Bearer
// @Param player body Player true "Player object"
func CreatePlayer(context *gin.Context) {
	var player Player
	context.Bind(&player)

	res, error := services.CreatePlayer(player.Name, player.Email, player.Password, player.Class)

	if error != nil {
		context.Status(http.StatusBadRequest)
		return
	}

	context.JSON(http.StatusCreated, res)
}

type UPlayer struct {
	Name     string
	Email    string
	Password string
	Class    string
	Level    int
	GuildID  int
}

// @Summary Update a player
// @Description Update an existing player
// @Tags players
// @Accept json
// @Produce json
// @Param id path int true "Player ID"
// @Success 200
// @Router /player/{id} [put]
// @Security Bearer
// @Param player body UPlayer true "Player object"
func UpdatePlayer(context *gin.Context) {
	var player UPlayer
	context.Bind(&player)
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.Status(http.StatusBadRequest)
		return
	}

	err = services.UpdatePlayer(id, player.Name, player.Email, player.Password, player.Class, player.Level, player.GuildID)
	if err != nil {
		context.Status(http.StatusBadRequest)
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Player updated successfully"})
}

// @Summary Delete a player
// @Description Delete a player by its ID
// @Tags players
// @Accept json
// @Produce json
// @Param id path int true "Player ID"
// @Success 200 {string} string "Player deleted successfully"
// @Router /player/{id} [delete]
// @Security Bearer
func DeletePlayer(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.Status(http.StatusBadRequest)
		return
	}

	err = services.DeletePlayer(id)
	if err != nil {
		context.Status(http.StatusBadRequest)
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Player deleted successfully"})
}

// @Summary Join a player to a quest
// @Description Associate a player with a quest
// @Tags players
// @Accept json
// @Produce json
// @Param player_id query int true "Player ID"
// @Param quest_id query int true "Quest ID"
// @Success 200 {string} string "Player associated with the quest successfully"
// @Router /player/quest/join [post]
// @Security Bearer
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
		context.Status(http.StatusBadRequest)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "player associated with the quest successfully"})
}

// @Summary Join a player to a guild
// @Description Allows a player to join a guild
// @Tags players
// @Accept  json
// @Produce  json
// @Param player_id query int true "Player ID"
// @Param guild_id query int true "Guild ID"
// @Success 200
// @Router /player/guild/join/ [post]
// @Security Bearer
func JoinGuildController(context *gin.Context) {

	playerIDStr := context.Query("player_id")
	guildIDStr := context.Query("guild_id")

	playerID, err := strconv.Atoi(playerIDStr)
	if err != nil {
		context.Status(http.StatusBadRequest)
		return
	}

	guildID, err := strconv.Atoi(guildIDStr)
	if err != nil {
		context.Status(http.StatusBadRequest)
		return
	}

	err = services.JoinGuildService(playerID, guildID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Player joined guild successfully"})
}
