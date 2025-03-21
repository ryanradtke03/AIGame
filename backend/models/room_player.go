package models

import "gorm.io/gorm"

type RoomPlayer struct {
	gorm.Model
	RoomID     uint
	Room       Room
	PlayerID   *uint       // Nullable for AI
	Player     *Player     // Nullable for AI
	Username   string      `gorm:"not null"`
	SessionID  string      `gorm:"not null"` // to reconnect
	IsAI       bool        `gorm:"default:false"`
	Points     int         `gorm:"default:0"`
	Score      int         `gorm:"default:0"` // Optional: if used separately
}