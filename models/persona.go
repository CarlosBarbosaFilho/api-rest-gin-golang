package models

import "gorm.io/gorm"

type Persona struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Document string `json:"document"`
}
