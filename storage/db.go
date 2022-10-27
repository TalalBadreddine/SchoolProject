package storage

import (
	"log"

	config "server/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func NewDB(params ...string) *gorm.DB {
	var err error
	conString := config.GetPostgresConnectionString()

	DB, err = gorm.Open(config.GetDBType(), conString)

	if err != nil {
		log.Panic(err)
	}

	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}
