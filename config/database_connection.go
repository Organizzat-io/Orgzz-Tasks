package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func OpenDbConnection() *gorm.DB {
	dsn := "host=localhost user=postgres password=053079 dbname=orgzz-tasks port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}
