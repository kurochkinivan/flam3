package fractal_config_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/fractal_config"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/infrastructure/input_config"
)

func TestNew_HappyPath(t *testing.T) {
	in := &input_config.Config{
		Size: input_config.Size{
			Width:  800,
			Height: 600,
		},
		Seed:          42,
		Iterations:    1000,
		Output:        "out.png",
		Threads:       4,
		Gamma:         2.2,
		SymmetryLevel: 3,
		AffineParams: []input_config.AffineParams{
			{A: 1, B: 0, C: 0, D: 1, E: 0, F: 0},
		},
		WeightedFunctions: []input_config.WeightedFunction{
			{Name: "linear", Weight: 1},
		},
	}

	cfg, err := fractal_config.New(in, fractal_config.DefaultSamples)

	require.NoError(t, err)
	require.NotNil(t, cfg)

	// публичный контракт
	require.Equal(t, fractal_config.DefaultSamples, cfg.Samples)
	require.Equal(t, in.Iterations, cfg.Iterations)
	require.Equal(t, in.Output, cfg.OutputPath)
	require.Equal(t, in.Threads, cfg.Threads)
	require.InDelta(t, in.Gamma, cfg.Gamma, 0.01)
	require.Equal(t, in.SymmetryLevel, cfg.SymmetryLevel)

	// наблюдаемые эффекты
	require.Len(t, cfg.Coeffs, len(in.AffineParams))
	require.Len(t, cfg.Variations, len(in.WeightedFunctions))
	require.NotNil(t, cfg.Rand)
}

func TestNew_UnknownVariation(t *testing.T) {
	in := &input_config.Config{
		Size: input_config.Size{
			Width:  800,
			Height: 600,
		},
		Seed:       1,
		Iterations: 10,
		Output:     "out.png",
		Threads:    1,
		AffineParams: []input_config.AffineParams{
			{A: 1, B: 0, C: 0, D: 1, E: 0, F: 0},
		},
		WeightedFunctions: []input_config.WeightedFunction{
			{Name: "definitely_not_exists", Weight: 1},
		},
	}

	cfg, err := fractal_config.New(in, 100)

	require.Error(t, err)
	require.Nil(t, cfg)
}

func TestNew_EmptyAffineAndVariations(t *testing.T) {
	in := &input_config.Config{
		Size:       input_config.Size{Width: 10, Height: 10},
		Seed:       1,
		Iterations: 1,
		Output:     "out",
		Threads:    1,
	}

	cfg, err := fractal_config.New(in, 1)

	require.NoError(t, err)
	require.NotNil(t, cfg)
	require.Empty(t, cfg.Coeffs)
	require.Empty(t, cfg.Variations)
}
