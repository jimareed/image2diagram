package main

import (
	"testing"
)

func TestEmptyDiagram(t *testing.T) {
	diagram := Diagram{}
	diagram.width = 0
	diagram.height = 0

	if len(diagram.blocks) != 0 {
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
