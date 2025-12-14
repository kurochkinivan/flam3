package terminal

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v3"
)

func TestHandler_FractalFlameCommand(t *testing.T) {
	t.Parallel()

	mockUsecase := NewMockFractalUsecase(t)
	h := New("1.0.0", mockUsecase)

	// Create a temporary config file for testing
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "config.json")
	configContent := `{
		"size": {"width": 100, "height": 100},
		"seed": 123.0,
		"iteration_count": 1000000,
		"output_path": "output.png",
		"threads": 4,
		"affine_params": [{"A": 1.0, "B": 0.0, "C": 0.0, "D": 0.0, "E": 1.0, "F": 0.0}],
		"functions": [{"name": "linear", "weight": 1.0}],
		"gamma_correction": true,
		"gamma": 2.2,
		"symmetry_level": 0
	}`
	err := os.WriteFile(configPath, []byte(configContent), 0644)
	require.NoError(t, err)

	cmd := &cli.Command{}
	cmd.Flags = []cli.Flag{
		&cli.StringFlag{Name: "config"},
	}
	cmd.Set("config", configPath)

	mockUsecase.EXPECT().Execute(context.Background(), mock.Anything).Return(nil).Once()

	err = h.FractalFlameCommand(context.Background(), cmd)
	assert.NoError(t, err)
}

func TestHandler_loadInputConfig(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "config.json")
	configContent := `{
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
		}`
	err := os.WriteFile(configPath, []byte(configContent), 0644)
	require.NoError(t, err)

	cmd := &cli.Command{}
	cmd.Flags = []cli.Flag{
		&cli.StringFlag{Name: "config"},
	}
	cmd.Set("config", configPath)

	h := &Handler{}
	cfg, err := h.loadInputConfig(cmd)

	require.NoError(t, err)
	assert.Equal(t, 1920, cfg.Size.Width)
	assert.Equal(t, 1080, cfg.Size.Height)
	assert.InDelta(t, 5.0, cfg.Seed, 0.01)
	assert.Equal(t, 2500, cfg.Iterations)
	assert.Equal(t, "fractal.png", cfg.Output)
	assert.Equal(t, 1, cfg.Threads)
	assert.Len(t, cfg.AffineParams, 1)
	assert.Len(t, cfg.WeightedFunctions, 1)
	assert.False(t, cfg.GammaCorrection)
	assert.InDelta(t, 2.2, cfg.Gamma, 0.01)
	assert.Equal(t, 1, cfg.SymmetryLevel)
}
