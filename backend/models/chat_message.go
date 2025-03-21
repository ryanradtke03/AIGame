package models

import "gorm.io/gorm"

type ChatMessage struct {
	gorm.Model
	RoomID    uint
	Room      Room
	PlayerID  *uint      // Nullable if AI
	Username  string     `gorm:"not null"`
	Message   string     `gorm:"not null"`
}