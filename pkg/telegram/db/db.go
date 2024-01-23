package db

import (
	"fmt"
	"log"

	"github.com/sha1sof/bot_tg_menstruation.git/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Get().DBHost,
		config.Get().DBPort,
		config.Get().DBUser,
		config.Get().DBPassword,
		config.Get().DBName)

	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	log.Println("Connected to the database")

	if err := Migrate(db); err != nil {
		log.Fatal("Error applying migrations:", err)
	}

	log.Println("Database migrations applied successfully")
}
