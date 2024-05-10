package services

import (
	"cdc_mailer/models"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dbURL := os.Getenv("GROM_DB_CONFIG")

	DB, err = gorm.Open(postgres.Open(dbURL))

	if err != nil {
		panic("Failed to connect to DB!!")
	}

	fmt.Println("Connected to DB successfully!!")
	errMigrate := DB.AutoMigrate(&models.User{})
	if errMigrate != nil {
		panic("Error in DB migration, Table: User")
	}
	_ = DB.AutoMigrate(&models.Company{})
	fmt.Println("Database has been migrated!!")
}
