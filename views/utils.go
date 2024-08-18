package views

import (
	"context"
	"fmt"
	"gofiber-templ-bench/utils/i18n"
	"log/slog"
)

func Translate(ctx context.Context, path string) string {
	slog.Info("translate", "context.Value('translate')", ctx.Value("translate"))
	if translate, ok := ctx.Value("translate").(i18n.TranslateFunc); ok {
		return translate(path)
	}
	return fmt.Sprintf("t:%s", path)
}

func GetLanguage(ctx context.Context) string {
	slog.Info("language", "context.Value('language')", ctx.Value("language"))
	if x, ok := ctx.Value("language").(string); ok {
		return x
	}
	return "no value"
}
