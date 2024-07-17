package config

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var ctx = context.Background()

func ConnectToMysql() *gorm.DB {
	var err error
	var dbUser = os.Getenv("DB_MYSQL_USERNAME")
	var dbPassword = os.Getenv("DB_MYSQL_PASSWORD")
	var dbHost = os.Getenv("DB_MYSQL_HOST")
	var dbPort = os.Getenv("DB_MYSQL_PORT")
	var dbName = os.Getenv("DB_MYSQL_NAME")

	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database MySql. Error: ", err)
	}

	/**
	Create a new migration
		migrate create -ext sql -dir db/migrations/mysql -seq <MIGRATION_NAME>

	To run migration 'up' and 'down' action
		migrate -database "mysql://root@tcp(localhost:3306)/shift_local" -path db/migrations/mysql up
		migrate -database "mysql://root@tcp(localhost:3306)/shift_local" -path db/migrations/mysql down
	*/

	return db
}

func ConnectToMongoDb() *mongo.Client {
	dbHost := os.Getenv("DB_MONGO_HOST")
	dbPort := os.Getenv("DB_MONGO_PORT")

	uri := fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " +
			"www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	return client
}
