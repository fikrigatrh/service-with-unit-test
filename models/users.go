package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" validate:"required"`
	Email string `json:"email" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
	Status string `json:"status" validate:"required" gorm:"default:active"`
	Role string `json:"role" gorm:"default:user"`
}