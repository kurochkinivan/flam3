package entities

import (
	"fmt"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/variation"
)

type WeightedVariation struct {
	variation.NamedFunction
	
	Weight float64
}

func NewWeightedVariation(namedFunction variation.NamedFunction, weight float64) WeightedVariation {
	return WeightedVariation{
		NamedFunction: namedFunction,
		Weight:        weight,
	}
}

func (w WeightedVariation) String() string {
	return fmt.Sprintf("%s(%.2f)", w.Name, w.Weight)
}
