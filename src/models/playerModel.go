package models

import (
	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Name    string  `json:"name" validate:"required"`
	Email   string  `json:"email" validate:"required"`
	Class   string  `json:"class" validate:"required"`
	Level   int     `json:"level" validate:"required"`
	GuildID int     `json:"guild_id"`
	Quests  []Quest `gorm:"many2many:player_quests;"`
}
