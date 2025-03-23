package models

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	Username  string `gorm:"not null"`
	SessionID string `gorm:"uniqueIndex;not null"` // UUID to allow reconnecting
}