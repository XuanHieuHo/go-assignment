package main

import (
	"log"

	"github.com/XuanHieuHo/go-assignment/config"
	"github.com/XuanHieuHo/go-assignment/models"
)

func main() {
	db := config.DatabaseConnect()
	log.Println("Running database migrations...")
	models := []any{
		&models.User{},
		&models.FriendShip{},
	}

	if err := db.AutoMigrate(models...); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
}
