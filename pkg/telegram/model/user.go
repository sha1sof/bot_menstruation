package model

import "time"

type User struct {
	ID               uint `gorm:"primaryKey;autoIncrement"`
	ChatID           uint
	UserName         string
	Gender           string
	TimeMenstruation time.Time
}

type Couple struct {
	ID    uint `gorm:"primaryKey;autoIncrement"`
	Man   uint
	Woman uint
}
