package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type Config struct {
	DBUSER     string `yaml:"DBUSER"`
	DBPASSWORD string `yaml:"DBPASSWORD"`
	DBHOST     string `yaml:"DBHOST"`
	DBPORT     string `yaml:"DBPORT"`
	DBNAME     string `yaml:"DBNAME"`
	DBTYPE     string `yaml:"DBTYPE"`
	SERVERPORT string `yaml:"SERVERPORT"`
}

var ConfigInstance Config

func (c *Config) GetDbConf() *Config {

	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func GetPostgresConnectionString() string {
	ConfigInstance.GetDbConf()
	dataBase := fmt.Sprintf("host=%s port=%s dbname=%s password=%s sslmode=disable", ConfigInstance.DBHOST, ConfigInstance.DBPORT, ConfigInstance.DBNAME, ConfigInstance.DBPASSWORD)
	return dataBase
}
