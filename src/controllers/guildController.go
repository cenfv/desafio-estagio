package controllers

import (
	"net/http"
	"strconv"

	"github.com/desafio-estagio/src/services"
	"github.com/gin-gonic/gin"
)

func GetGuilds(context *gin.Context) {
	guilds := services.GetGuilds()

	context.JSON(http.StatusOK, guilds)
}

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

func CreateGuild(context *gin.Context) {
	var guild struct {
		Name        string
		Description string
		Members     []struct{ ID int }
	}
	context.Bind(&guild)

	res, error := services.CreateGuild(guild.Name, guild.Description, guild.Members)

	if error != nil {
		context.Status(http.StatusBadRequest)
		return
	}

	context.JSON(http.StatusCreated, res)
}
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
