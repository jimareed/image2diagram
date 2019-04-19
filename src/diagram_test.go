package main

import (
	"testing"
)

func TestEmptyDiagram(t *testing.T) {
	diagram := Diagram{}
	diagram.width = 0
	diagram.height = 0

	if len(diagram.shapes) != 0 {
		t.Log("Failed to create a diagram with empty blocks")
		t.Fail()
	}

	s := diagram2String(diagram)

	if s != "{"+
		"\"width\": 0,"+
		"\"height\": 0,"+
		"\"blocks\": [],"+
		"\"connectors\": []"+
		"}\n" {
		t.Log("Failed to convert an empty diagram to string")
		t.Fail()
	}
}

func TestSingleBlock(t *testing.T) {
	bwi := bwiTestOneBlock()

	diagram := blackWhiteImage2Diagram(bwi)

	if len(diagram.shapes) != 1 {
		t.Log("Failed to convert a single block, invalid block count")
		t.Fail()
	}
}

func TestTwoBlocks(t *testing.T) {
	bwi := bwiTestTwoBlocks()

	diagram := blackWhiteImage2Diagram(bwi)

	if len(diagram.shapes) != 2 {
		t.Log("Failed to convert two blocks, invalid block count")
		t.Fail()
	}

	if !(diagram.shapes[0].x == 2 && diagram.shapes[0].y == 2 &&
		diagram.shapes[1].x == 42 && diagram.shapes[1].y == 12) {
		t.Log("Failed to convert two blocks, invalid blocks")
		t.Fail()
	}
}
