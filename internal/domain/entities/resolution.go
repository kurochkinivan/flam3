package entities

// Resolution представляет разрешение экрана в пикселях
type Resolution struct {
	width  int
	height int
}

func NewResolution(width, height int) Resolution {
	return Resolution{
		width:  width,
		height: height,
	}
}

func (r Resolution) Width() int {
	return r.width
}

func (r Resolution) Height() int {
	return r.height
}

func (r Resolution) InBoundsPixelX(x int) bool {
	return 0 <= x && x < r.width
}

func (r Resolution) InBoundsPixelY(y int) bool {
	return 0 <= y && y < r.height
}
