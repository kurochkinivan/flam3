package input_config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Size              Size               `json:"size"`
	Seed              float64            `json:"seed"`
	Iterations        int                `json:"iteration_count"`
	Output            string             `json:"output_path"`
	Threads           int                `json:"threads"`
	AffineParams      []AffineParams     `json:"affine_params"`
	WeightedFunctions []WeightedFunction `json:"functions"`
	GammaCorrection   bool               `json:"gamma_correction"`
	Gamma             float64            `json:"gamma"`
	SymmetryLevel     int                `json:"symmetry_level"`
}

type Size struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func NewSize(width, height int) Size {
	return Size{
		Width:  width,
		Height: height,
	}
}

type AffineParams struct {
	A, B, C, D, E, F float64
}

func NewAffineParams(a, b, c, d, e, f float64) AffineParams {
	return AffineParams{
		A: a,
		B: b,
		C: c,
		D: d,
		E: e,
		F: f,
	}
}

type WeightedFunction struct {
	Name   string  `json:"name"`
	Weight float64 `json:"weight"`
}

func NewWeightedFunction(name string, weight float64) WeightedFunction {
	return WeightedFunction{
		Name:   name,
		Weight: weight,
	}
}

func New(
	size Size,
	seed float64,
	iterations int,
	outputPath string,
	threads int,
	affineParams []AffineParams,
	weightedFunctions []WeightedFunction,
	gammaCorrection bool,
	gamma float64,
	symmetryLevel int,
) *Config {
	return &Config{
		Size:              size,
		Seed:              seed,
		Iterations:        iterations,
		Output:            outputPath,
		Threads:           threads,
		AffineParams:      append([]AffineParams(nil), affineParams...),
		WeightedFunctions: append([]WeightedFunction(nil), weightedFunctions...),
		GammaCorrection:   gammaCorrection,
		Gamma:             gamma,
		SymmetryLevel:     symmetryLevel,
	}
}

func ReadConfig(configPath string) (*Config, error) {
	f, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file %q: %w", configPath, err)
	}
	defer f.Close()

	cfg := new(Config)
	if err := cleanenv.ParseJSON(f, cfg); err != nil {
		return nil, fmt.Errorf("failed to parse json: %w", err)
	}

	return cfg, nil
}
