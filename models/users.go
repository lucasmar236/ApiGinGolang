package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nome    string `json:"nome"`
	Email   string `json:"email"`
	Usuario string `json:"usuario"`
	Senha   string `json:"senha,omitempty"`
}
