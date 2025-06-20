package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Role     string `json:"role" gorm:"not null"`
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
}
