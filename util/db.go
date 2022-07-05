package util

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config *Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.DBURI), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
		return nil, err
	}

	return db, nil
}
