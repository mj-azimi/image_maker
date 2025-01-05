package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	width, height := 3750, 6670
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			c := color.RGBA{R: uint8(y/150), G: uint8(5), B: uint8(x/90), A: uint8(250)}
			img.Set(x, y, c)
		}
	}

	file, err := os.Create("image_50x50.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}

