package models

import (
	"time"

	"gorm.io/gorm"
)

type Predicate struct {
	ID          int       `gorm:"primaryKey;autoIncrement;uniqueIndex;not null"`
	Name        string    `gorm:"size:50;not null"`
	Description string    `gorm:"size:255"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func SeedPredicates(db *gorm.DB) {
	predicates := []Predicate{
		{
			Name:        "Summa Cum Laude",
			Description: "Highest honors",
		},
		{
			Name:        "Magna Cum Laude",
			Description: "High honors",
		},
		{
			Name:        "Cum Laude",
			Description: "Honors",
		},
		{
			Name:        "Sangat Memuaskan",
			Description: "Very satisfactory",
		},
		{
			Name:        "Memuaskan",
			Description: "Satisfactory",
		},
		{
			Name:        "Cukup",
			Description: "Enough",
		},
	}

	for _, predicate := range predicates {
		db.FirstOrCreate(&predicate, Predicate{Name: predicate.Name})
	}
}
