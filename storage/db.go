package storage

import (
	"log"

	config "server/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB(params ...string) *gorm.DB {
	var err error
	conString := config.GetPostgresConnectionString()

	// DB, err = gorm.Open(config.GetDBType(), conString)
	DB, err = gorm.Open(postgres.Open(conString), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}

	return DB.Debug()
}

func GetDBInstance() *gorm.DB {
	return DB.Debug()
}
