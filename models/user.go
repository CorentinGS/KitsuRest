package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `json:"user_name" example:"Yume"`
	UserID   string `json:"user_id" example:"879730620157292636"`
	Vip      bool   `json:"vip" example:"true"`
}
