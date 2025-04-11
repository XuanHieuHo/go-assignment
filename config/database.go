package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnect() *gorm.DB {
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	var (
		port     = config.DBConfig.DBPort
		host     = config.DBConfig.DBHost
		user     = config.DBConfig.DBUsername
		password = config.DBConfig.DBPassword
		dbName   = config.DBConfig.DBName
	)

	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	fmt.Printf("Database connected successfully: host=%s port=%s \n", host, port)

	

	return db
}
