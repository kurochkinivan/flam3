package fractal_generator_test

import (
	"context"
	"log/slog"
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/suite"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/application/fractal_generator"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/coefficients"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/entities"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/fractal_config"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/variation"
)

type GenerateFractalSuite struct {
	suite.Suite

	ctx       context.Context
	generator *fractal_generator.FractalGenerator
}

func TestGenerateFractalSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(GenerateFractalSuite))
}

func (suite *GenerateFractalSuite) SetupTest() {
	log := slog.New(slog.DiscardHandler)
	slog.SetDefault(log)

	suite.ctx = context.Background()
	suite.generator = fractal_generator.New()
}

func (suite *GenerateFractalSuite) TestGenerateFractal_OneThread() {
	cfg, err := fractalConfig(10, 10, 1)
	suite.Require().NoError(err)

	pixels := suite.generator.GenerateFractal(suite.ctx, cfg)
	suite.Require().NotNil(pixels)
}

func (suite *GenerateFractalSuite) TestGenerateFractal_MultipleThreads() {
	cfg, err := fractalConfig(10, 10, 4)
	suite.Require().NoError(err)

	pixels := suite.generator.GenerateFractal(suite.ctx, cfg)
	suite.Require().NotNil(pixels)
}

func (suite *GenerateFractalSuite) TestGenerateFractal_MultipleThreads_RaceTest() {
	cfg, err := fractalConfig(1000, 1000, 8)
	suite.Require().NoError(err)

	pixels := suite.generator.GenerateFractal(suite.ctx, cfg)
	suite.Require().NotNil(pixels)
}

func fractalConfig(samples, iterations, threads int) (*fractal_config.Config, error) {
	rng := rand.New(rand.NewChaCha8([32]byte{}))

	resolution := entities.NewResolution(100, 100)
	bounds := entities.DefaultMathBounds()

	fName := variation.Cosine
	f, err := variation.Provider(fName)
	if err != nil {
		return nil, err
	}

	return &fractal_config.Config{
		Viewport:      entities.NewViewport(resolution, bounds),
		Coeffs:        []coefficients.Coefficients{coefficients.NewRandom(rng)},
		Variations:    []entities.WeightedVariation{entities.NewWeightedVariation(f, 1.0)},
		Rand:          rand.New(rand.NewChaCha8([32]byte{})),
		Samples:       samples,
		Iterations:    iterations,
		Threads:       threads,
		SymmetryLevel: 1,
	}, nil
}

const (
	benchSamples    = 1_000
	benchIterations = 1_000
	benchWidth      = 100
	benchHeight     = 100
)

func BenchmarkGenerateFractal_1Thread(b *testing.B) {
	benchmarkGenerateFractal(b, 1)
}

func BenchmarkGenerateFractal_2Threads(b *testing.B) {
	benchmarkGenerateFractal(b, 2)
}

func BenchmarkGenerateFractal_4Threads(b *testing.B) {
	benchmarkGenerateFractal(b, 4)
}

func BenchmarkGenerateFractal_8Threads(b *testing.B) {
	benchmarkGenerateFractal(b, 8)
}

func benchmarkGenerateFractal(b *testing.B, threads int) {
	slog.SetDefault(slog.New(slog.DiscardHandler))
	generator := fractal_generator.New()

	cfg, err := benchFractalConfig(benchSamples, benchIterations, threads)
	if err != nil {
		b.Fatalf("failed to create config: %v", err)
	}

	for b.Loop() {
		pixels := generator.GenerateFractal(context.TODO(), cfg)
		if pixels == nil {
			b.Fatal("generated pixels are nil")
		}
	}
}

func benchFractalConfig(samples, iterations, threads int) (*fractal_config.Config, error) {
	return benchFractalConfigCustomSize(samples, iterations, threads, benchWidth, benchHeight)
}

func benchFractalConfigCustomSize(samples, iterations, threads, width, height int) (*fractal_config.Config, error) {
	rng := rand.New(rand.NewChaCha8([32]byte{}))

	resolution := entities.NewResolution(width, height)
	bounds := entities.DefaultMathBounds()

	// Используем несколько вариаций для более реалистичного теста
	variations := []entities.WeightedVariation{}
	variationNames := []variation.Name{
		variation.Sinusoidal,
		variation.Spherical,
		variation.Swirl,
		variation.Horseshoe,
	}

	for _, vName := range variationNames {
		f, err := variation.Provider(vName)
		if err != nil {
			return nil, err
		}
		variations = append(variations, entities.NewWeightedVariation(f, 0.25))
	}

	// Несколько коэффициентов для большей сложности
	coeffs := []coefficients.Coefficients{
		coefficients.NewRandom(rng),
		coefficients.NewRandom(rng),
		coefficients.NewRandom(rng),
	}

	return &fractal_config.Config{
		Viewport:      entities.NewViewport(resolution, bounds),
		Coeffs:        coeffs,
		Variations:    variations,
		Rand:          rand.New(rand.NewChaCha8([32]byte{})),
		Samples:       samples,
		Iterations:    iterations,
		Threads:       threads,
		SymmetryLevel: 1,
	}, nil
}
