package fractal_usecase_test

import (
	"log/slog"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/application/fractal_usecase"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/fractal_config"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/pixels"
)

func TestExecute_HappyPath(t *testing.T) {
	log := slog.New(slog.DiscardHandler)

	path := "out.png"
	cfg := &fractal_config.Config{OutputPath: path}
	pixels := &pixels.Pixels{}
	img := pixels.Image()

	generator := NewMockGenerator(t)
	generator.EXPECT().GenerateFractal(cfg).Return(pixels).Once()

	imageSaver := NewMockImageSaver(t)
	imageSaver.EXPECT().SaveImage(img, path).Return(nil).Once()

	usecase := fractal_usecase.New(log, generator, imageSaver)
	err := usecase.Execute(t.Context(), cfg)
	require.NoError(t, err)
}
