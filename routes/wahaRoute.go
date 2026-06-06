package routes

import (
	"waha-bot/controllers"

	"github.com/gofiber/fiber/v3"
)

func SetupWahaRoutes(app *fiber.App) {
	app.Post("/webhook/waha", controllers.WebhookHandler)
}