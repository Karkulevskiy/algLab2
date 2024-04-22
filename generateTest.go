package main

import (
	"fmt"
	"math/rand"
	"time"

	. "github.com/karkulevskiy/algLab2/models"
)

func startTests() {

}

func createLogs(t time.Duration, algName string) {
	fmt.Println(algName, t)
}

func generatePoints(countPoints int64) []Point {
	points := make([]Point, countPoints)
	for i := int64(0); i < countPoints; i++ {
		x := rand.Int63n(countPoints)
		y := rand.Int63n(countPoints)
		points = append(points, Point{X: x, Y: y})
	}
	return points
}

func generateRectangles(countRectangles int64) []Rectangle {
	rectangles := make([]Rectangle, countRectangles)
	for i := int64(0); i < countRectangles; i++ {
		pointLeft := Point{
			X: 10 * i,
			Y: 10 * i,
		}
		pointRight := Point{
			X: 10 * (2*countRectangles - i),
			Y: 10 * (2*countRectangles - i),
		}
		rectangles = append(rectangles, Rectangle{
			LeftPoint:  pointLeft,
			RightPoint: pointRight,
		})
	}
	return rectangles
}
