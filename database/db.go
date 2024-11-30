package database

import (
	"log"
	"order-management/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:root@tcp(db:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	log.Println("Running migrations...")
	err = DB.AutoMigrate(&models.User{}, &models.Order{})
	if err != nil {
		log.Fatal("Failed to run migrations: ", err)
	}

	log.Println("Database connected and migrations completed.")
}
