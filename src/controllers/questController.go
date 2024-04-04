package controllers

import (
	"net/http"
	"strconv"

	"github.com/desafio-estagio/src/services"
	"github.com/gin-gonic/gin"
)

// @Summary List all quests
// @Description Get all quests
// @Tags quests
// @Accept json
// @Produce json
// @Success 200
// @Router /quests [get]
func GetQuests(context *gin.Context) {
	quests := services.GetQuests()

	context.JSON(http.StatusOK, quests)
}

// @Summary Get quest by ID
// @Description Get a quest by its ID
// @Tags quests
// @Accept json
// @Produce json
// @Param id path int true "Quest ID"
// @Success 200
// @Failure 400
// @Router /quests/{id} [get]
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

type Quest struct {
	Name        string
	Description string
	Difficulty  int
	Reward      string
}

// @Summary Create a new quest
// @Description Create a new quest
// @Tags quests
// @Accept json
// @Produce json
// @Param quest body Quest true "Quest object"
// @Success 201
// @Failure 400
// @Router /quests [post]
func CreateQuest(context *gin.Context) {
	var quest Quest
	context.Bind(&quest)

	res, error := services.CreateQuest(quest.Name, quest.Description, quest.Difficulty, quest.Reward)

	if error != nil {
		context.Status(http.StatusBadRequest)
		return
	}

	context.JSON(http.StatusCreated, res)
}

// @Summary Update an existing quest
// @Description Update an existing quest by ID
// @Tags quests
// @Accept json
// @Produce json
// @Param id path int true "Quest ID"
// @Param quest body Quest true "Quest object"
// @Success 200
// @Failure 400
// @Router /quests/{id} [put]
func UpdateQuest(context *gin.Context) {
	var quest Quest
	context.Bind(&quest)
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.Status(http.StatusBadRequest)
		return
	}

	err = services.UpdateQuest(id, quest.Name, quest.Description, quest.Difficulty, quest.Reward)
	if err != nil {
		context.Status(http.StatusBadRequest)
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Quest updated successfully"})
}

// @Summary Delete an existing quest
// @Description Delete an existing quest by ID
// @Tags quests
// @Accept json
// @Produce json
// @Param id path int true "Quest ID"
// @Success 200
// @Failure 400
// @Router /quests/{id} [delete]
func DeleteQuest(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.Status(http.StatusBadRequest)
		return
	}

	err = services.DeleteQuest(id)
	if err != nil {
		context.Status(http.StatusBadRequest)
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Quest deleted successfully"})
}
