package config

import (
	"fmt"
	"log"
	"os"

	"github.com/XuanHieuHo/go-assignment/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnect() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cannot load config!: ", err)
	}
	var (
		port     = os.Getenv("DB_PORT")
		host     = os.Getenv("DB_HOST")
		user     = os.Getenv("DB_USERNAME")
		password = os.Getenv("DB_PASSWORD")
		dbName   = os.Getenv("DB_NAME")
	)

	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	fmt.Printf("Database connected successfully: host=%s port=%s \n", host, port)

	models := []any{
		&models.User{},
		&models.FriendShip{},
	}
	err = db.AutoMigrate(models...)
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
	err = db.Exec(`
				ALTER TABLE friend_ships DROP CONSTRAINT IF EXISTS unique_friendship;
				ALTER TABLE friend_ships DROP CONSTRAINT IF EXISTS check_friendship;

				ALTER TABLE friend_ships ADD CONSTRAINT unique_friendship UNIQUE (user_id, friend_id); 
				ALTER TABLE friend_ships ADD CONSTRAINT check_friendship CHECK (user_id < friend_id);
			`).Error
	if err != nil {
		log.Fatal("Failed to add constraints:", err)
	}

	err = db.Exec(`
	        	CREATE INDEX IF NOT EXISTS idx_user_friend ON friend_ships (user_id, friend_id);
	    `).Error
	if err != nil {
		log.Fatal("Failed to add index:", err)
	}

	return db
}
