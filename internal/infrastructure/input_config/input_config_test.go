package input_config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/infrastructure/input_config"
)

func TestNewSize(t *testing.T) {
	width, height := 1920, 1080
	size := input_config.NewSize(width, height)

	assert.Equal(t, width, size.Width)
	assert.Equal(t, height, size.Height)
}

func TestNewAffineParams(t *testing.T) {
	a, b, c, d, e, f := 1.0, 2.0, 3.0, 4.0, 5.0, 6.0
	params := input_config.NewAffineParams(a, b, c, d, e, f)

	assert.InDelta(t, a, params.A, 0.01)
	assert.InDelta(t, b, params.B, 0.01)
	assert.InDelta(t, c, params.C, 0.01)
	assert.InDelta(t, d, params.D, 0.01)
	assert.InDelta(t, e, params.E, 0.01)
	assert.InDelta(t, f, params.F, 0.01)
}

func TestNewWeightedFunction(t *testing.T) {
	name := "linear"
	weight := 1.0
	wf := input_config.NewWeightedFunction(name, weight)

	assert.Equal(t, name, wf.Name)
	assert.InDelta(t, weight, wf.Weight, 0.01)
}

func TestNewConfig(t *testing.T) {
	size := input_config.NewSize(1920, 1080)
	seed := 5.0
	iterations := 2500
	output := "fractal.png"
	threads := 1
	affineParams := []input_config.AffineParams{
		input_config.NewAffineParams(1, 2, 3, 4, 5, 6),
	}
	weightedFunctions := []input_config.WeightedFunction{
		input_config.NewWeightedFunction("linear", 1.0),
	}
	gammaCorrection := false
	gamma := 2.2
	symmetryLevel := 1

	cfg := input_config.New(
		size, seed, iterations, output, threads,
		affineParams, weightedFunctions,
		gammaCorrection, gamma, symmetryLevel,
	)

	assert.Equal(t, size, cfg.Size)
	assert.InDelta(t, seed, cfg.Seed, 0.01)
	assert.Equal(t, iterations, cfg.Iterations)
	assert.Equal(t, output, cfg.Output)
	assert.Equal(t, threads, cfg.Threads)
	assert.Equal(t, affineParams, cfg.AffineParams)
	assert.Equal(t, weightedFunctions, cfg.WeightedFunctions)
	assert.Equal(t, gammaCorrection, cfg.GammaCorrection)
	assert.InDelta(t, gamma, cfg.Gamma, 0.01)
	assert.Equal(t, symmetryLevel, cfg.SymmetryLevel)
}

func TestReadConfig(t *testing.T) {
	tests := []struct {
		name        string
		jsonContent string
		wantErr     bool
		expected    *input_config.Config
	}{
		{
			name: "valid config",
			jsonContent: `{
				"size": {"width": 1920, "height": 1080},
				"seed": 5.0,
				"iteration_count": 2500,
				"output_path": "fractal.png",
				"threads": 1,
				"affine_params": [{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5, "F": 6}],
				"functions": [{"name": "linear", "weight": 1.0}],
				"gamma_correction": false,
				"gamma": 2.2,
				"symmetry_level": 1
			}`,
			wantErr: false,
			expected: input_config.New(
				input_config.NewSize(1920, 1080), 5.0, 2500, "fractal.png", 1,
				[]input_config.AffineParams{input_config.NewAffineParams(1, 2, 3, 4, 5, 6)},
				[]input_config.WeightedFunction{input_config.NewWeightedFunction("linear", 1.0)},
				false, 2.2, 1,
			),
		},
		{
			name:        "invalid json",
			jsonContent: `{invalid}`,
			wantErr:     true,
		},
		{
			name:        "nonexistent file",
			jsonContent: "",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var filePath string
			if tt.jsonContent != "" {
				tempDir := t.TempDir()
				filePath = filepath.Join(tempDir, "config.json")
				err := os.WriteFile(filePath, []byte(tt.jsonContent), 0644)
				require.NoError(t, err)
			} else {
				filePath = "nonexistent.json"
			}

			cfg, err := input_config.ReadConfig(filePath)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, cfg)
			}
		})
	}
}

func TestConfigApplyDefaults(t *testing.T) {
	tests := []struct {
		name     string
		initial  *input_config.Config
		expected *input_config.Config
	}{
		{
			name: "all defaults",
			initial: input_config.New(
				input_config.NewSize(0, 0), 0, 0, "", 0,
				nil, nil, false, 0, 0,
			),
			expected: input_config.New(
				input_config.NewSize(1920, 1080), 5.0, 2500, "fractal.png", 1,
				nil, nil, false, 2.2, 1,
			),
		},
		{
			name: "partial defaults",
			initial: input_config.New(
				input_config.NewSize(800, 600), 10.0, 1000, "output.png", 2,
				nil, nil, false, 1.5, 2,
			),
			expected: input_config.New(
				input_config.NewSize(800, 600), 10.0, 1000, "output.png", 2,
				nil, nil, false, 1.5, 2,
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.initial.ApplyDefaults()
			assert.Equal(t, tt.expected, tt.initial)
		})
	}
}

