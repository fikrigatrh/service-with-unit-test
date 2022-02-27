package models

import "gorm.io/gorm"

type AboutUs struct {
	gorm.Model
	Profil string `json:"profil"`
	Visi   string `json:"visi"`
	Misi   string `json:"misi"`
	Motto  string `json:"motto"`
}

type Footer struct {
}

type Header struct {
}
