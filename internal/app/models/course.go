package models

import "time"

type Course struct {
	ID           int    `gorm:"primaryKey;autoIncrement;uniqueIndex;not null;primaryKey"`
	Code         string `gorm:"not null"`
	CourseName   string `gorm:"not null"`
	CreditCourse int    `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
