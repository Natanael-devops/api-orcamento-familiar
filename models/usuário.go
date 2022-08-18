package models

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}
