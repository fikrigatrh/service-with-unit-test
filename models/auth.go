package models

import "gorm.io/gorm"

type Auth struct {
	gorm.Model
	AuthUUID string `gorm:"size:255;not null;" json:"auth_uuid"`
	Username string `gorm:"not null;" json:"username"`
	Email   string `gorm:"not null;" json:"email"`
	Role     string `json:"role" gorm:"default:user"`
}

type EncryptData struct {
	Encrypt string `json:"encrypt"`
}

type TokenStruct struct {
	Token string `json:"token"`
}
