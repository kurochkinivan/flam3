package terminal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/infrastructure/input_config"
)

func TestHandler_parseFunctionsSlice(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected []input_config.WeightedFunction
		hasError bool
	}{
		{
			name:  "happy path",
			input: "linear:1.0,spherical:0.5",
			expected: []input_config.WeightedFunction{
				input_config.NewWeightedFunction("linear", 1.0),
				input_config.NewWeightedFunction("spherical", 0.5),
			},
			hasError: false,
		},
		{
			name:     "empty input",
			input:    "",
			expected: nil,
			hasError: true,
		},
		{
			name:     "invalid format",
			input:    "linear",
			expected: nil,
			hasError: true,
		},
		{
			name:     "invalid weight",
			input:    "linear:abc",
			expected: nil,
			hasError: true,
		},
	}

	h := &Handler{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := h.parseFunctionsSlice(tt.input)
			if tt.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestHandler_parseAffineParamsSlice(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected []input_config.AffineParams
		hasError bool
	}{
		{
			name:  "happy path",
			input: "1.0,2.0,3.0,4.0,5.0,6.0/0.5,1.5,2.5,3.5,4.5,5.5",
			expected: []input_config.AffineParams{
				input_config.NewAffineParams(1.0, 2.0, 3.0, 4.0, 5.0, 6.0),
				input_config.NewAffineParams(0.5, 1.5, 2.5, 3.5, 4.5, 5.5),
			},
			hasError: false,
		},
		{
			name:     "empty input",
			input:    "",
			expected: nil,
			hasError: true,
		},
		{
			name:     "invalid float",
			input:    "1.0,2.0,abc,4.0,5.0,6.0",
			expected: nil,
			hasError: true,
		},
		{
			name:     "wrong number of params",
			input:    "1.0,2.0,3.0,4.0,5.0",
			expected: nil,
			hasError: true,
		},
	}

	h := &Handler{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := h.parseAffineParamsSlice(tt.input)
			if tt.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
