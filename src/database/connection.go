package database

import (
	"fmt"
	"golang-rest-api/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var (
	DbHost     = "localhost"
	DbPort     = "5432"
	DbUser     = os.Getenv("DB_USERNAME")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbName     = os.Getenv("DB_NAME")
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DbHost, DbPort, DbUser, DbPassword, DbName)

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("could not connect to database...")
	}

	DB = connection

	panic(connection.AutoMigrate(&models.User{}))

}
