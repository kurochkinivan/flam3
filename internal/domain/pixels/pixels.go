package pixels

import (
	"image"
	"image/color"
	"math"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/entities"
)

type Pixels struct {
	pixels [][]*entities.Pixel
	entities.Resolution
}

func NewPixels(resolution entities.Resolution) *Pixels {
	pixels := make([][]*entities.Pixel, resolution.Height())

	for i := range pixels {
		pixels[i] = make([]*entities.Pixel, resolution.Width())
		for j := range pixels[i] {
			pixels[i][j] = entities.NewPixel(color.RGBA{A: 255})
		}
	}

	return &Pixels{
		pixels:     pixels,
		Resolution: resolution,
	}
}

func (p *Pixels) ApplyGammaFactor(gamma float64) {
	var maxNormal float64

	for y := range p.pixels {
		for x := range p.pixels[y] {
			if p.Pix(x, y).Count != 0 {
				p.Pix(x, y).Normal = math.Log10(float64(p.Pix(x, y).Count))
				maxNormal = max(maxNormal, p.Pix(x, y).Normal)
			}
		}
	}

	for y := range p.pixels {
		for x := range p.pixels[y] {
			pix := p.Pix(x, y)
			alphaNormal := pix.Normal / maxNormal

			pix.Color = color.RGBA{
				R: uint8(float64(pix.Color.R) * math.Pow(alphaNormal, 1.0/gamma)),
				G: uint8(float64(pix.Color.G) * math.Pow(alphaNormal, 1.0/gamma)),
				B: uint8(float64(pix.Color.B) * math.Pow(alphaNormal, 1.0/gamma)),
				A: 255,
			}
		}
	}
}

func (p *Pixels) Merge(p2 *Pixels) {
	for y := range p.pixels {
		for x := range p.pixels[y] {
			dst := p.Pix(x, y)
			src := p2.Pix(x, y)

			if src.Count == 0 {
				continue
			}

			if dst.Count == 0 {
				dst.Color = src.Color
				dst.Count = src.Count
				continue
			}

			oldWeight := float64(dst.Count)
			newWeight := float64(src.Count)
			totalWeight := oldWeight + newWeight

			dst.Color = color.RGBA{
				R: uint8((float64(dst.Color.R)*oldWeight + float64(src.Color.R)*newWeight) / totalWeight),
				G: uint8((float64(dst.Color.G)*oldWeight + float64(src.Color.G)*newWeight) / totalWeight),
				B: uint8((float64(dst.Color.B)*oldWeight + float64(src.Color.B)*newWeight) / totalWeight),
				A: 255,
			}

			dst.Count += src.Count
		}
	}
}

func (p Pixels) Image() image.Image {
	rect := image.Rect(0, 0, p.Width(), p.Height())
	img := image.NewRGBA(rect)

	for y := range p.pixels {
		for x := range p.pixels[y] {
			pix := p.Pix(x, y)

			img.SetRGBA(x, y, pix.Color)
		}
	}

	return img
}

func (p *Pixels) Pix(x, y int) *entities.Pixel {
	return p.pixels[y][x]
}
