package main

import (
	"github.com/joho/godotenv"
	"go-shift/cmd/app/provider"
	"go-shift/cmd/app/routes"
	"os"
)

func init() {
	godotenv.Load()
}

func main() {
	port := os.Getenv("APP_PORT")

	initialization := provider.Wire()
	r := routes.Init(initialization)

	r.Run(":" + port)
}
