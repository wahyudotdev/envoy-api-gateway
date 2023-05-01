package main

import (
	"embed"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"log"
	"net/http"
	"os"
)

//go:embed docs
var docs embed.FS

func main() {
	app := fiber.New()
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))
	app.Use("/v1", filesystem.New(filesystem.Config{
		Root:   http.FS(docs),
		Browse: false,
		Index:  "index.html",
	}))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatalln(app.Listen(fmt.Sprintf(":%s", port)))
}
