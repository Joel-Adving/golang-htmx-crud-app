package main

import (
	"log"

	"github.com/Joel-Adving/testing-stuff/db"
	"github.com/Joel-Adving/testing-stuff/router"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	app := fiber.New(FiberConfig)
	app.Static("/static", "./public", StaticConfig)
	db.Init()
	router.Routes(app)
	log.Fatal(app.Listen(":3000"))
}
