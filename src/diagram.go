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
	blocks []Block
}

func diagram2String(diagram Diagram) string {

	b := ""
	if len(diagram.blocks) > 0 {
		b = fmt.Sprintf("{\"x\": %d, \"y\": %d}", diagram.blocks[0].x, diagram.blocks[0].y)
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
