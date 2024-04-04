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

func JoinGuildService(playerID, guildID int) error {
	var player models.Player
	var guild models.Guild

	res := database.DB.First(&player, playerID)
	if res.Error != nil {
		return errors.New("player not found")
	}

	res = database.DB.First(&guild, guildID)
	if res.Error != nil {
		return errors.New("guild not found")
	}

	for _, member := range guild.Members {
		if member.ID == uint(playerID) {
			return errors.New("player already in the guild")
		}
	}

	res = database.DB.Model(&player).Update("GuildID", guildID)
	if res.Error != nil {
		return errors.New("failed to join guild")
	}

	return nil
}
