package terminal

import (
	"fmt"
	"strconv"
	"strings"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/infrastructure/input_config"
)

func (h *Handler) parseFunctionsSlice(input string) ([]input_config.WeightedFunction, error) {
	if strings.TrimSpace(input) == "" {
		return nil, fmt.Errorf("functions were not provided")
	}

	blocks := strings.Split(input, ",")
	result := make([]input_config.WeightedFunction, 0, len(blocks))

	for i, block := range blocks {
		fn, err := parseFunction(block)
		if err != nil {
			return nil, fmt.Errorf("block %d: %w", i, err)
		}
		result = append(result, fn)
	}

	return result, nil
}

func parseFunction(block string) (input_config.WeightedFunction, error) {
	parts := strings.Split(block, ":")
	if len(parts) != 2 {
		return input_config.WeightedFunction{}, fmt.Errorf("expected format name:weight")
	}

	name := strings.TrimSpace(parts[0])
	if name == "" {
		return input_config.WeightedFunction{}, fmt.Errorf("function name is empty")
	}

	weight, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
	if err != nil {
		return input_config.WeightedFunction{}, fmt.Errorf("invalid weight: %w", err)
	}

	return input_config.NewWeightedFunction(name, weight), nil
}

func (h *Handler) parseAffineParamsSlice(input string) ([]input_config.AffineParams, error) {
	if strings.TrimSpace(input) == "" {
		return nil, fmt.Errorf("affine params were not provided")
	}

	blocks := strings.Split(input, "/")
	result := make([]input_config.AffineParams, 0, len(blocks))

	for i, block := range blocks {
		parts := strings.Split(block, ",")

		affineParams, err := parseAffineParams(parts)
		if err != nil {
			return nil, fmt.Errorf("block %d: %w", i, err)
		}

		result = append(result, affineParams)
	}

	return result, nil
}

func parseAffineParams(parts []string) (input_config.AffineParams, error) {
	if len(parts) != 6 {
		return input_config.AffineParams{}, fmt.Errorf("expected 6 values, got %d", len(parts))
	}

	vals := make([]float64, 6)
	for i, s := range parts {
		v, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
		if err != nil {
			return input_config.AffineParams{}, fmt.Errorf("invalid float at position %d: %w", i, err)
		}

		vals[i] = v
	}

	return input_config.NewAffineParams(vals[0], vals[1], vals[2], vals[3], vals[4], vals[5]), nil
}
