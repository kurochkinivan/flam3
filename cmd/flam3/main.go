package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/application/fractal_generator"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/infrastructure/image_saver"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/infrastructure/terminal"
)

var version = "dev"

func main() {
	ctx := context.Background()
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(log)

	imageSaver := image_saver.NewPNGSaver()
	fractalGenerator := fractal_generator.New()

	handler := terminal.New(version, fractalGenerator, imageSaver)

	if err := handler.Run(ctx, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "failed to run handler: %s\n", err.Error())
	}
}
