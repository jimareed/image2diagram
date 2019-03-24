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
}
