package main

import (
	"testing"
)

func TestEmptyImage(t *testing.T) {
	bwImage := BlackWhiteImage{}
	bwImage.width = 300
	bwImage.height = 200

	diagram := blackWhiteImage2Diagram(bwImage)

	if len(diagram.blocks) != 0 {
		t.Log("Failed to convert an empty image")
		t.Fail()
	}
}

func TestSingleBlock(t *testing.T) {
	bwImage := BlackWhiteImage{}
	bwImage.width = 300
	bwImage.height = 200

	bwImage.points = append(bwImage.points, Point{2, 2})
	bwImage.points = append(bwImage.points, Point{3, 2})
	bwImage.points = append(bwImage.points, Point{4, 2})
	bwImage.points = append(bwImage.points, Point{2, 3})
	bwImage.points = append(bwImage.points, Point{4, 3})
	bwImage.points = append(bwImage.points, Point{2, 4})
	bwImage.points = append(bwImage.points, Point{3, 4})
	bwImage.points = append(bwImage.points, Point{4, 4})

	diagram := blackWhiteImage2Diagram(bwImage)

	if len(diagram.blocks) != 1 {
		t.Log("Failed to convert a single block, invalid block count")
		t.Fail()
	}
}
