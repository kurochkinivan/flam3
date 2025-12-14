package entities_test

import (
	"math"
	"testing"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/entities"
)

func TestMathBounds_InBoundsX(t *testing.T) {
	t.Parallel()

	b := entities.NewMathBounds(-1.0, 2.0, -5.0, 5.0)

	tests := []struct {
		name string
		x    float64
		want bool
	}{
		{"at min", -1.0, true},
		{"at max", 2.0, true},
		{"middle", 0.5, true},
		{"just above min", -0.999, true},
		{"just below max", 1.999, true},
		{"below min", -1.0001, false},
		{"above max", 2.0001, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b.InBoundsX(tt.x); got != tt.want {
				t.Errorf("InBoundsX(%v) = %v, want %v", tt.x, got, tt.want)
			}
		})
	}
}

func TestMathBounds_InBoundsY(t *testing.T) {
	t.Parallel()

	b := entities.NewMathBounds(-1.0, 2.0, -5.0, 5.0)

	tests := []struct {
		name string
		y    float64
		want bool
	}{
		{"at min", -5.0, true},
		{"at max", 5.0, true},
		{"middle", 0.0, true},
		{"just above min", -4.999, true},
		{"just below max", 4.999, true},
		{"below min", -5.0001, false},
		{"above max", 5.0001, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b.InBoundsY(tt.y); got != tt.want {
				t.Errorf("InBoundsY(%v) = %v, want %v", tt.y, got, tt.want)
			}
		})
	}
}

func TestMathBounds_RangeX(t *testing.T) {
	t.Parallel()

	b := entities.NewMathBounds(-1.5, 2.5, -10.0, 10.0)

	want := 4.0
	got := b.RangeX()

	if math.Abs(got-want) > 1e-9 {
		t.Errorf("RangeX() = %v, want %v", got, want)
	}
}

func TestMathBounds_RangeY(t *testing.T) {
	t.Parallel()

	b := entities.NewMathBounds(-1.5, 2.5, -10.0, 10.0)

	want := 20.0
	got := b.RangeY()

	if math.Abs(got-want) > 1e-9 {
		t.Errorf("RangeY() = %v, want %v", got, want)
	}
}
