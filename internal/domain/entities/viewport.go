package entities

// Viewport unites Resolution and MathBounds
// and provides methods for converting math coordinates to pixels.
type Viewport struct {
	Resolution
	MathBounds
}

func NewViewport(resolution Resolution, bounds MathBounds) Viewport {
	return Viewport{
		Resolution: resolution,
		MathBounds: bounds,
	}
}

// XToPixel converts math X-coordinate to pixel.
func (v Viewport) XToPixel(x float64) int {
	pixel := ((v.MathBounds.xMax - x) / v.MathBounds.RangeX()) * float64(v.Resolution.Width())
	return int(pixel)
}

// YToPixel converts math Y-coordinate to pixel.
func (v Viewport) YToPixel(y float64) int {
	pixel := ((v.MathBounds.yMax - y) / v.MathBounds.RangeY()) * float64(v.Resolution.Height())
	return int(pixel)
}
