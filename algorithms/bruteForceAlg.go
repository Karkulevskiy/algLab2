package algorithms

import . "github.com/karkulevskiy/algLab2/models"

// Простой перебор прямоугольников, инкрементируем, если точка входит в границы прямоугольника\

type BruteForceAlg struct {
	Rectangles []Rectangle
}

func NewBruteForceAlg(rectangles []Rectangle) *BruteForceAlg {
	return &BruteForceAlg{
		Rectangles: rectangles,
	}
}

func (bf *BruteForceAlg) Test(p Point) int64 {
	totalRectangles := int64(0)
	for _, rect := range bf.Rectangles {
		if rect.IsInside(p) {
			totalRectangles++
		}
	}
	return totalRectangles
}
