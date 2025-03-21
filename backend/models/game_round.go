package models

import "gorm.io/gorm"

type GameRound struct {
	gorm.Model
	RoomID      uint
	Room        Room
	RoundNumber int
	AIRevealed  bool `gorm:"default:false"`
}