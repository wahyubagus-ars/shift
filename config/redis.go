package config

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

func ConnectToRedis() *redis.Client {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDb := os.Getenv("REDIS_DB")

	redisDbInt, err := strconv.Atoi(redisDb)
	if err != nil {
		log.Error("Error when convert redis db to integer")
		panic(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: redisPassword,
		DB:       redisDbInt,
	})

	return rdb
}
