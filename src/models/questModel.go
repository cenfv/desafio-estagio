package models

import (
	"gorm.io/gorm"
)

type Quest struct {
	gorm.Model
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Difficulty  int    `json:"difficulty" validate:"required"`
	Reward      string `json:"reward" validate:"required"`
}
