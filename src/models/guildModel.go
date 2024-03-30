package models

import (
	"gorm.io/gorm"
)

type Guild struct {
	gorm.Model
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Members     []Player `gorm:"many2many:guild_members;"`
}
