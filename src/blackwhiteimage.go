package main

// Point is a single point
type Point struct {
	x, y int
}

// BlackWhiteImage contains a black and white version of the image
type BlackWhiteImage struct {
	width  int
	height int
	points []Point
}

func blackWhiteImage2Diagram(bwImage BlackWhiteImage) Diagram {

	diagram := Diagram{}
	diagram.width = bwImage.width
	diagram.height = bwImage.height

	if len(bwImage.points) > 0 {
		diagram.blocks = append(diagram.blocks, Block{bwImage.points[0].x, bwImage.points[0].y})
	}

	return diagram
}
