package model

import "time"

type Man struct {
	ID             uint   `gorm:"primaryKey;autoIncrement"`
	ChatID         uint   `gorm:"uniqueIndex"`
	UserName       string `gorm:"uniqueIndex"`
	PartnerManName string
	PartnerManID   uint
}

type Woman struct {
	ID                uint   `gorm:"primaryKey;autoIncrement"`
	ChatID            uint   `gorm:"uniqueIndex"`
	UserName          string `gorm:"uniqueIndex"`
	PartnerWomanName  string
	PartnerWomanID    uint
	StartMenstruation time.Time
	AverageDuration   uint
}
