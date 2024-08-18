package i18n

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type TranslateFunc func(path string) string

func Translate(c *fiber.Ctx, path string) string {
	language, ok := c.Locals("language").(string)
	if !ok {
		language = "DE"
	}
	return fmt.Sprintf("I translate %s to %s", path, language)
}

func MakeTranslateFn(c *fiber.Ctx) TranslateFunc {
	return func(path string) string {
		return Translate(c, path)
	}
}

// NewMiddleware attaches language to all requests
func NewMiddleware() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.Locals("translate", MakeTranslateFn(c))
		c.Locals("language", "DE")
		return c.Next()
	}
}
