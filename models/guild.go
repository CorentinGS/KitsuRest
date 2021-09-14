package models

import (
	"gorm.io/gorm"
)

type Guild struct {
	gorm.Model
	GuildID   string `json:"guild_id" example:"879730620157292636"`
	GuildName string `json:"guild_name" example:"Les etoiles"`
	Vip       bool   `json:"vip" example:"true"`
}
