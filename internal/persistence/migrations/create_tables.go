package main

import (
	"Orgzz-Tasks/config"
	"Orgzz-Tasks/internal/models"
)

func main() {
	db := config.OpenDbConnection()
	// Migrate the schema
	db.AutoMigrate(&models.Task{})
}
