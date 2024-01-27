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
	existingMan := model.Man{}
	if dbInstance == nil {
		log.Println("Error: dbInstance is not initialized.")
		return errors.New("dbInstance is not initialized")
	}

	if err := dbInstance.Where("user_name = ?", userName).First(&existingMan).Error; err == nil {
		existingMan.PartnerManName = partnerName
		if err := dbInstance.Save(&existingMan).Error; err != nil {
			log.Printf("Error: when updating information about a man: %v\n", err)
			return err
		}
		return nil
	}

	newMan := model.Man{
		ChatID:         chatID,
		UserName:       userName,
		PartnerManName: partnerName,
	}

	if err := dbInstance.Create(&newMan).Error; err != nil {
		log.Printf("Error: when creating a record about a man: %v\n", err)
		return err
	}

	return nil
}

func AddWoman(chatID uint, userName, partnerWoman, averageDuration string, startMenstruation time.Time) error {
	existingWoman := model.Woman{}
	if dbInstance == nil {
		log.Println("Error: dbInstance is not initialized.")
		return errors.New("dbInstance is not initialized")
	}

	if err := dbInstance.Where("user_name = ?", userName).First(&existingWoman).Error; err == nil {
		existingWoman.PartnerWomanName = partnerWoman
		existingWoman.StartMenstruation = startMenstruation

		if err := dbInstance.Save(&existingWoman).Error; err != nil {
			log.Printf("Error: when updating information about a woman: %v\n", err)
			return err
		}
		return nil
	}

	newWoman := model.Woman{
		ChatID:            chatID,
		UserName:          userName,
		PartnerWomanName:  partnerWoman,
		StartMenstruation: startMenstruation,
	}

	if err := dbInstance.Create(&newWoman).Error; err != nil {
		log.Printf("Error: when creating a record about a woman: %v\n", err)
		return err
	}

	return nil
}
