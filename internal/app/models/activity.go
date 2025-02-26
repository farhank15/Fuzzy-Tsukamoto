package models

import "time"

type Activity struct {
	ID           int    `gorm:"primaryKey;autoIncrement;uniqueIndex;not null;primaryKey"`
	UserID       int    `gorm:"not null;index"`
	Organization string `gorm:"not null"`
	Year         int    `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
