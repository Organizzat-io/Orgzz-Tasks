package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func OpenDbConnection() *gorm.DB {
	dsn := "postgresql://doadmin:AVNS_HXrvyZ0xXpIgkRQNoaH@db-postgresql-nyc3-51555-do-user-15529344-0.c.db.ondigitalocean.com:25060/defaultdb?sslmode=require"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}
