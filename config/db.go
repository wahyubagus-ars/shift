package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func connectToMysql() *gorm.DB {
	var err error
	dsn := os.Getenv("DB_DSN")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database. Error: ", err)
	}

	return db
}
