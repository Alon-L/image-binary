package main

import (
	"github.com/urfave/cli/v2"
	"image-wav/decoder"
	"image-wav/encoder"
	"image/png"
	"io/ioutil"
	"os"
)

// Encodes a given binary file into an image with given dimensions
func encode(in string, out string, width int, height int) {

	binary, err := ioutil.ReadFile(in)
	if err != nil {
		panic(err)
	}

	encoder := encoder.NewEncoder(binary, width, height)
	encoder.Encode()

	if err = encoder.Out(out); err != nil {
		panic(err)
	}
}

// Decodes a given image file into a binary file
func decode(in string, out string) {
	file, err := os.Open(in)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err = file.Seek(0, 0); err != nil {
		panic(err)
	}

	img, err := png.Decode(file)
	if err != nil {
		panic(err)
	}

	decoder := decoder.NewDecoder(img)
	decoder.Decode()

	err = decoder.Out(out)
	if err != nil {
		panic(err)
	}
}

func actionMiddleware(c *cli.Context) error {
	if c.Args().Len() < 2 {
		return cli.Exit("Too little arguments were given", 1)
	}

	return nil
}

func main() {
	app := &cli.App{
		Name:  "image-binary",
		Usage: "encode and decode binary files into images",
		Commands: []*cli.Command{
			{
				Name:      "encode",
				Usage:     "Encodes a binary file into an image",
				Aliases:   []string{"e"},
				ArgsUsage: "[in] [out]",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:     "width",
						Aliases:  []string{"x"},
						Usage:    "Rendered image width",
						Required: false,
						Value:    1000,
					},
					&cli.IntFlag{
						Name:     "height",
						Aliases:  []string{"y"},
						Usage:    "Rendered image height",
						Required: false,
						Value:    1000,
					},
				},
				Action: func(c *cli.Context) error {
					if err := actionMiddleware(c); err != nil {
						return err
					}
					encode(c.Args().Get(0), c.Args().Get(1), c.Int("width"), c.Int("height"))
					return nil
				},
			},
			{
				Name:      "decode",
				Usage:     "Decodes an image into a binary file",
				Aliases:   []string{"e"},
				ArgsUsage: "[in] [out]",
				Action: func(c *cli.Context) error {
					if err := actionMiddleware(c); err != nil {
						return err
					}
					decode(c.Args().Get(0), c.Args().Get(1))
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
