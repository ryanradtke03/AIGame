package models

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Code       string `gorm:"uniqueIndex;not null"` // Join code
	OwnerID    uint
	Owner      Player `gorm:"foreignKey:OwnerID"`
	Rounds     int    `gorm:"default:3"`
	Status     string `gorm:"default:'waiting'"` // waiting, in_progress, finished
}