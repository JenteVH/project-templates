package main

import (
	"fmt"
	"log"

	"go-template/internal/config"
	"go-template/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.Load()

	app := fiber.New(fiber.Config{
		AppName: "Go Template API",
	})

	routes.Setup(app)

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Server starting on %s (environment: %s)", addr, cfg.Env)
	log.Fatal(app.Listen(addr))
}
