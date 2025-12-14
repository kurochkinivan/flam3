package coefficients_test

import (
	"fmt"
	"image/color"
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/coefficients"
)

func TestNew(t *testing.T) {
	a, b, d, e := 1.0, 2.0, 3.0, 4.0
	c, f := 5.0, 6.0
	testColor := color.RGBA{R: 100, G: 150, B: 200, A: 255}

	coeff := coefficients.New(a, b, d, e, c, f, testColor)

	assert.Equal(t, a, coeff.A)
	assert.Equal(t, b, coeff.B)
	assert.Equal(t, d, coeff.D)
	assert.Equal(t, e, coeff.E)
	assert.Equal(t, c, coeff.C)
	assert.Equal(t, f, coeff.F)
	assert.Equal(t, testColor, coeff.Color)
}

func TestCoefficientsString(t *testing.T) {
	a, b, d, e := 1.5, -2.3, 0.7, -1.8
	c, f := 3.2, -0.5
	testColor := color.RGBA{R: 255, G: 128, B: 64, A: 255}

	coeff := coefficients.New(a, b, d, e, c, f, testColor)

	str := coeff.String()
	expected := fmt.Sprintf(
		"Coefficients{A:%.3f, B:%.3f, C:%.3f, D:%.3f, E:%.3f, F:%.3f, Color:RGBA(%d,%d,%d,%d)}",
		a, b, c, d, e, f,
		testColor.R, testColor.G, testColor.B, testColor.A,
	)

	assert.Equal(t, expected, str)
}

func TestRandomColor(t *testing.T) {
	N := 100
	seed := uint64(42)
	rnd := rand.New(rand.NewPCG(seed, seed))

	for range N {
		c := coefficients.RandomColor(rnd)
		assert.Equal(t, uint8(255), c.A)
		assert.GreaterOrEqual(t, c.R, uint8(64))
		assert.GreaterOrEqual(t, c.G, uint8(64))
		assert.GreaterOrEqual(t, c.B, uint8(64))
	}
}

func TestNewRandom(t *testing.T) {
	N := 100
	seed := uint64(42)
	rnd := rand.New(rand.NewPCG(seed, seed))

	for i := range N {
		coeff := coefficients.NewRandom(rnd)

		if coeff.A*coeff.A+coeff.D*coeff.D >= 1 {
			t.Errorf("iteration %d: A^2 + D^2 >= 1 (A=%.4f, D=%.4f)", i, coeff.A, coeff.D)
		}
		if coeff.B*coeff.B+coeff.E*coeff.E >= 1 {
			t.Errorf("iteration %d: B^2 + E^2 >= 1 (B=%.4f, E=%.4f)", i, coeff.B, coeff.E)
		}
		if coeff.A*coeff.A+coeff.B*coeff.B+coeff.D*coeff.D+coeff.E*coeff.E >= 1+(coeff.A*coeff.E-coeff.B*coeff.D)*(coeff.A*coeff.E-coeff.B*coeff.D) {
			t.Errorf("iteration %d: A^2+B^2+D^2+E^2 >= 1+(AE-BD)^2 (A=%.4f, B=%.4f, D=%.4f, E=%.4f)", i, coeff.A, coeff.B, coeff.D, coeff.E)
		}
	}
}
