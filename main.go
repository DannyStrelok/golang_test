package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hola holita, vecinito")
	})

	port := os.Getenv("PORT")

	if os.Getenv("PORT") == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))

}
