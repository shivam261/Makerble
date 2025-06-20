package models

import "gorm.io/gorm"

type Patient struct {
	gorm.Model
	Name    string `json:"name" gorm:"not null"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address" gorm:"not null"`
	Age     int    `json:"age" gorm:"not null"`
	Disease string `json:"disease" gorm:"not null"`
	Status  string `json:"status" gorm:"not null;default:'Scheduled'"`
}
