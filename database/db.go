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
	Migrate()

	log.Println("Database connected and migrations completed.")
}

func Migrate() {
	if err := DB.AutoMigrate(&models.User{}, &models.Order{}); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	if DB.Migrator().HasColumn(&models.Order{}, "cod_fee") {
		if err := DB.Migrator().RenameColumn(&models.Order{}, "cod_fee", "cash_on_delivery_fee"); err != nil {
			log.Fatalf("Failed to rename column: %v", err)
		}
	}

	if err := DB.AutoMigrate(&models.User{}, &models.Order{}); err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
}
