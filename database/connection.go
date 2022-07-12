package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	DB = Connect()
	return DB
}

func Connect() *gorm.DB {
	err := godotenv.Load(".env")

	User := os.Getenv("DB_USER")
	Pass := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	Host := os.Getenv("DB_HOST")
	Port := os.Getenv("DB_PORT")

	DNS := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Mexico_City",
		Host, User, Pass, DbName, Port)

	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})

	if err != nil {
		panic("Connection to the DataBase: lost")
	}
	return DB
}
