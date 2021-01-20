package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hola holita, vecinito")
	})

	log.Fatal(app.Listen(":3000"))

	// // GET /john
	// app.Get("/:name", func(c *fiber.Ctx) error {
	// 	msg := fmt.Sprintf("Hello, %s 👋!", c.Params("name"))
	// 	return c.SendString(msg) // => Hello john 👋!
	// })

	// log.Fatal(app.Listen(":3000"))
}
