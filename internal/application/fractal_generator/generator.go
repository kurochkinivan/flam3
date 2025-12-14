package fractal_generator

import (
	"log/slog"
	"sync"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/fractal"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/fractal_config"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/pixels"
)

type FractalGenerator struct {
	log *slog.Logger
}

func New(log *slog.Logger) *FractalGenerator {
	return &FractalGenerator{
		log: log,
	}
}

// GenerateFractal generates a fractal image based on the provided configuration.
// It uses multiple workers to generate the image in parallel, and then merges the results.
// The number of workers, samples per worker, and symmetry level are all configurable.
// The function returns a pointer to the generated image.
func (f *FractalGenerator) GenerateFractal(cfg *fractal_config.Config) *pixels.Pixels {
	fractal := fractal.New(cfg.Rand, cfg.Viewport, cfg.Variations, cfg.Coeffs)

	pixelsChan := make(chan *pixels.Pixels, cfg.Threads)
	wg := new(sync.WaitGroup)

	samplesPerWorker := cfg.Samples / cfg.Threads
	remainder := cfg.Samples % cfg.Threads
	firstWorkerSamples := samplesPerWorker + remainder

	f.log.Info("starting workers",
		slog.Int("workers", cfg.Threads),
		slog.Int("samples_per_worker", samplesPerWorker),
		slog.Int("first_worker_samples", firstWorkerSamples),
	)

	for i := range cfg.Threads {
		workerID := i
		samplesForWorker := samplesPerWorker
		if i == 0 {
			samplesForWorker = firstWorkerSamples
		}

		wg.Go(func() {
			f.log.Debug("worker started",
				slog.Int("worker_id", workerID),
				slog.Int("samples", samplesForWorker),
			)

			pixels := fractal.Generate(samplesForWorker, cfg.Iterations, cfg.SymmetryLevel)

			f.log.Debug("worker completed",
				slog.Int("worker_id", workerID),
			)

			pixelsChan <- pixels
		})
	}

	go func() {
		f.log.Debug("waiting for all workers to complete")
		wg.Wait()
		f.log.Debug("all workers completed, closing channel")
		close(pixelsChan)
	}()

	f.log.Debug("merging pixels from workers")

	var mergedPixels *pixels.Pixels
	for pixels := range pixelsChan {
		if mergedPixels == nil {
			mergedPixels = pixels
			continue
		}

		mergedPixels.Merge(pixels)
	}

	f.log.Debug("all pixels were merged")

	return mergedPixels
}
