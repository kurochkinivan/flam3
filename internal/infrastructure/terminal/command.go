package terminal

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/infrastructure/input_config"
)

func (h *Handler) Run(ctx context.Context, osArgs []string) error {
	if err := h.app().Run(ctx, osArgs); err != nil {
		return fmt.Errorf("failed to run app: %w", err)
	}

	return nil
}

func (h *Handler) app() *cli.Command {
	return &cli.Command{
		Name:    "flam3",
		Version: h.version,
		Usage:   "Fractal Flame image generator",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "width",
				Aliases: []string{"w"},
				Usage:   fmt.Sprintf("Output image width (int), default: %d", input_config.DefaultWidth),
			},
			&cli.IntFlag{
				Name:    "height",
				Aliases: []string{"h"},
				Usage:   fmt.Sprintf("Output image height (int), default: %d", input_config.DefaultHeight),
			},
			&cli.Float64Flag{
				Name:  "seed",
				Usage: fmt.Sprintf("Random generator seed (float64), default: %.2f", input_config.DefaultSeed),
			},
			&cli.IntFlag{
				Name:    "iteration-count",
				Aliases: []string{"i"},
				Usage:   fmt.Sprintf("Number of generation iterations (int), default: %d", input_config.DefaultIterations),
			},
			&cli.StringFlag{
				Name:    "output-path",
				Aliases: []string{"o"},
				Usage:   fmt.Sprintf("Relative path to output PNG file, default: %s", input_config.DefaultOutputPath),
			},
			&cli.IntFlag{
				Name:    "threads",
				Aliases: []string{"t"},
				Usage:   fmt.Sprintf("Number of worker threads (int), default: %d", input_config.DefaultThreads),
			},
			&cli.StringFlag{
				Name:      "affine-params",
				Aliases:   []string{"ap"},
				Validator: h.validateAffineParams,
				Usage: "Affine transformation configuration in format: " +
					"a,b,c,d,e,f/a,b,c,d,e,f. " +
					"a,e — scale/rotation; b,d — axis mixing; c,f — translations. " +
					"Example: 0.1,0.2,0,0,0.1,0/0.5,0,0,0,0.5,0",
			},
			&cli.StringFlag{
				Name:      "functions",
				Aliases:   []string{"f"},
				Validator: h.validateFunctions,
				Usage: "List of variation functions in format: name:weight,name:weight. " +
					"name — function name (e.g. swirl, horseshoe), " +
					"weight — transformation weight (float64). " +
					"Example: swirl:1.0,horseshoe:0.8",
			},
			&cli.BoolFlag{
				Name:    "gamma-correction",
				Aliases: []string{"g"},
				Usage:   fmt.Sprintf("Enable gamma correction (bool), default: %t", input_config.DefaultGammaCorrection),
			},

			&cli.Float64Flag{
				Name:  "gamma",
				Usage: fmt.Sprintf("Gamma value for correction (float64), default: %.2f", input_config.DefaultGamma),
			},

			&cli.IntFlag{
				Name:    "symmetry-level",
				Aliases: []string{"s"},
				Usage:   fmt.Sprintf("Number of rotational symmetries N >= 1, default: %d", input_config.DefaultSymmetryLevel),
			},
			&cli.StringFlag{
				Name:  "config",
				Usage: "Relative path to optional configuration file. If provided, other flags will be ignored",
			},
		},
		Action: h.FractalFlameCommand,
	}
}

func (h *Handler) validateFunctions(functions string) error {
	_, err := h.parseFunctionsSlice(functions)
	if err != nil {
		return fmt.Errorf("failed to parse functions: %w", err)
	}
	return nil
}

func (h *Handler) validateAffineParams(params string) error {
	_, err := h.parseAffineParamsSlice(params)
	if err != nil {
		return fmt.Errorf("failed to parse affine params: %w", err)
	}

	return nil
}
