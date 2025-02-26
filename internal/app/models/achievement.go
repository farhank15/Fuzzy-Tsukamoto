package models

import (
	"database/sql/driver"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Level string

const (
	LevelNasional      Level = "nasional"
	LevelInternasional Level = "internasional"
	LevelInternal      Level = "internal"
)

func (l *Level) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		*l = Level(v)
	case string:
		*l = Level(v)
	default:
		return errors.New("invalid type for Level")
	}
	return nil
}

func (l Level) Value() (driver.Value, error) {
	return string(l), nil
}

type Achievement struct {
	ID          int    `gorm:"primaryKey;autoIncrement;uniqueIndex;not null"`
	UserID      int    `gorm:"not null;index"`
	Title       string `gorm:"not null"`
	Certificate bool   `gorm:"not null"`
	Rank        int    `gorm:"not null"`
	Level       Level  `gorm:"not null;type:text"`
	Year        int    `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (a *Achievement) BeforeSave(tx *gorm.DB) (err error) {
	switch a.Level {
	case LevelNasional, LevelInternasional, LevelInternal:
		// valid level
	default:
		err = errors.New("invalid level")
	}
	return
}
