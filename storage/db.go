package storage

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"server/config"
)

func NewDB() *gorm.DB {
	var DB *gorm.DB
	var err error
	conString := config.GetPostgresConnectionString()

	DB, err = gorm.Open(postgres.Open(conString), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}

	return DB.Debug()
}
