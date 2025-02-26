package models

import "time"

type Thesis struct {
	ID        int    `gorm:"primaryKey;autoIncrement;uniqueIndex;not null;primaryKey"`
	UserID    int    `gorm:"not null;index"`
	Title     string `gorm:"size:255;not null"`
	Year      int
	Semester  int
	Value     string `gorm:"size:50"`
	Level     string `gorm:"size:50;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
