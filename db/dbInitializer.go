package db

import (
	// "gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect() {
	var err error
	
	// connStr := "user='ritesh' password=ryiuTVcC36QD host=ep-curly-fog-a28s2wmg.eu-central-1.pg.koyeb.app dbname='koyebdb'"
	// DB, err = gorm.Open(postgres.Open(connStr),&gorm.Config{})
	DB, err = gorm.Open(sqlite.Open("portfolio.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}