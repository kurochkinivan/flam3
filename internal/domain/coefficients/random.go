package coefficients

import (
	"image/color"
	"math"
	"math/rand/v2"
)

const (
	minColor   = 64
	maxColor   = 256
	colorRange = maxColor - minColor

	coordMin = -2.0
	coordMax = 2.0
)

func RandomColor(rnd *rand.Rand) color.RGBA {
	return color.RGBA{
		R: uint8(minColor + rnd.IntN(colorRange)),
		G: uint8(minColor + rnd.IntN(colorRange)),
		B: uint8(minColor + rnd.IntN(colorRange)),
		A: 255,
	}
}

func NewRandom(rnd *rand.Rand) Coefficients {
	a, b, d, e := generateABDE(rnd)
	c, f := randFloat(rnd, coordMin, coordMax), randFloat(rnd, coordMin, coordMax)

	return New(a, b, d, e, c, f, RandomColor(rnd))
}

func generateABDE(rnd *rand.Rand) (a, b, d, e float64) {
	for {
		a, d = generateAD(rnd)
		b, e = generateBE(rnd)

		if conditionABDE(a, b, d, e) {
			return
		}
	}
}

func generateAD(rnd *rand.Rand) (a, d float64) {
	for {
		a = rnd.Float64()
		d = randFloat(rnd, a*a, 1)
		if randBool(rnd) {
			d = -d
		}
		if conditionAD(a, d) {
			return
		}
	}
}

func generateBE(rnd *rand.Rand) (b, e float64) {
	for {
		b = rnd.Float64()
		e = randFloat(rnd, b*b, 1)
		if randBool(rnd) {
			e = -e
		}
		if conditionBE(b, e) {
			return
		}
	}
}

func conditionABDE(a, b, d, e float64) bool {
	return a*a+b*b+d*d+e*e < 1+math.Pow(a*e-b*d, 2)
}

func conditionAD(a, d float64) bool {
	return a*a+d*d < 1
}

func conditionBE(b, e float64) bool {
	return b*b+e*e < 1
}

// randFloatN возвращает случайное число в [a,b)
func randFloat(rnd *rand.Rand, a, b float64) float64 {
	return a + rnd.Float64()*(b-a)
}

// randBool возвращает true или false с равной вероятностью
func randBool(rnd *rand.Rand) bool {
	return rnd.IntN(2) == 1
}
