package models

import (
	"gorm.io/gorm"
)

type Skill struct {
	gorm.Model
    ID     uint   `gorm:"primaryKey"`
    Name   string `gorm:"size:100;not null" json:"name"`
    Level  string `gorm:"size:50" json:"level"` // Optional: e.g., Beginner, Intermediate, Expert
}
