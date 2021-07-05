package utils

import (
	"image"
	"image/color"
)

// Canvas is a utility structure helped to keep track of the image's pixel values
type Canvas struct {
	width  int
	height int
	RGBA   *image.RGBA

	context struct {
		x int
		y int
	}
}

// NewCanvas creates a new canvas using given dimensions
func NewCanvas(width int, height int) *Canvas {
	upLeft := image.Point{}
	lowRight := image.Point{X: width, Y: height}

	rgba := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

	return &Canvas{
		width:  width,
		height: height,
		RGBA:   rgba,
	}
}

// ReadPixel reads the pixel behind the context's current position. Increments the context position
func (c *Canvas) ReadPixel() (uint8, uint8, uint8, uint8) {
	r, g, b, a := c.RGBA.At(c.context.x, c.context.y).RGBA()
	c.IncContext()

	return uint8(r / 257), uint8(g / 257), uint8(b / 257), uint8(a / 257)
}

// AddPixel paints a given RGBA value into the context's current position. Increments the context position
func (c *Canvas) AddPixel(rgba *color.RGBA) {
	c.RGBA.Set(c.context.x, c.context.y, rgba)
	c.IncContext()
}

// IncContext increments the current context position. Move to the next line if the current ends.
func (c *Canvas) IncContext() {
	c.context.x++
	if c.context.x >= c.width {
		c.context.x = 0
		c.context.y++
	}
}

// IsComplete checks whether the canvas space has been all used.
func (c *Canvas) IsComplete() bool {
	return c.context.y >= c.height
}
