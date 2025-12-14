package fractal_config

import (
	"fmt"
	"math/rand/v2"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/coefficients"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/entities"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/variation"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/infrastructure/input_config"
)

func mapCoefficients(rnd *rand.Rand, cfgCoeffs []input_config.AffineParams) []coefficients.Coefficients {
	coeffs := make([]coefficients.Coefficients, len(cfgCoeffs))

	for idx, element := range cfgCoeffs {
		coeffs[idx] = coefficients.New(
			element.A,
			element.B,
			element.C,
			element.D,
			element.E,
			element.F,
			coefficients.RandomColor(rnd),
		)
	}

	return coeffs
}

func mapVariations(cfgVariations []input_config.WeightedFunction) ([]entities.WeightedVariation, error) {
	variations := make([]entities.WeightedVariation, len(cfgVariations))

	for idx, element := range cfgVariations {
		vName := variation.VariationName(element.Name)
		namedFunction, err := variation.VariationProvider(vName)
		if err != nil {
			return nil, fmt.Errorf("failed to get function %q: %w", element.Name, err)
		}

		variations[idx] = entities.NewWeightedVariation(namedFunction, element.Weight)
	}

	return variations, nil
}
