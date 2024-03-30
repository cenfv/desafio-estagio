package services

import (
	"github.com/desafio-estagio/database"
	"github.com/desafio-estagio/src/models"
)

func GetGuilds() []models.Guild {
	var guilds []models.Guild
	database.DB.Preload("Members").Find(&guilds)

	return guilds
}
func GetGuild(id int) models.Guild {
	var guild models.Guild
	database.DB.Preload("Members").First(&guild, id)

	return guild
}

func CreateGuild(name string, description string, members []struct{ ID int }) (models.Guild, error) {

	var structuredMembers []models.Player
	for _, member := range members {
		var player models.Player
		database.DB.First(&player, member.ID)
		structuredMembers = append(structuredMembers, player)
	}

	guild := models.Guild{
		Name:        name,
		Description: description,
		Members:     structuredMembers,
	}

	res := database.DB.Create(&guild)

	if res.Error != nil {
		return models.Guild{}, res.Error
	}

	return guild, res.Error
}

func UpdateGuild(id int, name string, description string, members []struct{ ID int }) error {
	var guildToUpdate models.Guild
	database.DB.First(&guildToUpdate, id)

	database.DB.Model(&guildToUpdate).Association("Members").Clear()

	var structuredMembers []models.Player
	for _, member := range members {
		var player models.Player
		database.DB.First(&player, member.ID)
		structuredMembers = append(structuredMembers, player)
	}
	database.DB.Model(&guildToUpdate).Association("Members").Append(structuredMembers)

	guildToUpdate.Name = name
	guildToUpdate.Description = description

	res := database.DB.Save(&guildToUpdate)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func DeleteGuild(id int) error {
	var guildToDelete models.Guild
	database.DB.First(&guildToDelete, id)

	res := database.DB.Delete(&guildToDelete)

	if res.Error != nil {
		return res.Error
	}

	return nil
}
