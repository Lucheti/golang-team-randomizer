package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func InitializeDB() *gorm.DB {
	//get DATABASE_URL environment variable
	connector := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: connector,
	}), &gorm.Config{})
	if err != nil {
		panic(any("failed to connect database"))
	}

	return db
}
