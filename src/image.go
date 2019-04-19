package main

import (
	"image/color"
	"image/png"
	"log"
	"os"
)

func image2BlackWhiteImage(imageFileName string) BlackWhiteImage {
	file, err := os.Open(imageFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	b := img.Bounds()

	bwImage := BlackWhiteImage{}
	bwImage.width = b.Max.X
	bwImage.height = b.Max.Y

	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			level := c.Y / 51 // 51 * 5 = 255
			if level == 5 {
				level--
			}
			if level == 1 {
				level = 0
			}
			if level == 0 {
				bwImage.points = append(bwImage.points, Point{x, y})
			}
		}
	}

	return bwImage
}
