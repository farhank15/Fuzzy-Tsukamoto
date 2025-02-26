package models

import "time"

type Users struct {
	ID        int    `gorm:"primaryKey;autoIncrement;uniqueIndex;not null"`
	Username  string `gorm:"size:50;index"`
	Name      string `gorm:"size:50"`
	Nim       string `gorm:"size:20;uniqueIndex;not null"`
	Password  string `gorm:"size:255;not null"`
	StartYear int    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
