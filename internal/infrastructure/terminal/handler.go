package terminal

import (
	"context"
	"fmt"
	"image"

	"github.com/urfave/cli/v3"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/fractal_config"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/pixels"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/infrastructure/input_config"
)

const (
	ExitCodeUnexpectedError = 1
	ExitCodeInvalidInput    = 2
)

type Handler struct {
	version          string
	fractalGenerator FractalGenerator
	imageSaver       ImageSaver
}

type FractalGenerator interface {
	GenerateFractal(ctx context.Context, cfg *fractal_config.Config) *pixels.Pixels
}

type ImageSaver interface {
	SaveImage(img image.Image, path string) error
}

func New(version string, fractalGenerator FractalGenerator, imageSaver ImageSaver) *Handler {
	return &Handler{
		version:          version,
		fractalGenerator: fractalGenerator,
		imageSaver:       imageSaver,
	}
}

func (h *Handler) FractalFlameCommand(ctx context.Context, cmd *cli.Command) error {
	inputCfg, err := h.loadInputConfig(cmd)
	if err != nil {
		return cli.Exit(err, ExitCodeInvalidInput)
	}

	fractalCfg, err := fractal_config.New(inputCfg, fractal_config.DefaultSamples)
	if err != nil {
		return cli.Exit(err, ExitCodeInvalidInput)
	}

	pixels := h.fractalGenerator.GenerateFractal(ctx, fractalCfg)

	if err := h.imageSaver.SaveImage(pixels.Image(), inputCfg.Output); err != nil {
		return cli.Exit(err, ExitCodeUnexpectedError)
	}

	return nil
}

func (h *Handler) loadInputConfig(cmd *cli.Command) (*input_config.Config, error) {
	var cfg *input_config.Config

	configPath := cmd.String("config")
	switch configPath {
	case "":
		width := cmd.Int("width")
		height := cmd.Int("height")
		seed := cmd.Float64("seed")
		iterations := cmd.Int("iteration-count")
		output := cmd.String("output-path")
		threads := cmd.Int("threads")
		affineParams := cmd.String("affine-params")
		functions := cmd.StringSlice("functions")
		gammaCorrection := cmd.Bool("gamma-correction")
		gamma := cmd.Float64("gamma")
		symmetry := cmd.Int("symmetry-level")

		size := input_config.NewSize(width, height)
		parsedAffineParams, err := h.parseAffineParamsSlice(affineParams)
		if err != nil {
			return nil, fmt.Errorf("failed to parse affine params: %w", err)
		}
		parsedAffineFunctions, err := h.parseFunctionsSlice(functions)
		if err != nil {
			return nil, fmt.Errorf("failed to parse functions: %w", err)
		}

		cfg = input_config.New(
			size, seed, iterations, output,
			threads, parsedAffineParams, parsedAffineFunctions,
			gammaCorrection, gamma, symmetry,
		)
	default:
		var err error
		cfg, err = input_config.ReadConfig(configPath)
		if err != nil {
			return nil, fmt.Errorf("failed to read config: %w", err)
		}
	}

	cfg.ApplyDefaults()

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate config: %w", err)
	}

	return cfg, nil
}
