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
