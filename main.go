package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"test/utils"
	"test/views"
)

func main() {
	app := fiber.New()

	// i18n language middleware
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("language", "DE")
		return c.Next()
	})

	//app.Use(func(c *fiber.Ctx) error {
	//	c.Locals("translate", translations.MakeTranslateFn(c))
	//	return c.Next()
	//})

	app.Get("/", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		return utils.Render(c, views.Greeter())
	})

	log.Fatal(app.Listen(":3000"))
}
