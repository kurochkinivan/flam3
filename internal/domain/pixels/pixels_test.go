package pixels_test

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/suite"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/entities"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain/pixels"
)

type PixelsTestSuite struct {
	suite.Suite

	resolution entities.Resolution
	pixels     *pixels.Pixels
}

func (suite *PixelsTestSuite) SetupTest() {
	suite.resolution = entities.NewResolution(10, 10)
	suite.pixels = pixels.NewPixels(suite.resolution)
}

func (suite *PixelsTestSuite) TestNewPixels() {
	suite.Equal(suite.resolution.Width(), suite.pixels.Width())
	suite.Equal(suite.resolution.Height(), suite.pixels.Height())

	// Check that all pixels are initialized
	for y := range suite.pixels.Height() {
		for x := range suite.pixels.Width() {
			pix := suite.pixels.Pix(x, y)
			suite.NotNil(pix)
			suite.Equal(uint8(255), pix.Color.A)
			suite.Equal(0, pix.Count)
			suite.InDelta(0.0, pix.Normal, 0.01)
		}
	}
}

func (suite *PixelsTestSuite) TestPix() {
	// Test getting pixel at valid coordinates
	pix := suite.pixels.Pix(5, 5)
	suite.NotNil(pix)

	// Modify pixel and verify
	newCount := 42
	newColor := color.RGBA{R: 100, G: 150, B: 200, A: 255}
	newNormal := 3.14

	pix.Count = newCount
	pix.Color = newColor
	pix.Normal = newNormal

	retrieved := suite.pixels.Pix(5, 5)
	suite.Equal(newCount, retrieved.Count)
	suite.Equal(newColor, retrieved.Color)
	suite.InDelta(newNormal, retrieved.Normal, 0.01)
}

func (suite *PixelsTestSuite) TestApplyGammaFactor() {
	// Set up pixels with different counts
	suite.pixels.Pix(0, 0).Count = 1
	suite.pixels.Pix(1, 1).Count = 10
	suite.pixels.Pix(2, 2).Count = 100
	suite.pixels.Pix(3, 3).Count = 1000

	// Set initial colors
	for y := range suite.pixels.Height() {
		for x := range suite.pixels.Width() {
			suite.pixels.Pix(x, y).Color = color.RGBA{R: 255, G: 255, B: 255, A: 255}
		}
	}

	gamma := 2.2
	suite.pixels.ApplyGammaFactor(gamma)

	// Check that pixels with counts have been processed
	pix1 := suite.pixels.Pix(0, 0)
	pix10 := suite.pixels.Pix(1, 1)
	pix100 := suite.pixels.Pix(2, 2)
	pix1000 := suite.pixels.Pix(3, 3)

	// The pixel with highest count (1000, log10=3) should be brightest
	suite.Greater(pix1000.Normal, pix100.Normal)
	suite.Greater(pix100.Normal, pix10.Normal)
	suite.Greater(pix10.Normal, pix1.Normal)

	// All pixels should have A=255
	suite.Equal(uint8(255), pix1.Color.A)
	suite.Equal(uint8(255), pix10.Color.A)
	suite.Equal(uint8(255), pix100.Color.A)
	suite.Equal(uint8(255), pix1000.Color.A)
}

func (suite *PixelsTestSuite) TestApplyGammaFactor_ZeroCounts() {
	// All pixels have count = 0
	for y := range suite.pixels.Height() {
		for x := range suite.pixels.Width() {
			suite.pixels.Pix(x, y).Count = 0
		}
	}

	// Should not panic
	suite.NotPanics(func() {
		suite.pixels.ApplyGammaFactor(2.2)
	})
}

