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
}

func AddMan(chatID uint, userName, partnerName string) (bool, error) {
	existingMan := model.Man{}
	if dbInstance == nil {
		log.Println("Error: dbInstance is not initialized.")
		return false, errors.New("dbInstance is not initialized")
	}

	if err := dbInstance.Where("user_name = ?", userName).First(&existingMan).Error; err == nil {
		existingMan.PartnerManName = partnerName

		var partnerWoman model.Woman
		if result := dbInstance.Where("user_name = ?", partnerName).First(&partnerWoman); result.Error == nil {
			existingMan.PartnerManID = partnerWoman.ID
		}
		if err := dbInstance.Save(&existingMan).Error; err != nil {
			log.Printf("Error: when updating information about a man: %v\n", err)
			return false, err
		}
		return false, nil
	}

	newMan := model.Man{
		ChatID:         chatID,
		UserName:       userName,
		PartnerManName: partnerName,
	}

	var partnerWoman model.Woman
	if result := dbInstance.Where("user_name = ?", partnerName).First(&partnerWoman); result.Error == nil {
		newMan.PartnerManID = partnerWoman.ChatID
	} else {
		log.Printf("Error: when searching partner woman: %v\n", result.Error)
		return false, nil
	}

	if err := dbInstance.Create(&newMan).Error; err != nil {
		log.Printf("Error: when creating a record about a man: %v\n", err)
		return false, err
	}

	return true, nil
}

func AddWoman(chatID uint, userName, partnerWoman string, averageDuration uint, startMenstruation time.Time) (bool, error) {
	existingWoman := model.Woman{}
	if dbInstance == nil {
		log.Println("Error: dbInstance is not initialized.")
		return false, errors.New("dbInstance is not initialized")
	}

	if err := dbInstance.Where("user_name = ?", userName).First(&existingWoman).Error; err == nil {
		existingWoman.PartnerWomanName = partnerWoman
		existingWoman.StartMenstruation = startMenstruation
		var partnerMan model.Man
		if result := dbInstance.Where("user_name = ?", partnerWoman).First(&partnerMan); result.Error == nil {
			existingWoman.PartnerWomanID = partnerMan.ID
		}
		if err := dbInstance.Save(&existingWoman).Error; err != nil {
			log.Printf("Error: when updating information about a man: %v\n", err)
			return false, err
		}
		return false, nil
	}

	newWoman := model.Woman{
		ChatID:            chatID,
		UserName:          userName,
		PartnerWomanName:  partnerWoman,
		StartMenstruation: startMenstruation,
		AverageDuration:   averageDuration,
	}

	var partnerMan model.Man
	if result := dbInstance.Where("user_name = ?", partnerWoman).First(&partnerMan); result.Error == nil {
		newWoman.PartnerWomanID = partnerMan.ChatID
	} else {
		return false, nil
	}

	if err := dbInstance.Create(&newWoman).Error; err != nil {
		log.Printf("Error: when creating a record about a woman: %v\n", err)
		return false, err
	}

	return true, nil
}
