package encoder

import (
	"image-wav/utils"
	"image/png"
	"os"
)

// Encoder encodes a given slice of binary data into an image
type Encoder struct {
	binary []byte
	canvas *utils.Canvas

	width  int
	height int
}

// NewEncoder creates a new encoder using a given binary slice, and dimensions values
func NewEncoder(binary []byte, width int, height int) *Encoder {
	canvas := utils.NewCanvas(width, height)

	return &Encoder{
		binary: binary,
		canvas: canvas,
		width:  width,
		height: height,
	}
}

// Encode encodes the binary slice into image RGB pixels
func (e *Encoder) Encode() {
	for i := 0; i < len(e.binary) && !e.canvas.IsComplete(); i += 3 {
		pixel := utils.BytesToRGBA(e.binary[i : i+3])
		e.canvas.AddPixel(pixel)
	}
}

// Out renders the image using the collected pixel values from the Encode method
func (e *Encoder) Out(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	if err = png.Encode(f, e.canvas.RGBA); err != nil {
		return err
	}

	return nil
}
