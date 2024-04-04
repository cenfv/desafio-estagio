package controllers

import (
	"net/http"
	"strconv"

	"github.com/desafio-estagio/src/services"
	"github.com/gin-gonic/gin"
)

// @Summary Get all guilds
// @Description Get a list of all guilds
// @Tags guilds
// @Accept json
// @Produce json
// @Success 200
// @Router /guild [get]
// @Security Bearer
func GetGuilds(context *gin.Context) {
	guilds := services.GetGuilds()

	context.JSON(http.StatusOK, guilds)
}

// @Summary Get a guild by ID
// @Description Get a guild by its ID
// @Tags guilds
// @Accept json
// @Produce json
// @Param id path int true "Guild ID"
// @Success 200
// @Router /guild/{id} [get]
// @Security Bearer
func GetGuild(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.Status(http.StatusBadRequest)
		return
	}
	res := services.GetGuild(id)
	context.JSON(http.StatusOK, res)
}

type Guild struct {
	Name        string
	Description string
	Members     []struct{ ID int }
}

// @Summary Create a new guild
// @Description Create a new guild
// @Tags guilds
// @Accept json
// @Produce json
// @Param guild body Guild true "Guild object"
// @Success 201
// @Router /guild [post]
// @Security Bearer
func CreateGuild(context *gin.Context) {
	var guild Guild
	context.Bind(&guild)

	res, error := services.CreateGuild(guild.Name, guild.Description, guild.Members)

	if error != nil {
		context.Status(http.StatusBadRequest)
		return
	}

	context.JSON(http.StatusCreated, res)
}

// @Summary Update a guild
// @Description Update an existing guild
// @Tags guilds
// @Accept json
// @Produce json
// @Param id path int true "Guild ID"
// @Param guild body Guild true "Guild object"
// @Success 200
// @Router /guild/{id} [put]
// @Security Bearer
func UpdateGuild(context *gin.Context) {
	var guild struct {
		Name        string
		Description string
		Members     []struct{ ID int }
	}
	context.Bind(&guild)
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.Status(http.StatusBadRequest)
		return
	}

	err = services.UpdateGuild(id, guild.Name, guild.Description, guild.Members)
	if err != nil {
		context.Status(http.StatusBadRequest)
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Guild updated successfully"})
}

// @Summary Delete a guild
// @Description Delete a guild by its ID
// @Tags guilds
// @Accept json
// @Produce json
// @Param id path int true "Guild ID"
// @Success 200
// @Router /guild/{id} [delete]
// @Security Bearer
func DeleteGuild(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.Status(http.StatusBadRequest)
		return
	}

	err = services.DeleteGuild(id)
	if err != nil {
		context.Status(http.StatusBadRequest)
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Guild deleted successfully"})
}

// @Summary Remove a player from a guild
// @Description Remove a player from a guild by player ID and guild ID
// @Tags guilds
// @Accept  json
// @Produce  json
// @Param player_id query int true "Player ID"
// @Param guild_id query int true "Guild ID"
// @Success 200
// @Router /guild/kick [delete]
// @Security Bearer
func KickPlayerController(context *gin.Context) {
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

	err = services.KickPlayerService(playerID, guildID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Player removed from guild successfully"})
}
