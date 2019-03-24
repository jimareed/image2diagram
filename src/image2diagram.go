package main

import (
	"fmt"
	"image/color"
	"image/png"
	"log"
	"os"
)

func image2Diagram(imageFileName string) string {
	file, err := os.Open(imageFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(os.Stderr, "%s: %v\n", imageFileName, err)
	}

	b := img.Bounds()

	type point struct {
		x, y int
	}
	points := []point{}

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
				points = append(points, point{x, y})
			}
		}
	}

	blocks := ""
	if len(points) > 0 {
		blocks = fmt.Sprintf("{\"x\": %d, \"y\": %d}", points[0].x, points[0].y)
	}

	diagram := fmt.Sprintf(
		"{"+
			"\"width\": %d,"+
			"\"height\": %d,"+
			"\"blocks\": [%s],"+
			"\"connectors\": []"+
			"}\n",
		b.Max.X, b.Max.Y, blocks)

	return diagram
}
