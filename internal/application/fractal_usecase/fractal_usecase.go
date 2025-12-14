package fractal_usecase

import (
	"context"
	"fmt"
	"image"
	"log/slog"
	"time"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/fractal_config"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/pixels"
)

type FractalUsecase struct {
	log        *slog.Logger
	generator  Generator
	imageSaver ImageSaver
}

type Generator interface {
	GenerateFractal(cfg *fractal_config.Config) *pixels.Pixels
}

type ImageSaver interface {
	SaveImage(img image.Image, path string) error
}

func New(log *slog.Logger, generator Generator, imageSaver ImageSaver) *FractalUsecase {
	return &FractalUsecase{
		log:        log,
		generator:  generator,
		imageSaver: imageSaver,
	}
}

func (f *FractalUsecase) Execute(ctx context.Context, cfg *fractal_config.Config) error {
	// 1. Generate fractal
	start := time.Now()
	f.log.InfoContext(ctx, "generating fractal",
		slog.Int("samples", cfg.Samples),
		slog.Int("iterations", cfg.Iterations),
		slog.Int("threads", cfg.Threads),
		slog.Int("symmetry_level", cfg.SymmetryLevel),
		slog.String("coefficients", cfg.Coeffs.String()),
		slog.String("variations", cfg.Variations.String()),
	)

	pixels := f.generator.GenerateFractal(cfg)

	f.log.InfoContext(ctx, "completed fractal generation", slog.Duration("duration", time.Since(start)))

	// 2. Apply gamma factor if necessary
	if cfg.GammaCorrection {
		start = time.Now()
		f.log.InfoContext(ctx, "applying gamma factor", slog.Float64("gamma", cfg.Gamma))

		pixels.ApplyGammaFactor(cfg.Gamma)

		f.log.InfoContext(ctx, "gamma factor was applied", slog.Duration("duration", time.Since(start)))
	} else {
		f.log.InfoContext(ctx, "gamma correction not enabled", slog.Bool("enabled", cfg.GammaCorrection))
	}

	// 3. Convert pixels to image and save it
	f.log.InfoContext(ctx, "saving image", slog.String("output_path", cfg.OutputPath))

	if err := f.imageSaver.SaveImage(pixels.Image(), cfg.OutputPath); err != nil {
		return fmt.Errorf("failed to save image: %w", err)
	}

	f.log.InfoContext(ctx, "image was successfully saved")

	return nil
}
