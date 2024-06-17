package config

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func ConnectToMysql() *gorm.DB {
	godotenv.Load()
	var err error
	var dsn = os.ExpandEnv("host=${DB_HOST} user=${DB_USERNAME} password=${DB_PASSWORD} dbname=${DB_NAME} port=${DB_PORT} sslmode=disable TimeZone=Asia/Jakarta")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database. Error: ", err)
	}

	return db
}
