package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"math/rand/v2"
	"os"
	"path/filepath"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/application/fractal_generator"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/coefficients"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/entities"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/fractal_config"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/variation"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/infrastructure/image_saver"
)

const (
	samples    int = 1e6
	iterations int = 1e2

	width  = 1920
	height = 1080

	threads = 12

	gammaCorrection = true
	gamma           = 2.2
	symmetry        = 1
)

func main() {
	var numberOfImages int
	var outputDir string
	var seed1 uint64
	var seed2 uint64
	var useSeed bool

	flag.IntVar(&numberOfImages, "number", 1, "number of images to create")
	flag.StringVar(&outputDir, "output-dir", "", "directory for output files")
	flag.Uint64Var(&seed1, "seed1", 0, "first seed value for reproducibility")
	flag.Uint64Var(&seed2, "seed2", 0, "second seed value for reproducibility")
	flag.BoolVar(&useSeed, "use-seed", false, "use provided seeds instead of random")
	flag.Parse()

	// Если указаны оба seed, автоматически включаем use-seed
	if seed1 != 0 || seed2 != 0 {
		useSeed = true
	}

	f, err := os.OpenFile("output/logs.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	defer f.Close()

	log := slog.New(slog.NewTextHandler(f, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(log)

	fractalGenerator := fractal_generator.New()
	imageSaver := image_saver.NewPNGSaver()

	ctx := context.Background()
	for idx := range numberOfImages {
		var s1, s2 uint64
		var rng *rand.Rand

		if useSeed {
			// Воспроизводимый режим с указанными seed
			s1, s2 = seed1, seed2
			rng = rand.New(rand.NewPCG(s1, s2))
		} else {
			// Случайный режим
			s1, s2 = rand.Uint64(), rand.Uint64()
			rng = rand.New(rand.NewPCG(s1, s2))
		}

		outputFileName := fmt.Sprintf("fractal_%d_%d.png", s1, s2)
		outputPath := filepath.Join(outputDir, outputFileName)

		cfg := RandomConfig(rng, outputPath)
		log.InfoContext(ctx, "start fractal generation",
			slog.Int("image_index", idx),
			slog.Uint64("seed_1", s1),
			slog.Uint64("seed_2", s2),
			slog.Int("samples", cfg.Samples),
			slog.Int("iterations", cfg.Iterations),
			slog.Int("width", cfg.Viewport.Width()),
			slog.Int("height", cfg.Viewport.Height()),
			slog.Bool("gamma_correction", cfg.GammaCorrection),
			slog.Float64("gamma", cfg.Gamma),
			slog.Int("symmetry", cfg.SymmetryLevel),
			slog.String("affine_params", cfg.Coeffs.String()),
			slog.String("functions", cfg.Variations.String()),
			slog.String("output_path", cfg.OutputPath),
			slog.Int("threads", cfg.Threads),
		)

		pixels := fractalGenerator.GenerateFractal(ctx, cfg)

		err := imageSaver.SaveImage(pixels.Image(), cfg.OutputPath)
		if err != nil {
			log.ErrorContext(ctx, "failed to execute fractal generation",
				slog.Int("image_index", idx),
				slog.String("err", err.Error()),
			)
			fmt.Fprintf(os.Stderr, "failed to execute fractal generation: %v\n", err)
		}
	}
}

func RandomConfig(rng *rand.Rand, outputPath string) *fractal_config.Config {
	const maxCoeffs = 10
	n := rng.IntN(maxCoeffs) + 1
	coeffs := make([]coefficients.Coefficients, n)
	for i := range n {
		coeffs[i] = coefficients.NewRandom(rng)
	}

	const maxFuncs = 15
	n = rng.IntN(maxFuncs) + 1
	variations := make([]entities.WeightedVariation, n)
	for i := range n {
		randomIdx := rng.IntN(len(allVariations))
		namedVariation := allVariations[randomIdx]
		variations[i] = entities.NewWeightedVariation(namedVariation, rng.Float64())
	}

	return &fractal_config.Config{
		Viewport: entities.NewViewport(
			entities.NewResolution(width, height),
			entities.DefaultMathBounds(),
		),
		Variations:      variations,
		Coeffs:          coeffs,
		Samples:         samples,
		Iterations:      iterations,
		GammaCorrection: gammaCorrection,
		Gamma:           gamma,
		SymmetryLevel:   symmetry,
		Threads:         threads,
		Rand:            rng,
		OutputPath:      outputPath,
	}
}

var allVariations = []variation.NamedFunction{
	MustProvideVariation(variation.Cosine),
	MustProvideVariation(variation.Linear),
	MustProvideVariation(variation.Sinusoidal),
	MustProvideVariation(variation.Spherical),
	MustProvideVariation(variation.Swirl),
	MustProvideVariation(variation.Horseshoe),
	MustProvideVariation(variation.Polar),
	MustProvideVariation(variation.Handkerchief),
	MustProvideVariation(variation.Heart),
	MustProvideVariation(variation.Disk),
	MustProvideVariation(variation.Spiral),
	MustProvideVariation(variation.Hyperbolic),
	MustProvideVariation(variation.Diamond),
	MustProvideVariation(variation.Ex),
	MustProvideVariation(variation.Bent),
	MustProvideVariation(variation.Fisheye),
	MustProvideVariation(variation.Eyefish),
	MustProvideVariation(variation.Bubble),
	MustProvideVariation(variation.Cylinder),
	MustProvideVariation(variation.Tangent),
	MustProvideVariation(variation.Cross),
	MustProvideVariation(variation.Exponential),
	MustProvideVariation(variation.Power),
	MustProvideVariation(variation.Cosine),
}

func MustProvideVariation(name variation.Name) variation.NamedFunction {
	fn, err := variation.Provider(name)
	if err != nil {
		panic(err)
	}
	return fn
}
