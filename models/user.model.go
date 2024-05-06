package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string `gorm:"not null" json:"name"`
	EmailAddress string `gorm:"not null;unique" json:"emailAddress"`
	Password     string `gorm:"not null;" json:"password"`
	HasSignedUp  bool   `gorm:"default:false" json:"hasSignedUp"`
}
