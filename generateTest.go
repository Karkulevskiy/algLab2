package main

import (
	"math/rand"
)

func generateTests() {

}

func generatePoints(countPoints int64) []Point {
	points := make([]Point, countPoints)
	for i := int64(0); i < countPoints; i++ {
		x := rand.Int63n(countPoints)
		y := rand.Int63n(countPoints)
		points = append(points, Point{x: x, y: y})
	}
	return points
}

func generateRectangles(countRectangles int64) []Rectangle {
	rectangles := make([]Rectangle, countRectangles)
	for i := int64(0); i < countRectangles; i++ {
		pointLeft := Point{
			x: 10 * i,
			y: 10 * i,
		}
		pointRight := Point{
			x: 10 * (2*countRectangles - i),
			y: 10 * (2*countRectangles - i),
		}
		rectangles = append(rectangles, Rectangle{
			leftPoint:  pointLeft,
			rightPoint: pointRight,
		})
	}
	return rectangles
}
