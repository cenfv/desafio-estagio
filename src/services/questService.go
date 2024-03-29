package services

import (
	"github.com/desafio-estagio/database"
	"github.com/desafio-estagio/src/models"
)

func GetQuests() []models.Quest {
	var Quests []models.Quest
	database.DB.Find(&Quests)

	return Quests
}
func GetQuest(id int) models.Quest {
	var quest models.Quest
	database.DB.First(&quest, id)

	return quest
}

func CreateQuest(name string, description string, difficulty int, reward string) (models.Quest, error) {
	quest := models.Quest{
		Name:        name,
		Description: description,
		Difficulty:  difficulty,
		Reward:      reward,
	}

	res := database.DB.Create(&quest)

	if res.Error != nil {
		return models.Quest{}, res.Error
	}

	return quest, res.Error
}

func UpdateQuest(id int, name string, description string, difficulty int, reward string) error {
	var questToUpdate models.Quest
	database.DB.First(&questToUpdate, id)

	res := database.DB.Model(&questToUpdate).Updates(models.Quest{
		Name:        name,
		Description: description,
		Difficulty:  difficulty,
		Reward:      reward,
	})

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func DeleteQuest(id int) error {
	var questToDelete models.Quest
	database.DB.First(&questToDelete, id)

	res := database.DB.Delete(&questToDelete)

	if res.Error != nil {
		return res.Error
	}

	return nil
}
