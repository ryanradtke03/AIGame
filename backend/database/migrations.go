package database

import (
	"backend/models"
	"fmt"
)

// AutoMigrateTables creates tables automatically
func AutoMigrateTables() {
    if DB == nil {
        fmt.Println("❌ Error: Database connection is not initialized")
        return
    }

    DB = DB.Debug() // This will print SQL queries to the console
    err := DB.AutoMigrate(models.AllModels...) // Add other models here

    if err != nil {
        fmt.Println("❌ Error migrating tables:", err)
    } else {
        fmt.Println("✅ Database tables migrated successfully!")
    }
}