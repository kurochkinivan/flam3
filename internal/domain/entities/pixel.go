package entities

import (
	"image/color"
)

type Pixel struct {
	Color  color.RGBA
	Count  int
	Normal float64
}

func NewPixel(color color.RGBA) *Pixel {
	return &Pixel{
		Count:  0,
		Normal: 0,
		Color:  color,
	}
}
