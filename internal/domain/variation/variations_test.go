package variation

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinear(t *testing.T) {
	tests := []struct {
		name         string
		x, y         float64
		wantX, wantY float64
	}{
		{"positive", 1.0, 2.0, 1.0, 2.0},
		{"negative", -1.0, -2.0, -1.0, -2.0},
		{"zero", 0.0, 0.0, 0.0, 0.0},
		{"mixed", -3.5, 4.2, -3.5, 4.2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := linear(tt.x, tt.y)
			assert.Equal(t, tt.wantX, gotX)
			assert.Equal(t, tt.wantY, gotY)
		})
	}
}

func TestSinusoidal(t *testing.T) {
	tests := []struct {
		name         string
		x, y         float64
		wantX, wantY float64
	}{
		{"pi_half", math.Pi / 2, math.Pi, 1.0, 0.0},
		{"zero", 0.0, 0.0, 0.0, 0.0},
		{"negative", -math.Pi / 2, -math.Pi, -1.0, 0.0},
		{"pi_quarter", math.Pi / 4, math.Pi / 6, math.Sqrt(2) / 2, 0.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := sinusoidal(tt.x, tt.y)
			assert.InDelta(t, tt.wantX, gotX, 1e-10)
			assert.InDelta(t, tt.wantY, gotY, 1e-10)
		})
	}
}

func TestSpherical(t *testing.T) {
	tests := []struct {
		name         string
		x, y         float64
		wantX, wantY float64
		isNaN        bool
	}{
		{"unit_x", 1.0, 0.0, 1.0, 0.0, false},
		{"zero_division", 0.0, 0.0, 0.0, 0.0, true},
		{"quadrant_2", -1.0, 1.0, -0.5, 0.5, false},
		{"unit_circle", math.Sqrt(0.5), math.Sqrt(0.5), math.Sqrt(0.5), math.Sqrt(0.5), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := spherical(tt.x, tt.y)
			if tt.isNaN {
				assert.True(t, math.IsNaN(gotX) || math.IsInf(gotX, 0))
				assert.True(t, math.IsNaN(gotY) || math.IsInf(gotY, 0))
			} else {
				assert.InDelta(t, tt.wantX, gotX, 1e-10)
				assert.InDelta(t, tt.wantY, gotY, 1e-10)
			}
		})
	}
}

func TestSwirl(t *testing.T) {
	tests := []struct {
		name string
		x, y float64
	}{
		{"zero", 0.0, 0.0},
		{"unit_x", 1.0, 0.0},
		{"both_positive", 1.0, 1.0},
		{"mixed", -1.0, 1.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := swirl(tt.x, tt.y)
			assert.False(t, math.IsNaN(gotX))
			assert.False(t, math.IsNaN(gotY))
		})
	}
}

func TestHorseshoe(t *testing.T) {
	tests := []struct {
		name  string
		x, y  float64
		isNaN bool
	}{
		{"unit_x", 1.0, 0.0, false},
		{"zero_division", 0.0, 0.0, true},
		{"negative", -1.0, 1.0, false},
		{"equal", 1.0, 1.0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := horseshoe(tt.x, tt.y)
			if tt.isNaN {
				assert.True(t, math.IsNaN(gotX) || math.IsInf(gotX, 0))
				assert.True(t, math.IsNaN(gotY) || math.IsInf(gotY, 0))
			} else {
				assert.False(t, math.IsNaN(gotX))
				assert.False(t, math.IsNaN(gotY))
			}
		})
	}
}

func TestPolar(t *testing.T) {
	tests := []struct {
		name         string
		x, y         float64
		wantX, wantY float64
	}{
		{"positive_x", 1.0, 0.0, 0.0, 0.0},
		{"positive_y", 0.0, 1.0, 0.5, 0.0},
		{"zero", 0.0, 0.0, 0.0, -1.0},
		{"negative_x", -1.0, 0.0, 1.0, 0.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := polar(tt.x, tt.y)
			assert.InDelta(t, tt.wantX, gotX, 1e-10)
			assert.InDelta(t, tt.wantY, gotY, 1e-10)
		})
	}
}

func TestHandkerchief(t *testing.T) {
	tests := []struct {
		name string
		x, y float64
	}{
		{"zero", 0.0, 0.0},
		{"unit_x", 1.0, 0.0},
		{"both_positive", 1.0, 1.0},
		{"negative", -1.0, -1.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := handkerchief(tt.x, tt.y)
			assert.False(t, math.IsNaN(gotX))
			assert.False(t, math.IsNaN(gotY))
		})
	}
}

