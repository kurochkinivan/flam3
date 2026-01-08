package image_saver

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

type PNGSaver struct{}

func NewPNGSaver() *PNGSaver {
	return &PNGSaver{}
}

// SaveImage saves the given image to the specified path as a PNG file.
func (i *PNGSaver) SaveImage(img image.Image, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file %q: %w", path, err)
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		return fmt.Errorf("failed to encode png: %w", err)
	}

	return nil
}
