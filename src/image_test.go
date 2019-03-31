package main

import (
	"testing"
)

func TestConvertBlock1(t *testing.T) {

	bwImage := image2BlackWhiteImage("../images/blocks-1.png")

	if bwImage.points[0].x != 316 || bwImage.points[0].y != 248 {
		t.Log("Failed to convert single block image, first point is invalid")
		t.Fail()
	}

	if bwImage.width != 1416 || bwImage.height != 1096 {
		t.Log("Failed to convert single block image, image width/height invalid")
		t.Fail()
	}

	diagram := blackWhiteImage2Diagram(bwImage)

	if len(diagram.shapes) != 1 {
		t.Log("Failed to convert single block image to diagram")
		t.Fail()
	}
}

func TestConvertBlock2(t *testing.T) {

	bwImage := image2BlackWhiteImage("../images/blocks-2.png")

	if bwImage.points[0].x != 316 || bwImage.points[0].y != 248 {
		t.Log("Failed to convert block-2, first point is invalid")
		t.Fail()
	}

	diagram := blackWhiteImage2Diagram(bwImage)

	if len(diagram.shapes) != 2 {
		t.Log("Failed to convert a two block image to diagram")
		t.Fail()
	}
}
