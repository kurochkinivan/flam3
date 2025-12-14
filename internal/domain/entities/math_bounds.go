package entities

const (
	DefaultXMax = 1.777
	DefaultXMin = -1.777
	DefaultYMax = 1.0
	DefaultYMin = -1.0
)

// MathBounds представляет математические границы области просмотра
type MathBounds struct {
	xMin, xMax float64
	yMin, yMax float64
}

func NewMathBounds(xMin, xMax, yMin, yMax float64) MathBounds {
	return MathBounds{
		xMin: xMin,
		xMax: xMax,
		yMin: yMin,
		yMax: yMax,
	}
}

func DefaultMathBounds() MathBounds {
	return NewMathBounds(DefaultXMin, DefaultXMax, DefaultYMin, DefaultYMax)
}

func (m MathBounds) XMin() float64 {
	return m.xMin
}

func (m MathBounds) XMax() float64 {
	return m.xMax
}

func (m MathBounds) YMin() float64 {
	return m.yMin
}

func (m MathBounds) YMax() float64 {
	return m.yMax
}

func (m MathBounds) InBoundsX(x float64) bool {
	return m.xMin <= x && x <= m.xMax
}

func (m MathBounds) InBoundsY(y float64) bool {
	return m.yMin <= y && y <= m.yMax
}

func (m MathBounds) RangeX() float64 {
	return m.xMax - m.xMin
}

func (m MathBounds) RangeY() float64 {
	return m.yMax - m.yMin
}
