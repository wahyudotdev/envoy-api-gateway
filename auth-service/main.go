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
	app.Post("/v1/auth/login", handler.Login())
	log.Fatalln(app.Listen(fmt.Sprintf(":%s", port)))
}
