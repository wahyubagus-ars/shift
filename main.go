package main

import (
	"github.com/joho/godotenv"
	"go-shift/cmd/app/routes"
	"go-shift/config"
	"os"
)

func init() {
	godotenv.Load()
}

func main() {
	port := os.Getenv("APP_PORT")

	initialization := config.Init()
	r := routes.Init(initialization)

	r.Run(":" + port)
}
