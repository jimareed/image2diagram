package main

import (
	"fmt"
	"log"
)

const maxShapes = 1000
const minDistanceBetweenShapes = 20

// Shape contains the x,y coordinates of the start of the shape
type Shape struct {
	x, y, width, height   int
	points []Point
}

func shape2String(shape Shape) string {

	s := fmt.Sprintf("{\"x\": %d, \"y\": %d}", shape.x, shape.y)

	return s
}

func blackWhiteImage2Shapes(bwi BlackWhiteImage) []Shape {

	shapes := []Shape{}

	return findShapes(bwi, shapes)
}

func findShapes(bwi BlackWhiteImage, shapes []Shape) []Shape {

	if len(bwi.points) == 0 {
		return shapes
	}

	if len(shapes) > maxShapes {
		log.Print("Max shapes exceeded")
		return shapes
	}

	newImage := BlackWhiteImage{}
	i := len(shapes)
	shapes = append(shapes, Shape{bwi.points[0].x, bwi.points[0].y, 0, 0, []Point{}})

	for _, p := range bwi.points {
		if belongsToShape(shapes[i], p) {
			shapes[i].points = append(shapes[i].points, p)
			width := p.x - shapes[i].x
			if (width > shapes[i].width) {
				shapes[i].width = width
			}
			height := p.y - shapes[i].y
			if (height > shapes[i].height) {
				shapes[i].height = height
			}
		} else {
			newImage.points = append(newImage.points, p)
		}
	}
	return findShapes(newImage, shapes)
}

func belongsToShape(shape Shape, point Point) bool {
	if len(shape.points) == 0 {
		return true
	}
	for _, p := range shape.points {
		if distance(p, point) < minDistanceBetweenShapes {
			return true
		}
	}
	return false
}

func distance(p1 Point, p2 Point) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
