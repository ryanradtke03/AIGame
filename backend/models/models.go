// Central model registry
package models

type ModelList []interface{}

var AllModels = ModelList{
	&Player{},
	&Room{},
	&RoomPlayer{},
	&ChatMessage{},
	&GameRound{},
	&Vote{},
}