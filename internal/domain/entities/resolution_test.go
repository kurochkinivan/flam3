package entities_test

import (
	"testing"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/entities"
)

func TestResolutionInBounds(t *testing.T) {
	width, height := 1920, 1080
	res := entities.NewResolution(width, height)

	tests := []struct {
		name  string
		x     int
		y     int
		wantX bool
		wantY bool
	}{
		{"top-left", 0, 0, true, true},
		{"bottom-right", width - 1, height - 1, true, true},
		{"x negative", -1, height / 2, false, true},
		{"y negative", width / 2, -1, true, false},
		{"x equal width", width, height / 2, false, true},
		{"y equal height", width / 2, height, true, false},
		{"just inside top-left", 1, 1, true, true},
		{"just inside bottom-right", width - 2, height - 2, true, true},
		{"just outside top-left", -1, -1, false, false},
		{"just outside bottom-right", width, height, false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := res.InBoundsPixelX(tt.x); got != tt.wantX {
				t.Errorf("InBoundsPixelX(%d) = %v, want %v", tt.x, got, tt.wantX)
			}
			if got := res.InBoundsPixelY(tt.y); got != tt.wantY {
				t.Errorf("InBoundsPixelY(%d) = %v, want %v", tt.y, got, tt.wantY)
			}
		})
	}
}
