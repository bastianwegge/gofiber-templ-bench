package main

import (
	"context"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"log"
	. "test/views"
)

func main() {
	app := fiber.New()

	app.Get("/with", func(c *fiber.Ctx) error {
		c.Set("X", "Y")
		return Render(c, Greeter())
	})
	app.Get("/without", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		c.Set("X", "Y")
		ctx := context.WithValue(c.Context(), "X", "Y")

		return Greeter().Render(ctx, c)
	})

	log.Fatal(app.Listen(":3000"))
}

func Render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component, options...)
	return adaptor.HTTPHandler(componentHandler)(c)
}
