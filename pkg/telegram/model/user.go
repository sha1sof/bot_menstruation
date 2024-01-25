package model

import "time"

type Man struct {
	ID       uint `gorm:"primaryKey;autoIncrement"`
	ChatID   uint
	UserName string
	Partner  uint
}

type Couple struct {
	ID    uint `gorm:"primaryKey;autoIncrement"`
	Man   uint
	Woman uint
}

type Woman struct {
	ID                uint `gorm:"primaryKey;autoIncrement"`
	ChatID            uint
	UserName          string
	StartMenstruation time.Time
	Partner           uint
}
