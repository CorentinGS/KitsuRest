package models

import (
	"gorm.io/gorm"
)

type Channel struct {
	gorm.Model
	GuildID     uint   `json:"guild_id" example:"1"`
	ChannelName string `json:"channel_name" example:"main"`
	ChannelID   string `json:"channel_id" example:"879730620157292636"`

	Ignored bool `json:"ignored" example:"false"`
}
