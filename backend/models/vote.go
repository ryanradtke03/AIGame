package models

import "gorm.io/gorm"

type Vote struct {
	gorm.Model
	RoomID     uint
	Room       Room
	VoterID    uint
	Voter      Player `gorm:"foreignKey:VoterID"`
	TargetID   uint
	Target     Player `gorm:"foreignKey:TargetID"`
	RoundNumber int
}