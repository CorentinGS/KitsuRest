package models

import (
	"gorm.io/gorm"
)

// Message is a model for message
type Message struct {
	gorm.Model
	GuildID   int64  `json:"guild_id" example:"1"`
	UserID    int64  `json:"user_id" example:"2"`
	ChannelID int64  `json:"channel_id" example:"1"`
	MessageID string `json:"message_id" example:"879730620157292638"`
	DateID    int64  `json:"date_id" example:"20210101"`
}
