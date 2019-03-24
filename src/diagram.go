package main

import (
	"fmt"
)

type Block struct {
	x, y int
}

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
