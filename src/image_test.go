package main

import (
	"testing"
)

func TestBlock1(t *testing.T) {
	diagram := Diagram{}
	diagram.width = 0
	diagram.height = 0

	if len(diagram.blocks) != 0 {
		t.Log("Failed to create a diagram with empty blocks")
		t.Fail()
	}
}
