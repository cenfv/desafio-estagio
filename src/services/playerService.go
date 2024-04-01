package services

import (
	"errors"

	"github.com/desafio-estagio/database"
	"github.com/desafio-estagio/src/models"
)

func GetPlayers() []models.Player {
	var players []models.Player
	database.DB.Find(&players)

	return players
}

func GetPlayer(id int) models.Player {
	var player models.Player
	database.DB.First(&player, id)

	return player
}

func CreatePlayer(name string, email string, password string, class string) (models.Player, error) {
	player := models.Player{
		Name:     name,
		Email:    email,
		Password: password,
		Class:    class,
		Level:    1,
	}

	res := database.DB.Create(&player)

	if res.Error != nil {
		return models.Player{}, res.Error
	}

	return player, res.Error
}

func UpdatePlayer(id int, name string, email string, password string, class string, level int, guildID int) error {
	var playerToUpdate models.Player
	database.DB.First(&playerToUpdate, id)

	res := database.DB.Model(&playerToUpdate).Updates(models.Player{
		Name:     name,
		Email:    email,
		Password: password,
		Class:    class,
		Level:    level,
		//GuildID: guildID,
	})

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func DeletePlayer(id int) error {
	var playerToDelete models.Player
	database.DB.First(&playerToDelete, id)

	res := database.DB.Delete(&playerToDelete)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func JoinQuest(playerID, questID int) error {
	var player models.Player
	var quest models.Quest

	database.DB.First(&player, playerID)
	database.DB.First(&quest, questID)

	err := database.DB.Model(&player).Association("Quests").Append(&quest)

	if err != nil {
		return errors.New("couldn't associate the player with the quest")
	}

	return nil
}

var jwtKey = []byte("your_secret_key")
