package config

import (
	"fmt"
)

const (
	DBUSER     = "root"
	DBPASSWORD = "root"
	DBHOST     = "0.0.0.0"
	DBPORT     = "5432"
	DBNAME     = "postgres"
	DBTYPE     = "postgres"
)

func GetDBType() string {
	return DBTYPE
}

func GetPostgresConnectionString() string {
	dataBase := fmt.Sprintf("host=%s port=%s dbname=%s password=%s sslmode=disable", DBHOST, DBPORT, DBNAME, DBPASSWORD)
	return dataBase
}

// .yml
