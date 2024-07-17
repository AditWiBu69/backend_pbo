package main

import (
	"api/config"
	_ "api/docs"
	"api/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"strings"
)

// @title TES SWAGGER ULBI
// @version 1.0
// @description This is a sample swagger for Fiber

// @contact.name API Support
// @contact.url https://github.com/indrariksa
// @contact.email indra@ulbi.ac.id

// @host pbo-tugasbesar-c553f4142173.herokuapp.com
// @BasePath /
// @schemes https http
func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	router.SetupRoutes(app)

	app.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Join(config.GetAllowedOrigins(), ","),
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	port := "80"

	err := app.Listen(":" + port)
	if err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
