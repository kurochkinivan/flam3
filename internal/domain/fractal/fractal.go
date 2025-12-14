package fractal

import (
	"image/color"
	"math"
	"math/rand/v2"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/coefficients"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/entities"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/pixels"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/variation"
)

type Fractal struct {
	rand         *rand.Rand
	viewPort     entities.Viewport
	variations   weightedVariations
	coefficients []coefficients.Coefficients
}

type weightedVariations struct {
	totalWeight float64
	variations  []entities.WeightedVariation
}

func newWeightedVariations(variations []entities.WeightedVariation) weightedVariations {
	var totalWeight float64
	for _, variation := range variations {
		totalWeight += variation.Weight
	}

	return weightedVariations{
		totalWeight: totalWeight,
		variations:  variations,
	}
}

func New(
	rand *rand.Rand,
	viewPort entities.Viewport,
	variations []entities.WeightedVariation,
	coefficients []coefficients.Coefficients,
) *Fractal {

	return &Fractal{
		rand:         rand,
		viewPort:     viewPort,
		variations:   newWeightedVariations(variations),
		coefficients: coefficients,
	}
}

func (f *Fractal) Generate(samples, iterations int, symmetry int) *pixels.Pixels {
	pixels := pixels.NewPixels(f.viewPort.Resolution)

	for range samples {
		newX := f.randomFloat(f.viewPort.XMin(), f.viewPort.XMax())
		newY := f.randomFloat(f.viewPort.YMin(), f.viewPort.YMax())

		for step := -20; step < iterations; step++ {
			variation := f.randomVariation()
			coeff := f.randomCoefficients()

			x := coeff.A*newX + coeff.B*newY + coeff.C
			y := coeff.D*newX + coeff.E*newY + coeff.F

			newX, newY = variation(x, y)

			if step >= 0 {
				var theta float64

				for range symmetry {
					theta += ((2 * math.Pi) / float64(symmetry))

					rotX := newX*math.Cos(theta) - newY*math.Sin(theta)
					rotY := newX*math.Sin(theta) + newY*math.Cos(theta)

					if f.viewPort.InBoundsX(rotX) && f.viewPort.InBoundsY(rotY) {
						x1 := f.viewPort.XToPixel(rotX)
						y1 := f.viewPort.YToPixel(rotY)

						if f.viewPort.InBoundsPixelX(x1) && f.viewPort.InBoundsPixelY(y1) {
							point := pixels.Pix(x1, y1)

							if point.Count == 0 {
								point.Color = color.RGBA{
									R: coeff.Color.R,
									G: coeff.Color.G,
									B: coeff.Color.B,
									A: 255,
								}
							} else {
								old := point.Color
								point.Color = color.RGBA{
									R: uint8((float64(old.R) + float64(coeff.Color.R)) / 2.0),
									G: uint8((float64(old.G) + float64(coeff.Color.G)) / 2.0),
									B: uint8((float64(old.B) + float64(coeff.Color.B)) / 2.0),
									A: 255,
								}
							}

							point.Count++
						}
					}
				}
			}
		}
	}

	return pixels
}

func (f *Fractal) randomCoefficients() coefficients.Coefficients {
	idx := f.rand.IntN(len(f.coefficients))
	return f.coefficients[idx]
}

func (f *Fractal) randomVariation() variation.F {
	want := f.rand.Float64() * f.variations.totalWeight
	var acc float64

	for _, v := range f.variations.variations {
		acc += v.Weight

		if acc >= want {
			return v.Function
		}
	}

	return f.variations.variations[len(f.variations.variations)-1].Function
}

func (f *Fractal) randomFloat(a, b float64) float64 {
	return a + f.rand.Float64()*(b-a)
}
