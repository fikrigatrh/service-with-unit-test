package models

import "gorm.io/gorm"

type Auth struct {
	gorm.Model
	AuthUUID string `gorm:"size:255;not null;" json:"auth_uuid"`
	Username string `gorm:"not null;" json:"username"`
	UserID   string `gorm:"not null;" json:"user_id"`
	SessID   string `json:"sess_id"`
	Token    string `json:"token"`
}