func (suite *PixelsTestSuite) TestMerge() {
	// Create two pixels instances
	pixels1 := pixels.NewPixels(suite.resolution)
	pixels2 := pixels.NewPixels(suite.resolution)

	// Set up pixels1
	color1 := color.RGBA{R: 100, G: 100, B: 100, A: 255}
	count1 := 2
	pixels1.Pix(0, 0).Color = color1
	pixels1.Pix(0, 0).Count = count1

	color2 := color.RGBA{R: 200, G: 200, B: 200, A: 255}
	count2 := 4
	pixels1.Pix(1, 1).Color = color2
	pixels1.Pix(1, 1).Count = count2

	// Set up pixels2
	color3 := color.RGBA{R: 50, G: 50, B: 50, A: 255}
	count3 := 2
	pixels2.Pix(0, 0).Color = color3
	pixels2.Pix(0, 0).Count = count3

	color4 := color.RGBA{R: 150, G: 150, B: 150, A: 255}
	count4 := 3
	pixels2.Pix(2, 2).Color = color4
	pixels2.Pix(2, 2).Count = count4

	// Merge pixels2 into pixels1
	pixels1.Merge(pixels2)

	// Check merged pixel (0,0): should be average of color1 and color3 with weights count1 and count3
	merged00 := pixels1.Pix(0, 0)
	expectedCount00 := count1 + count3
	expectedR00 := uint8(
		(float64(color1.R)*float64(count1) + float64(color3.R)*float64(count3)) / float64(expectedCount00),
	)
	expectedG00 := uint8(
		(float64(color1.G)*float64(count1) + float64(color3.G)*float64(count3)) / float64(expectedCount00),
	)
	expectedB00 := uint8(
		(float64(color1.B)*float64(count1) + float64(color3.B)*float64(count3)) / float64(expectedCount00),
	)
	expectedColor00 := color.RGBA{R: expectedR00, G: expectedG00, B: expectedB00, A: 255}

	suite.Equal(expectedCount00, merged00.Count)
	suite.Equal(expectedColor00, merged00.Color)

	// Check pixel (1,1): should remain unchanged (no merge from pixels2)
	merged11 := pixels1.Pix(1, 1)
	suite.Equal(count2, merged11.Count)
	suite.Equal(color2, merged11.Color)

	// Check pixel (2,2): should be copied from pixels2
	merged22 := pixels1.Pix(2, 2)
	suite.Equal(count4, merged22.Count)
	suite.Equal(color4, merged22.Color)
}

func (suite *PixelsTestSuite) TestMergeEmptyPixels() {
	pixels1 := pixels.NewPixels(suite.resolution)
	pixels2 := pixels.NewPixels(suite.resolution)

	// pixels2 has all zero counts, should not affect pixels1
	originalColor := color.RGBA{R: 100, G: 100, B: 100, A: 255}
	originalCount := 5
	pixels1.Pix(5, 5).Color = originalColor
	pixels1.Pix(5, 5).Count = originalCount

	pixels1.Merge(pixels2)

	suite.Equal(originalCount, pixels1.Pix(5, 5).Count)
	suite.Equal(originalColor, pixels1.Pix(5, 5).Color)
}

func (suite *PixelsTestSuite) TestImage() {
	// Set some pixel colors
	redColor := color.RGBA{R: 255, G: 0, B: 0, A: 255}
	greenColor := color.RGBA{R: 0, G: 255, B: 0, A: 255}
	blueColor := color.RGBA{R: 0, G: 0, B: 255, A: 255}

	suite.pixels.Pix(0, 0).Color = redColor
	suite.pixels.Pix(1, 1).Color = greenColor
	suite.pixels.Pix(2, 2).Color = blueColor

	img := suite.pixels.Image()

	// Check image dimensions
	bounds := img.Bounds()
	suite.Equal(0, bounds.Min.X)
	suite.Equal(0, bounds.Min.Y)
	suite.Equal(suite.pixels.Width(), bounds.Max.X)
	suite.Equal(suite.pixels.Height(), bounds.Max.Y)

	// Check pixel colors in the image
	r1, g1, b1, a1 := img.At(0, 0).RGBA()
	suite.Equal(uint32(redColor.R), r1>>8) // RGBA returns 16-bit values
	suite.Equal(uint32(redColor.G), g1>>8)
	suite.Equal(uint32(redColor.B), b1>>8)
	suite.Equal(uint32(redColor.A), a1>>8)

	r2, g2, b2, a2 := img.At(1, 1).RGBA()
	suite.Equal(uint32(greenColor.R), r2>>8)
	suite.Equal(uint32(greenColor.G), g2>>8)
	suite.Equal(uint32(greenColor.B), b2>>8)
	suite.Equal(uint32(greenColor.A), a2>>8)

	r3, g3, b3, a3 := img.At(2, 2).RGBA()
	suite.Equal(uint32(blueColor.R), r3>>8)
	suite.Equal(uint32(blueColor.G), g3>>8)
	suite.Equal(uint32(blueColor.B), b3>>8)
	suite.Equal(uint32(blueColor.A), a3>>8)
}

func TestPixelsTestSuite(t *testing.T) {
	suite.Run(t, new(PixelsTestSuite))
}
