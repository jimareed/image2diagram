package main

// Point is a single point
type Point struct {
	x, y int
}

// Region identifies a rectangular area
type Region struct {
	x1, y1, x2, y2 int
}

// BlackWhiteImage contains a black and white version of the image
type BlackWhiteImage struct {
	width  int
	height int
	points []Point
}

func getRegion(bwi BlackWhiteImage, point Point) Region {

	region := Region{0, 0, 0, 0}
	started := false
	finishedX := false

	for _, p := range bwi.points {
		if !started {
			started = p.x == point.x && p.y == point.x
			region.x1 = p.x
			region.y1 = p.y
			region.x2 = p.x
			region.y2 = p.y
		} else {
			if !finishedX {
				if p.y == region.y1 && p.x == region.x2+1 {
					region.x2 = p.x
				} else {
					finishedX = true
				}
			} else {
				if p.x == region.x2 && p.y == region.y2+1 {
					region.y2 = p.y
				}
			}
		}
	}

	return region
}

func copyImage(bwImage BlackWhiteImage) BlackWhiteImage {

	bwi := BlackWhiteImage{}
	bwi.width = bwImage.width
	bwi.height = bwImage.height

	for _, point := range bwImage.points {
		bwi.points = append(bwi.points, Point{point.x, point.y})
	}

	return bwi
}

func clearRegion(bwImage BlackWhiteImage, region Region) BlackWhiteImage {

	bwi := BlackWhiteImage{}
	bwi.width = bwImage.width
	bwi.height = bwImage.height

	for _, point := range bwImage.points {
		if point.x >= region.x1 && point.x <= region.x2 &&
			point.y >= region.y1 && point.y <= region.y2 {
			// don't want this point
		} else {
			bwi.points = append(bwi.points, Point{point.x, point.y})
		}
	}

	return bwi
}
