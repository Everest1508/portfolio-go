package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:100;not null" json:"name"`
	Description string `gorm:"type:text;not null" json:"description"`
	GitHubURL   string `gorm:"size:255" json:"github_url"`
	LiveDemoURL string `gorm:"size:255" json:"live_demo_url"`
}
