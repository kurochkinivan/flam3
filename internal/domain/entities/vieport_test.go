package entities_test

import (
	"testing"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/entities"
)

func TestViewportPixelConversion(t *testing.T) {
	width, height := 200, 100
	xMin, xMax := 0.0, 10.0
	yMin, yMax := 0.0, 5.0

	res := entities.NewResolution(width, height)
	bounds := entities.NewMathBounds(xMin, xMax, yMin, yMax)
	vp := entities.NewViewport(res, bounds)

	tests := []struct {
		name  string
		x     float64
		y     float64
		wantX int
		wantY int
	}{
		{"top-left", xMin, yMax, width, 0},
		{"bottom-right", xMax, yMin, 0, height},
		{"center", (xMin + xMax) / 2, (yMin + yMax) / 2, width / 2, height / 2},
		{"just inside top-left", xMin + 0.0001, yMax - 0.0001, width - 1, 0},
		{"just inside bottom-right", xMax - 0.0001, yMin + 0.0001, 0, height - 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX := vp.XToPixel(tt.x)
			gotY := vp.YToPixel(tt.y)
			if gotX != tt.wantX {
				t.Errorf("XToPixel(%v) = %v, want %v", tt.x, gotX, tt.wantX)
			}
			if gotY != tt.wantY {
				t.Errorf("YToPixel(%v) = %v, want %v", tt.y, gotY, tt.wantY)
			}
		})
	}
}
