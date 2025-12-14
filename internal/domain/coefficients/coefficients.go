package coefficients

import (
	"fmt"
	"image/color"
)

type Coefficients struct {
	A, B, D, E float64
	C, F       float64
	Color      color.RGBA
}

func New(a, b, d, e float64, c, f float64, color color.RGBA) Coefficients {
	return Coefficients{
		A:     a,
		B:     b,
		D:     d,
		E:     e,
		C:     c,
		F:     f,
		Color: color,
	}
}

func (c Coefficients) String() string {
	return fmt.Sprintf(
		"Coefficients{A:%.3f, B:%.3f, C:%.3f, D:%.3f, E:%.3f, F:%.3f, Color:RGBA(%d,%d,%d,%d)}",
		c.A, c.B, c.C, c.D, c.E, c.F,
		c.Color.R, c.Color.G, c.Color.B, c.Color.A,
	)
}
