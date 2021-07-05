package utils

import (
	"image/color"
)

// BytesToRGBA encodes a given slice of bytes into an RGBA value
func BytesToRGBA(b []byte) *color.RGBA {
	return &color.RGBA{
		R: b[0],
		G: b[1],
		B: b[2],
		A: 255,
	}
}
