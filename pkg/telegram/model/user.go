package model

import "time"

type Man struct {
	ID             uint   `gorm:"primaryKey;autoIncrement"`
	ChatID         uint   `gorm:"uniqueIndex"`
	UserName       string `gorm:"uniqueIndex"`
	PartnerManName string
}

type Couple struct {
	ID    uint `gorm:"primaryKey;autoIncrement"`
	Man   uint
	Woman uint
}

type Woman struct {
	ID                uint   `gorm:"primaryKey;autoIncrement"`
	ChatID            uint   `gorm:"uniqueIndex"`
	UserName          string `gorm:"uniqueIndex"`
	PartnerWomanName  string
	StartMenstruation time.Time
	AverageDuration   string
}
