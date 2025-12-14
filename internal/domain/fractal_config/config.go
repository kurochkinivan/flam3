package fractal_config

import (
	"fmt"
	"math"
	"math/rand/v2"
	"strings"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/coefficients"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/entities"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/infrastructure/input_config"
)

const (
	DefaultSamples int = 1_000
)

type Config struct {
	Viewport        entities.Viewport
	Coeffs          Coefficients
	Variations      Variations
	Rand            *rand.Rand
	Samples         int
	Iterations      int
	OutputPath      string
	Threads         int
	GammaCorrection bool
	Gamma           float64
	SymmetryLevel   int
}

type Variations []entities.WeightedVariation

func (v Variations) String() string {
	parts := make([]string, len(v))

	for i, variation := range v {
		parts[i] = variation.String()
	}

	return "[" + strings.Join(parts, ", ") + "]"
}

type Coefficients []coefficients.Coefficients

func (c Coefficients) String() string {
	parts := make([]string, len(c))

	for i, coeff := range c {
		parts[i] = coeff.String()
	}

	return "[" + strings.Join(parts, ", ") + "]"
}

func New(cfg *input_config.Config, samples int) (*Config, error) {
	seed := math.Float64bits(cfg.Seed)
	rand := rand.New(rand.NewPCG(seed, seed))

	resolution := entities.NewResolution(cfg.Size.Width, cfg.Size.Height)
	mathBounds := entities.NewMathBounds(entities.DefaultXMin, entities.DefaultXMax, entities.DefaultYMin, entities.DefaultYMax)

	coeffs := mapCoefficients(rand, cfg.AffineParams)
	variations, err := mapVariations(cfg.WeightedFunctions)
	if err != nil {
		return nil, fmt.Errorf("failed to map variations: %w", err)
	}

	return &Config{
		Viewport:        entities.NewViewport(resolution, mathBounds),
		Coeffs:          coeffs,
		Variations:      variations,
		Rand:            rand,
		Samples:         samples,
		Iterations:      cfg.Iterations,
		OutputPath:      cfg.Output,
		Threads:         cfg.Threads,
		GammaCorrection: cfg.GammaCorrection,
		Gamma:           cfg.Gamma,
		SymmetryLevel:   cfg.SymmetryLevel,
	}, nil
}
