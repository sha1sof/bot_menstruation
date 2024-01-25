package db

import (
	"log"

	"github.com/sha1sof/bot_menstruation.git/pkg/telegram/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	log.Println("Starting database migration...")

	err := db.AutoMigrate(&model.Man{}, &model.Woman{}, &model.Couple{})
	if err != nil {
		log.Printf("Error during migration: %v\n", err)
		return err
	}

	log.Println("Database migration completed successfully.")
	return nil
}
