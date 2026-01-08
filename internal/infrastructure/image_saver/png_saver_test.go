package image_saver_test

import (
	"image"
	"os"
	"path/filepath"
	"testing"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/infrastructure/image_saver"
)

func TestSaveImageHappyPath(t *testing.T) {
	saver := image_saver.NewPNGSaver()

	img := image.NewRGBA(image.Rect(0, 0, 1, 1))
	path := filepath.Join(t.TempDir(), "out.png")

	if err := saver.SaveImage(img, path); err != nil {
		t.Fatalf("failed: %v", err)
	}

	if _, err := os.Stat(path); err != nil {
		t.Fatalf("file not created: %v", err)
	}
}