func TestHeart(t *testing.T) {
	tests := []struct {
		name string
		x, y float64
	}{
		{"zero", 0.0, 0.0},
		{"unit", 1.0, 0.0},
		{"positive", 1.0, 1.0},
		{"negative", -1.0, -1.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := heart(tt.x, tt.y)
			assert.False(t, math.IsNaN(gotX))
			assert.False(t, math.IsNaN(gotY))
		})
	}
}

func TestDisk(t *testing.T) {
	tests := []struct {
		name string
		x, y float64
	}{
		{"zero", 0.0, 0.0},
		{"unit_x", 1.0, 0.0},
		{"unit_y", 0.0, 1.0},
		{"both", 1.0, 1.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := disk(tt.x, tt.y)
			assert.False(t, math.IsNaN(gotX))
			assert.False(t, math.IsNaN(gotY))
		})
	}
}

func TestSpiral(t *testing.T) {
	tests := []struct {
		name  string
		x, y  float64
		isNaN bool
	}{
		{"zero_division", 0.0, 0.0, true},
		{"unit_x", 1.0, 0.0, false},
		{"positive", 1.0, 1.0, false},
		{"negative", -1.0, -1.0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := spiral(tt.x, tt.y)
			if tt.isNaN {
				assert.True(t, math.IsNaN(gotX) || math.IsInf(gotX, 0))
				assert.True(t, math.IsNaN(gotY) || math.IsInf(gotY, 0))
			} else {
				assert.False(t, math.IsNaN(gotX))
				assert.False(t, math.IsNaN(gotY))
			}
		})
	}
}

func TestHyperbolic(t *testing.T) {
	tests := []struct {
		name         string
		x, y         float64
		wantX, wantY float64
		checkNaN     bool
	}{
		{"zero_division", 0.0, 0.0, 0.0, 0.0, true}, // newX=NaN, newY=0
		{"unit_x", 1.0, 0.0, 0.0, 1.0, false},
		{"positive", 1.0, 1.0, math.Sin(math.Pi/4) / math.Sqrt(2), math.Sqrt(2) * math.Cos(math.Pi/4), false},
		{"negative", -1.0, 1.0, math.Sin(3*math.Pi/4) / math.Sqrt(2), math.Sqrt(2) * math.Cos(3*math.Pi/4), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := hyperbolic(tt.x, tt.y)
			if tt.checkNaN {
				assert.True(t, math.IsNaN(gotX) || math.IsInf(gotX, 0))
				assert.InDelta(t, tt.wantY, gotY, 1e-10) // newY = 0, не NaN
			} else {
				assert.InDelta(t, tt.wantX, gotX, 1e-10)
				assert.InDelta(t, tt.wantY, gotY, 1e-10)
			}
		})
	}
}

func TestDiamond(t *testing.T) {
	tests := []struct {
		name string
		x, y float64
	}{
		{"zero", 0.0, 0.0},
		{"unit_x", 1.0, 0.0},
		{"positive", 1.0, 1.0},
		{"negative", -1.0, -1.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := diamond(tt.x, tt.y)
			assert.False(t, math.IsNaN(gotX))
			assert.False(t, math.IsNaN(gotY))
		})
	}
}

func TestEx(t *testing.T) {
	tests := []struct {
		name string
		x, y float64
	}{
		{"zero", 0.0, 0.0},
		{"unit", 1.0, 0.0},
		{"positive", 1.0, 1.0},
		{"negative", -1.0, -1.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := ex(tt.x, tt.y)
			assert.False(t, math.IsNaN(gotX))
			assert.False(t, math.IsNaN(gotY))
		})
	}
}

func TestBent(t *testing.T) {
	tests := []struct {
		name         string
		x, y         float64
		wantX, wantY float64
	}{
		{"both_positive", 2.0, 3.0, 2.0, 3.0},
		{"x_neg_y_pos", -2.0, 3.0, -4.0, 3.0},
		{"x_pos_y_neg", 2.0, -3.0, 2.0, -1.5},
		{"both_negative", -2.0, -3.0, -4.0, -1.5},
		{"zero", 0.0, 0.0, 0.0, 0.0},
		{"x_zero_y_pos", 0.0, 1.0, 0.0, 1.0},
		{"x_zero_y_neg", 0.0, -1.0, 0.0, -0.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := bent(tt.x, tt.y)
			assert.Equal(t, tt.wantX, gotX)
			assert.Equal(t, tt.wantY, gotY)
		})
	}
}

