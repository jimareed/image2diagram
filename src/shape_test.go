package main

import (
	"testing"
)

func TestEmptyShape(t *testing.T) {
	shape := Shape{}

	s := shape2String(shape)

	if s != "{"+
		"\"x\": 0, "+
		"\"y\": 0"+
		"}" {
		t.Log("Failed to convert an empty shape to string:" + s)
		t.Fail()
	}

	bwi := BlackWhiteImage{}

	shapes := blackWhiteImage2Shapes(bwi)

	if len(shapes) != 0 {
		t.Log("Failed to convert an empty bwImage to shapes")
		t.Fail()
	}
}

func TestOneShape(t *testing.T) {

	shapes := blackWhiteImage2Shapes(bwiTestOneBlock())

	if len(shapes) != 1 {
		t.Log("Failed to convert a single block bwImage to shapes")
		t.Fail()
		return
	}

	if shapes[0].x != 2 || shapes[0].y != 2 {
		t.Log("Failed to convert a single block bwImage to shapes, invalid coordinates")
		t.Fail()
	}

	if shapes[0].width != 2 || shapes[0].height != 2 {
		t.Log("Failed to convert a single block bwImage to shapes, invalid width and/or height")
		t.Fail()
	}

}

func TestTwoShapes(t *testing.T) {

	shapes := blackWhiteImage2Shapes(bwiTestTwoBlocks())

	if len(shapes) != 2 {
		t.Log("Failed to convert a two block bwImage to shapes")
		t.Fail()
		return
	}

	if shapes[0].x != 2 || shapes[0].y != 2 || shapes[1].x != 42 || shapes[1].y != 12 {
		t.Log("Failed to convert a two block bwImage to shapes, invalid coordinates")
		t.Fail()
	}

}
