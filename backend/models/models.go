package models

// Import all models here
import "gorm.io/gorm"

type ModelList []interface{}

var AllModels = ModelList{
	&User{},
	// Add more models here, but no need to update anything else!
}

// Base models with GORM relationships
type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
}