package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func ConnectToMysql() *gorm.DB {
	godotenv.Load()
	var err error
	var dbUser = os.Getenv("DB_USERNAME")
	var dbPassword = os.Getenv("DB_PASSWORD")
	var dbHost = os.Getenv("DB_HOST")
	var dbPort = os.Getenv("DB_PORT")
	var dbName = os.Getenv("DB_NAME")

	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// Print the DSN string for debugging
	fmt.Println("DSN:", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database. Error: ", err)
	}

	/**
	To run migration 'up' and 'down' action
	migrate -database "mysql://root@tcp(localhost:3306)/shift_local" path db/migrations up
	migrate -database "mysql://root@tcp(localhost:3306)/shift_local" path db/migrations down
	*/

	return db
}
