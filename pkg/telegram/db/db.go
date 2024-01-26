package db

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/sha1sof/bot_menstruation.git/config"
	"github.com/sha1sof/bot_menstruation.git/pkg/telegram/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func InitDB() {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Get().DBHost,
		config.Get().DBPort,
		config.Get().DBUser,
		config.Get().DBPassword,
		config.Get().DBName)

	var err error
	dbInstance, err = gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	log.Println("Connected to the database")

	if err := Migrate(dbInstance); err != nil {
		log.Fatal("Error applying migrations:", err)
	}

	log.Println("Database migrations applied successfully")
}

func AddMan(chatID uint, userName, partnerName string) error {
	newMan := model.Man{
		ChatID:         chatID,
		UserName:       userName,
		PartnerManName: partnerName,
	}

	if dbInstance == nil {
		log.Println("Error: dbInstance is not initialized.")
		return errors.New("dbInstance is not initialized")
	}

	err := dbInstance.Create(&newMan).Error
	if err != nil {
		log.Printf("Error: when creating a record about a man: %v\n", err)
		return err
	}

	return nil
}

func AddWoman(chatID uint, userName, partnerWoman, averageDuration string, startMenstruation time.Time) error {
	newWoman := model.Woman{
		ChatID:            chatID,
		UserName:          userName,
		PartnerWomanName:  partnerWoman,
		StartMenstruation: startMenstruation,
	}

	if dbInstance == nil {
		log.Println("Error: dbInstance is not initialized.")
		return errors.New("dbInstance is not initialized")
	}

	err := dbInstance.Create(&newWoman).Error
	if err != nil {
		log.Printf("Error: when creating a record about a woman: %v\n", err)
		return err
	}

	return nil
}
