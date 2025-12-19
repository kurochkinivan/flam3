package input_config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/variation"
)

func (c *Config) Validate() error {
	if err := c.Size.Validate(); err != nil {
		return fmt.Errorf("invalid size: %w", err)
	}

	if err := ValidateGreaterThanZero(c.Iterations); err != nil {
		return fmt.Errorf("invalid iterations: %w", err)
	}

	if err := validateWeightedFunctions(c.WeightedFunctions); err != nil {
		return fmt.Errorf("invalid functions: %w", err)
	}

	if err := validateAffineParams(c.AffineParams); err != nil {
		return fmt.Errorf("invalid affine params: %w", err)
	}

	if c.GammaCorrection {
		if err := ValidateGreaterThanZero(c.Gamma); err != nil {
			return fmt.Errorf("invalid gamma: %w", err)
		}
	}

	if err := ValidateGreaterThanZero(c.Threads); err != nil {
		return fmt.Errorf("invalid threads: %w", err)
	}

	if err := ValidateGreaterThanZero(c.SymmetryLevel); err != nil {
		return fmt.Errorf("invalid symmetry_level: %w", err)
	}

	if err := ValidateWritableDir(c.Output); err != nil {
		return fmt.Errorf("invalid path: %w", err)
	}

	return nil
}

func (s Size) Validate() error {
	if err := ValidateGreaterThanZero(s.Width); err != nil {
		return fmt.Errorf("invalid width: %w", err)
	}

	if err := ValidateGreaterThanZero(s.Height); err != nil {
		return fmt.Errorf("invalid height: %w", err)
	}

	return nil
}

func ValidateGreaterThanZero[Number int | float64](v Number) error {
	if v <= 0 {
		return fmt.Errorf("must be > 0, got %v", v)
	}

	return nil
}

func ValidateWritableDir(path string) error {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); err != nil {
		return fmt.Errorf("dir %q does not exist: %w", dir, err)
	}

	f, err := os.CreateTemp(dir, "*.tmp")
	if err != nil {
		return fmt.Errorf("unable to create temporary file in provided directory: %w", err)
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("unable to close temporary file: %w", err)
	}

	if err := os.Remove(f.Name()); err != nil {
		return fmt.Errorf("unable to remove temporary file: %w", err)
	}

	return nil
}

func validateWeightedFunctions(funcs []WeightedFunction) error {
	if len(funcs) == 0 {
		return errors.New("no functions specified")
	}

	for i, wf := range funcs {
		if !variation.Name(wf.Name).IsValid() {
			return fmt.Errorf("function #%d: unknown variation %q", i, wf.Name)
		}

		if wf.Weight < 0 {
			return fmt.Errorf("function #%d: negative weight %.2f", i, wf.Weight)
		}
	}

	return nil
}

func validateAffineParams(params []AffineParams) error {
	if len(params) == 0 {
		return errors.New("no affine parameters specified")
	}

	return nil
}
