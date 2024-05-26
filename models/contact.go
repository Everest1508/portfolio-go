package models

import (
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
    ID      uint   `gorm:"primaryKey"`
    Name    string `gorm:"size:100;not null" json:"name"`
    Email   string `gorm:"size:100;not null" json:"email"`
    Message string `gorm:"type:text;not null" json:"message"`
}