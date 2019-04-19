package main

import (
	"fmt"
)

// Block contains the x,y coordinates of the start of the block
type Block struct {
	x, y int
}

// Diagram contains a list of blocks (structure used by block-diagram-editor)
type Diagram struct {
	width  int
	height int
	shapes []Shape
}

func blackWhiteImage2Diagram(bwImage BlackWhiteImage) Diagram {

	diagram := Diagram{}
	diagram.width = bwImage.width
	diagram.height = bwImage.height

	diagram.shapes = blackWhiteImage2Shapes(bwImage)

	return diagram
}

func diagram2String(diagram Diagram) string {

	b := ""

	for _, s := range diagram.shapes {
		if len(b) > 0 {
			b += ","
		}
		b += fmt.Sprintf("{\"x\": %d, \"y\": %d, \"width\": %d, \"height\": %d}", s.x, s.y, s.width, s.height)
	}

	s := fmt.Sprintf(
		"{"+
			"\"width\": %d,"+
			"\"height\": %d,"+
			"\"blocks\": [%s],"+
			"\"connectors\": []"+
			"}\n",
		diagram.width, diagram.height, b)

	return s
}
