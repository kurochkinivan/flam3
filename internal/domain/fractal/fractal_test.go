package fractal_test

import (
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/coefficients"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/entities"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/fractal"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/pixels"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/variation"
)

func TestGenerate_HappyPath(t *testing.T) {
	rng := rand.New(rand.NewChaCha8([32]byte{}))

	resolution := entities.NewResolution(100, 100)
	mathBounds := entities.DefaultMathBounds()
	viewPort := entities.NewViewport(resolution, mathBounds)

	f, err := variation.Provider(variation.Cosine)
	require.NoError(t, err)

	variations := []entities.WeightedVariation{entities.NewWeightedVariation(f, 1.0)}
	coeffs := []coefficients.Coefficients{coefficients.NewRandom(rng)}

	fr := fractal.New(rng, viewPort, variations, coeffs)

	result := fr.Generate(100, 1000, 1)
	require.NotNil(t, result)

	assert.Equal(t, resolution.Width(), result.Width(), "ширина изображения должна совпадать с resolution")
	assert.Equal(t, resolution.Height(), result.Height(), "высота изображения должна совпадать с resolution")

	touchedPixels := countTouchedPixels(result)
	assert.Positive(t, touchedPixels, "должен быть хотя бы один затронутый пиксель")
}

func countTouchedPixels(pixels *pixels.Pixels) int {
	count := 0

	for y := range pixels.Height() {
		for x := range pixels.Width() {
			if pixels.Pix(x, y).Count > 0 {
				count++
			}
		}
	}

	return count
}
