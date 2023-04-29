package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"os"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	app.Use(logger.New())

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	handler := NewAuthHandler()
	group := app.Group("/v1/auth")
	{
		group.Post("/login", handler.Login())
		group.Get("/verify", handler.VerifyToken())
	}
	log.Fatalln(app.Listen(fmt.Sprintf(":%s", port)))
}
