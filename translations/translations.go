package translations

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
