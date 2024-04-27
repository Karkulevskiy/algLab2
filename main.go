package main

import (
	"fmt"
	"math"
	"os"
	"time"

	. "github.com/karkulevskiy/algLab2/algorithms"
	. "github.com/karkulevskiy/algLab2/models"
)

const (
	N                    = 1000 // Максимальное количество прямоугольников и точек
	Repeats              = 3    //  Количество повторов алгоритма
	BruteForceAlgorithm  = "BruteForceAlgorithm"
	MapAlgorithm         = "MapAlgorithm"
	SegmentTreeAlgorithm = "SegmentTreeAlgorithm"
	FileName             = "results.txt"
)

func main() {
	TestBruteForceAlg()
	TestMapAlg()
	TestSegmentTreeAlg()
}

func TestSegmentTreeAlg() {
	fmt.Printf("\n\n\n")
	writeData("Segment Tree Algorithm testing: QUERYTIME | BUILDTIME | R, P count")
	for i := int64(1); i < N; i += 100 {
		rectangles := generateRectangles(i)
		points := generatePoints(i)

		var avgTime time.Duration
		var avgBuildTime time.Duration

		// Будем тест проводить по 3 раза и брать среднее время
		for totalRepeats := 0; totalRepeats < Repeats; totalRepeats++ {

			sgtAlg := NewPSTAlg(rectangles)

			//Считаем время построения
			start := time.Now()
			sgtAlg.CreateActions()
			elapsed := time.Since(start)

			avgBuildTime += elapsed

			// Считаем время запрос
			start = time.Now()
			for _, point := range points {
				_ = sgtAlg.PSTTesting(point)
			}
			elapsed = time.Since(start)

			avgTime += elapsed
		}

		avgTime /= Repeats
		avgBuildTime /= Repeats
		writeData(fmt.Sprintf("%v | %v | %v", avgTime, avgBuildTime, int64(len(rectangles))))
	}
	fmt.Printf("\n\n\n")
	writeData("\n\n\n")
}

func TestMapAlg() {
	fmt.Printf("\n\n\n")
	writeData("Map Algorithm testing: QUERYTIME | BUILDTIME | R, P count")
	for i := int64(1); i < N; i += 35 {
		rectangles := generateRectangles(i)
		points := generatePoints(i)

		var avgTime time.Duration
		var avgBuildTime time.Duration

		// Будем тест проводить по 3 раза и брать среднее время
		for totalRepeats := 0; totalRepeats < Repeats; totalRepeats++ {

			//Считаем время построения
			start := time.Now()
			mapAlg := NewMapAlg(rectangles)
			elapsed := time.Since(start)

			avgBuildTime += elapsed

			// Считаем время запрос
			start = time.Now()

			for _, point := range points {
				_ = mapAlg.Test(point)
			}
			elapsed = time.Since(start)
			avgTime += elapsed
		}

		avgTime /= Repeats
		avgBuildTime /= Repeats

		// Записываем логи
		writeData(fmt.Sprintf("%v | %v | %v", avgTime, avgBuildTime, int64(len(rectangles))))
	}
	fmt.Printf("\n\n\n")
	writeData("\n\n\n")
}

func TestBruteForceAlg() {
	fmt.Printf("\n\n\n")
	writeData("BruteForce Algorithm testing")
	for i := int64(1); i < N; i += 35 {
		rectangles := generateRectangles(i)
		points := generatePoints(i)

		// Здесь я не учитываю время построения прямоугольников, потому что оно происходит за O(1)
		bruteForceAlg := NewBruteForceAlg(rectangles)

		var avgTime time.Duration

		// Будем тест проводить по 3 раза и брать среднее время
		for totalRepeats := 0; totalRepeats < Repeats; totalRepeats++ {
			start := time.Now()
			for _, point := range points {
				_ = bruteForceAlg.Test(point)
			}
			elapsed := time.Since(start)
			avgTime += elapsed
		}

		avgTime /= Repeats

		// Записываем логи
		writeData(fmt.Sprintf("%v | %v", avgTime, int64(len(rectangles))))
	}
	fmt.Printf("\n\n\n")
	writeData("\n\n\n")
}

func writeData(data string) {
	fmt.Println(data)
	file, err := os.OpenFile(FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err := file.WriteString(data + "\n"); err != nil {
		panic(err)
	}
}

// Генерация точек
func generatePoints(countPoints int64) []Point {
	points := []Point{}
	for i := int64(0); i < countPoints; i++ {
		x := int64(math.Pow(float64(11161*i), 31)) % (20 * countPoints)
		y := int64(math.Pow(float64(10501*i), 31)) % (20 * countPoints)
		points = append(points, Point{X: x, Y: y})
	}
	return points
}

// Генерация прямоугольников
func generateRectangles(countRectangles int64) []Rectangle {
	rectangles := []Rectangle{}
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
