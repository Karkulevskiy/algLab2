package algorithms

import . "github.com/karkulevskiy/algLab2/models"

// Простой перебор прямоугольников, инкрементируем, если точка входит в границы прямоугольника
func TestingBruteForce(rectangles []Rectangle, p Point) {
	totalRectangles := int64(0)
	for _, rect := range rectangles {
		if rect.IsInside(p) {
			totalRectangles++
		}
	}
}