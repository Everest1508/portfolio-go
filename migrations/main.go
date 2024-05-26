package main

import (
	"github.com/everest1508/portfolio/db"
	"github.com/everest1508/portfolio/models"
)

func init() {
	db.DBConnect()
}

func main() {
	db.DB.AutoMigrate(&models.Bio{})
	db.DB.AutoMigrate(&models.Contact{})
	db.DB.AutoMigrate(&models.Skill{})
	db.DB.AutoMigrate(&models.Project{})
}