func TestConfigValidate(t *testing.T) {
	tests := []struct {
		name    string
		config  *input_config.Config
		wantErr bool
	}{
		{
			name: "valid config",
			config: input_config.New(
				input_config.NewSize(1920, 1080), 5.0, 2500, "fractal.png", 1,
				[]input_config.AffineParams{input_config.NewAffineParams(1, 2, 3, 4, 5, 6)},
				[]input_config.WeightedFunction{input_config.NewWeightedFunction("linear", 1.0)},
				false, 2.2, 1,
			),
			wantErr: false,
		},
		{
			name: "invalid width",
			config: input_config.New(
				input_config.NewSize(0, 1080), 5.0, 2500, "fractal.png", 1,
				[]input_config.AffineParams{input_config.NewAffineParams(1, 2, 3, 4, 5, 6)},
				[]input_config.WeightedFunction{input_config.NewWeightedFunction("linear", 1.0)},
				false, 2.2, 1,
			),
			wantErr: true,
		},
		{
			name: "invalid height",
			config: input_config.New(
				input_config.NewSize(1920, -1), 5.0, 2500, "fractal.png", 1,
				[]input_config.AffineParams{input_config.NewAffineParams(1, 2, 3, 4, 5, 6)},
				[]input_config.WeightedFunction{input_config.NewWeightedFunction("linear", 1.0)},
				false, 2.2, 1,
			),
			wantErr: true,
		},
		{
			name: "invalid iterations",
			config: input_config.New(
				input_config.NewSize(1920, 1080), 5.0, 0, "fractal.png", 1,
				[]input_config.AffineParams{input_config.NewAffineParams(1, 2, 3, 4, 5, 6)},
				[]input_config.WeightedFunction{input_config.NewWeightedFunction("linear", 1.0)},
				false, 2.2, 1,
			),
			wantErr: true,
		},
		{
			name: "no functions",
			config: input_config.New(
				input_config.NewSize(1920, 1080), 5.0, 2500, "fractal.png", 1,
				[]input_config.AffineParams{input_config.NewAffineParams(1, 2, 3, 4, 5, 6)},
				nil,
				false, 2.2, 1,
			),
			wantErr: true,
		},
		{
			name: "invalid function name",
			config: input_config.New(
				input_config.NewSize(1920, 1080), 5.0, 2500, "fractal.png", 1,
				[]input_config.AffineParams{input_config.NewAffineParams(1, 2, 3, 4, 5, 6)},
				[]input_config.WeightedFunction{input_config.NewWeightedFunction("invalid", 1.0)},
				false, 2.2, 1,
			),
			wantErr: true,
		},
		{
			name: "negative weight",
			config: input_config.New(
				input_config.NewSize(1920, 1080), 5.0, 2500, "fractal.png", 1,
				[]input_config.AffineParams{input_config.NewAffineParams(1, 2, 3, 4, 5, 6)},
				[]input_config.WeightedFunction{input_config.NewWeightedFunction("linear", -1.0)},
				false, 2.2, 1,
			),
			wantErr: true,
		},
		{
			name: "no affine params",
			config: input_config.New(
				input_config.NewSize(1920, 1080), 5.0, 2500, "fractal.png", 1,
				nil,
				[]input_config.WeightedFunction{input_config.NewWeightedFunction("linear", 1.0)},
				false, 2.2, 1,
			),
			wantErr: true,
		},
		{
			name: "gamma correction with invalid gamma",
			config: input_config.New(
				input_config.NewSize(1920, 1080), 5.0, 2500, "fractal.png", 1,
				[]input_config.AffineParams{input_config.NewAffineParams(1, 2, 3, 4, 5, 6)},
				[]input_config.WeightedFunction{input_config.NewWeightedFunction("linear", 1.0)},
				true, 0, 1,
			),
			wantErr: true,
		},
		{
			name: "invalid threads",
			config: input_config.New(
				input_config.NewSize(1920, 1080), 5.0, 2500, "fractal.png", 0,
				[]input_config.AffineParams{input_config.NewAffineParams(1, 2, 3, 4, 5, 6)},
				[]input_config.WeightedFunction{input_config.NewWeightedFunction("linear", 1.0)},
				false, 2.2, 1,
			),
			wantErr: true,
		},
		{
			name: "invalid symmetry level",
			config: input_config.New(
				input_config.NewSize(1920, 1080), 5.0, 2500, "fractal.png", 1,
				[]input_config.AffineParams{input_config.NewAffineParams(1, 2, 3, 4, 5, 6)},
				[]input_config.WeightedFunction{input_config.NewWeightedFunction("linear", 1.0)},
				false, 2.2, -1,
			),
			wantErr: true,
		},
		{
			name: "invalid output path",
			config: input_config.New(
				input_config.NewSize(1920, 1080), 5.0, 2500, "/nonexistent/dir/file.png", 1,
				[]input_config.AffineParams{input_config.NewAffineParams(1, 2, 3, 4, 5, 6)},
				[]input_config.WeightedFunction{input_config.NewWeightedFunction("linear", 1.0)},
				false, 2.2, 1,
			),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateGreaterThanZero(t *testing.T) {
	tests := []struct {
		name    string
		value   any
		wantErr bool
	}{
		{"int positive", 1, false},
		{"int zero", 0, true},
		{"int negative", -1, true},
		{"float positive", 1.0, false},
		{"float zero", 0.0, true},
		{"float negative", -1.0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			switch v := tt.value.(type) {
			case int:
				err = input_config.ValidateGreaterThanZero(v)
			case float64:
				err = input_config.ValidateGreaterThanZero(v)
			}
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidatePath(t *testing.T) {
	tempDir := t.TempDir()
	validDir := filepath.Join(tempDir, "subdir")
	require.NoError(t, os.MkdirAll(validDir, 0755))

	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{"valid path", filepath.Join(validDir, "file.png"), false},
		{"nonexistent dir", "/nonexistent/dir/file.png", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := input_config.ValidatePath(tt.path)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
