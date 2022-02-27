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

type UserRequest struct {
	Email string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null" json:"password"`
}