package routes

import (
	"go-template/internal/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Setup(app *fiber.App) {
	app.Use(logger.New())

	app.Get("/", handlers.Home)
	app.Get("/health", handlers.HealthCheck)
	app.Get("/api/health", handlers.HealthCheck)
}
