package terminal

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v3"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/pixels"
)

func TestRun_HappyPath_Flags(t *testing.T) {
	handler := MockedHandler(t)

	osArgs := []string{
		"flam3",
		"--width=1920",
		"--height=1080",
		"--seed=2.1324512",
		"-i=2500",
		"-o=result.png",
		"-t=8",
		"-f=swirl:10.0,horseshoe:0.7",
		"-ap=1.0,1.0,1.0,1.0,1.0,1.0/0.3,1.0,-0.2,0.4,1.0,1.0",
	}

	err := handler.Run(t.Context(), osArgs)
	require.NoError(t, err)
}

func TestRun_NoParams_Flags(t *testing.T) {
	handler := New("test_version", nil, nil)

	osArgs := []string{
		"flam3",
	}

	cli.OsExiter = func(code int) {
		assert.Equal(t, ExitCodeInvalidInput, code)
	}

	err := handler.Run(t.Context(), osArgs)
	require.Error(t, err)
}

func TestRun_HappyPath_JSONConfig(t *testing.T) {
	handler := MockedHandler(t)

	json := []byte(`{
  "size": {
    "width": 1920,
    "height": 1080
  },
  "iteration_count": 1000,
  "output_path": "result.png",
  "threads": 12,
  "seed": 2.1324512,
  "functions": [
    {
      "name": "swirl",
      "weight": 2.0
    }
  ],
  "affine_params": [
    {
      "a": 1,
      "b": 1,
      "c": 1,
      "d": 1,
      "e": 1,
      "f": 1
    },
    {
      "a": 0.3,
      "b": 1,
      "c": -0.2,
      "d": 0.4,
      "e": 1,
      "f": 1
    }
  ],
  "gamma_correction": true,
  "gamma": 2.2,
  "symmetry_level": 1
}`)

	f, err := os.CreateTemp(t.TempDir(), "*.json")
	require.NoError(t, err)

	written, err := f.Write(json)
	require.NoError(t, err)
	assert.Len(t, json, written)

	require.NoError(t, f.Close())

	osArgs := []string{
		"flam3",
		"--config=" + f.Name(),
	}

	err = handler.Run(t.Context(), osArgs)

	require.NoError(t, err)
}

func TestRun_NoParams_JSON(t *testing.T) {
	f, err := os.CreateTemp(t.TempDir(), "*.json")
	require.NoError(t, err)
	require.NoError(t, f.Close())

	osArgs := []string{
		"flam3",
		"--config=" + f.Name(),
	}

	cli.OsExiter = func(code int) {
		assert.Equal(t, ExitCodeInvalidInput, code)
	}

	handler := New("version", nil, nil)
	err = handler.Run(t.Context(), osArgs)
	require.Error(t, err)
}

func MockedHandler(t *testing.T) *Handler {
	mockFractalGenerator := NewMockFractalGenerator(t)
	mockImageSaver := NewMockImageSaver(t)
	pixels := &pixels.Pixels{}

	mockFractalGenerator.EXPECT().GenerateFractal(mock.Anything, mock.Anything).Return(pixels).Once()
	mockImageSaver.EXPECT().SaveImage(pixels.Image(), mock.Anything).Return(nil).Once()

	return New("test_version", mockFractalGenerator, mockImageSaver)
}
