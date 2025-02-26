package models

import "time"

type Academic struct {
	ID              int     `gorm:"primaryKey;autoIncrement;uniqueIndex;not null;primaryKey"`
	UserID          int     `gorm:"not null;index"`
	Ipk             float64 `gorm:"size:50"`
	Ips             float64 `gorm:"size:50"`
	RepeatedCourses int     `gorm:"not null"`
	Semester        int     `gorm:"not null"`
	Year            int     `gorm:"not null"`
	PredicateID     *int    `gorm:"default:null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
