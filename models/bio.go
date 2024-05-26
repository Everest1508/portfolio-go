package models

import (
	"gorm.io/gorm"
)

type Bio struct {
    gorm.Model
	ID      uint   `gorm:"primaryKey"`
    Name    string `gorm:"size:100;not null" json:"name"`
    Summary string `gorm:"type:text;not null" json:"summary"`
}
