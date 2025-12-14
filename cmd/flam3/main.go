package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/application/fractal_generator"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/application/fractal_usecase"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/infrastructure/image_utils"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/infrastructure/terminal"
)

var version = "dev"

func main() {
	ctx := context.Background()
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	imageSaver := image_utils.NewImageSaver()
	generator := fractal_generator.New(log)
	fractalUseCase := fractal_usecase.New(log, generator, imageSaver)

	handler := terminal.New(version, fractalUseCase)

	if err := handler.Run(ctx, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "failed to run handler: %s\n", err.Error())
	}
}
