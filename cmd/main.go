package main

import (
	"geocity/georedis"
	"geocity/handlers"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5173"
	}
	return port
}

func main() {
	geoDB := georedis.NewClientMust(os.Getenv("REDIS_PUBLIC_URL"))

	handler := handlers.Handler{DB: geoDB}
	e := echo.New()

	e.Static("/", "./static")

	handler.AttachRoutes(e)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${status} ${latency_human} ${method} ${path}\n",
	}))

	e.Logger.Fatal(e.Start(":" + getPort()))
}
