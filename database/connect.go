package database

import (
	"log"
	"os"
	"react_go_catalog_web/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		log.Println("Connect To Database")
	}

	DB = db
	db.AutoMigrate(
		&models.User{},
		&models.Content{},
	)

}