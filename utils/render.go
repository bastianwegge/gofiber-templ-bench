package utils

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"test/views/layouts"
)

func Render(c *fiber.Ctx, component templ.Component) error {
	c.Set("Content-Type", "text/html")
	return layouts.MainLayoutView(component).Render(c.Context(), c.Response().BodyWriter())
}

func RenderPartial(c *fiber.Ctx, component templ.Component) error {
	c.Set("Content-Type", "text/html")
	return component.Render(c.Context(), c.Response().BodyWriter())
}
