package main

import (
	"testing"
)

func TestGetRegion(t *testing.T) {

	bwi := bwiTestOneBlock()
	point := Point{2, 2}

	region := getRegion(bwi, point)

	if !(region.x1 == 2 && region.y1 == 2 && region.x2 == 4 && region.y2 == 4) {
		t.Log("Failed to get region for single block")
		t.Fail()
	}

	bwi = bwiTestTwoBlocks()

	region = getRegion(bwi, point)

	if !(region.x1 == 2 && region.y1 == 2 && region.x2 == 4 && region.y2 == 4) {
		t.Log("Failed to get region for two blocks")
		t.Fail()
	}

}

func TestClearRegion(t *testing.T) {

	bwi := bwiTestOneBlock()
	point := Point{2, 2}

	region := getRegion(bwi, point)

	bwi = clearRegion(bwi, region)

	if len(bwi.points) > 0 {
		t.Log("Failed to clear region for single block")
		t.Fail()
	}

	bwi = bwiTestTwoBlocks()
	point = Point{2, 2}

	region = getRegion(bwi, point)

	bwi = clearRegion(bwi, region)

	if len(bwi.points) == 0 {
		t.Log("Failed to clear region for two blocks, cleared too much")
		t.Fail()
	}

	if bwi.points[0].x != 12 && bwi.points[0].y != 12 {
		t.Log("Failed to clear region for two blocks")
		t.Fail()
	}

}

func bwiTestOneBlock() BlackWhiteImage {
	bwi := BlackWhiteImage{}
	bwi.width = 300
	bwi.height = 200

	bwi.points = append(bwi.points, Point{2, 2})
	bwi.points = append(bwi.points, Point{3, 2})
	bwi.points = append(bwi.points, Point{4, 2})
	bwi.points = append(bwi.points, Point{2, 3})
	bwi.points = append(bwi.points, Point{4, 3})
	bwi.points = append(bwi.points, Point{2, 4})
	bwi.points = append(bwi.points, Point{3, 4})
	bwi.points = append(bwi.points, Point{4, 4})

	return bwi
}

func bwiTestTwoBlocks() BlackWhiteImage {
	bwi := BlackWhiteImage{}
	bwi.width = 300
	bwi.height = 200

	bwi.points = append(bwi.points, Point{2, 2})
	bwi.points = append(bwi.points, Point{3, 2})
	bwi.points = append(bwi.points, Point{4, 2})
	bwi.points = append(bwi.points, Point{2, 3})
	bwi.points = append(bwi.points, Point{4, 3})
	bwi.points = append(bwi.points, Point{2, 4})
	bwi.points = append(bwi.points, Point{3, 4})
	bwi.points = append(bwi.points, Point{4, 4})

	bwi.points = append(bwi.points, Point{42, 12})
	bwi.points = append(bwi.points, Point{43, 12})
	bwi.points = append(bwi.points, Point{44, 12})
	bwi.points = append(bwi.points, Point{42, 13})
	bwi.points = append(bwi.points, Point{44, 13})
	bwi.points = append(bwi.points, Point{42, 14})
	bwi.points = append(bwi.points, Point{43, 14})
	bwi.points = append(bwi.points, Point{44, 14})

	return bwi
}
