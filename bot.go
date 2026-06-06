package main

import (
	"log"

	"waha-bot/config"
	"waha-bot/routes"
	"github.com/gofiber/fiber/v3"
)

func main() {
	// Load environment variables
	config.LoadConfig()

	app := fiber.New(fiber.Config{
		AppName: "Msytc WAHA Bot",
	})
	routes.SetupWahaRoutes(app)
	log.Fatal(app.Listen(":3001"))
}