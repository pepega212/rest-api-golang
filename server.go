package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"go.mod/configs"
	"go.mod/routes"
)

func main() {
	// loadEnv()
	configs.InitDatabase()

	e := echo.New()
	routes.InitRoute(e)

	// start the server, and log if it fails
	// e.Logger.Fatal(e.Start(":8000"))
	e.Start(getPort())
}

func getPort() string {
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	return ":8000"
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}