func TestFisheye(t *testing.T) {
	tests := []struct {
		name         string
		x, y         float64
		wantX, wantY float64
	}{
		{"zero", 0.0, 0.0, 0.0, 0.0},
		{"unit_x", 1.0, 0.0, 0.0, 1.0},
		{"unit_y", 0.0, 1.0, 1.0, 0.0},
		{"positive", 1.0, 1.0, 2.0 / (1.0 + math.Sqrt(2)), 2.0 / (1.0 + math.Sqrt(2))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := fisheye(tt.x, tt.y)
			assert.InDelta(t, tt.wantX, gotX, 1e-10)
			assert.InDelta(t, tt.wantY, gotY, 1e-10)
		})
	}
}

func TestEyefish(t *testing.T) {
	tests := []struct {
		name         string
		x, y         float64
		wantX, wantY float64
	}{
		{"zero", 0.0, 0.0, 0.0, 0.0},
		{"unit_x", 1.0, 0.0, 1.0, 0.0},
		{"unit_y", 0.0, 1.0, 0.0, 1.0},
		{"positive", 1.0, 1.0, 2.0 / (1.0 + math.Sqrt(2)), 2.0 / (1.0 + math.Sqrt(2))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := eyefish(tt.x, tt.y)
			assert.InDelta(t, tt.wantX, gotX, 1e-10)
			assert.InDelta(t, tt.wantY, gotY, 1e-10)
		})
	}
}

func TestBubble(t *testing.T) {
	tests := []struct {
		name         string
		x, y         float64
		wantX, wantY float64
	}{
		{"zero", 0.0, 0.0, 0.0, 0.0},
		{"unit_x", 1.0, 0.0, 0.8, 0.0},
		{"unit_y", 0.0, 1.0, 0.0, 0.8},
		{"positive", 2.0, 2.0, 8.0 / 12.0, 8.0 / 12.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := bubble(tt.x, tt.y)
			assert.InDelta(t, tt.wantX, gotX, 1e-10)
			assert.InDelta(t, tt.wantY, gotY, 1e-10)
		})
	}
}

func TestCylinder(t *testing.T) {
	tests := []struct {
		name         string
		x, y         float64
		wantX, wantY float64
	}{
		{"zero", 0.0, 2.0, 0.0, 2.0},
		{"pi_half", math.Pi / 2, 3.0, 1.0, 3.0},
		{"pi", math.Pi, 1.0, 0.0, 1.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := cylinder(tt.x, tt.y)
			assert.InDelta(t, tt.wantX, gotX, 1e-10)
			assert.Equal(t, tt.wantY, gotY)
		})
	}
}

func TestTangent(t *testing.T) {
	tests := []struct {
		name string
		x, y float64
	}{
		{"zero", 0.0, 0.0},
		{"small_values", 0.1, 0.1},
		{"pi_quarter", math.Pi / 4, math.Pi / 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := tangent(tt.x, tt.y)
			assert.False(t, math.IsNaN(gotX))
			assert.False(t, math.IsNaN(gotY))
		})
	}
}

func TestCross(t *testing.T) {
	tests := []struct {
		name  string
		x, y  float64
		isNaN bool
	}{
		{"equal_xy", 1.0, 1.0, true},
		{"different", 2.0, 1.0, false},
		{"negative", -2.0, 1.0, false},
		{"zero_x", 0.0, 1.0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := cross(tt.x, tt.y)
			if tt.isNaN {
				assert.True(t, math.IsNaN(gotX) || math.IsInf(gotX, 0))
				assert.True(t, math.IsNaN(gotY) || math.IsInf(gotY, 0))
			} else {
				assert.False(t, math.IsNaN(gotX))
				assert.False(t, math.IsNaN(gotY))
			}
		})
	}
}

func TestExponential(t *testing.T) {
	tests := []struct {
		name string
		x, y float64
	}{
		{"zero", 0.0, 0.0},
		{"one", 1.0, 0.0},
		{"positive", 2.0, 0.5},
		{"negative", -1.0, -0.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := exponential(tt.x, tt.y)
			assert.False(t, math.IsNaN(gotX))
			assert.False(t, math.IsNaN(gotY))
		})
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		name string
		x, y float64
	}{
		{"unit_x", 1.0, 0.0},
		{"positive", 1.0, 1.0},
		{"small", 0.5, 0.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := power(tt.x, tt.y)
			assert.False(t, math.IsNaN(gotX))
			assert.False(t, math.IsNaN(gotY))
		})
	}
}

func TestCosine(t *testing.T) {
	tests := []struct {
		name string
		x, y float64
	}{
		{"zero", 0.0, 0.0},
		{"positive", 0.5, 0.5},
		{"one", 1.0, 1.0},
		{"negative", -0.5, -0.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := cosine(tt.x, tt.y)
			assert.False(t, math.IsNaN(gotX))
			assert.False(t, math.IsNaN(gotY))
		})
	}
}
