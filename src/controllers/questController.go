package controllers

import (
	"net/http"
	"strconv"

	"github.com/desafio-estagio/src/services"
	"github.com/gin-gonic/gin"
)

func GetQuests(context *gin.Context) {
	quests := services.GetQuests()

	context.JSON(http.StatusOK, quests)
}

func GetQuest(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.Status(http.StatusBadRequest)
		return
	}
	res := services.GetQuest(id)
	context.JSON(http.StatusOK, res)
}

func CreateQuest(context *gin.Context) {
	var quest struct {
		Name        string
		Description string
		Difficulty  int
		Reward      string
	}
	context.Bind(&quest)

	res, error := services.CreateQuest(quest.Name, quest.Description, quest.Difficulty, quest.Reward)

	if error != nil {
		context.Status(400)
		return
	}

	context.JSON(http.StatusOK, res)
}
func UpdateQuest(context *gin.Context) {
	var quest struct {
		Name        string
		Description string
		Difficulty  int
		Reward      string
	}
	context.Bind(&quest)
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.Status(http.StatusBadRequest)
		return
	}

	err = services.UpdateQuest(id, quest.Name, quest.Description, quest.Difficulty, quest.Reward)
	if err != nil {
		context.Status(400)
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Quest updated successfully"})
}

func DeleteQuest(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.Status(http.StatusBadRequest)
		return
	}

	err = services.DeleteQuest(id)
	if err != nil {
		context.Status(400)
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Quest deleted successfully"})
}
