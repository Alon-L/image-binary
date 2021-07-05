package decoder

import (
	"image"
	"image-wav/utils"
	"image/draw"
	"os"
)

// Decoder decodes a given image into binary bytes
type Decoder struct {
	img    image.Image
	binary []byte
	canvas *utils.Canvas
}

// NewDecoder creates a new decoder using a given image
func NewDecoder(img image.Image) *Decoder {
	b := img.Bounds()

	width := b.Size().X
	height := b.Size().Y

	canvas := utils.NewCanvas(width, height)
	draw.Draw(canvas.RGBA, canvas.RGBA.Bounds(), img, b.Min, draw.Src)

	return &Decoder{
		img:    img,
		canvas: canvas,
	}
}

// Decode decodes the image's pixels into the binary slice
func (d *Decoder) Decode() {
	for !d.canvas.IsComplete() {
		r, g, b, _ := d.canvas.ReadPixel()
		d.binary = append(d.binary, r, g, b)
	}
}

// Out writes the binary file using the bytes collected from the Decode method
func (d *Decoder) Out(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	if _, err = file.Write(d.binary); err != nil {
		return err
	}

	return nil
}
