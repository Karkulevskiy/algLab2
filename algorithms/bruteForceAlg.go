package algorithms

import . "github.com/karkulevskiy/algLab2/models"

// Простой перебор прямоугольников, инкрементируем, если точка входит в границы прямоугольника
func BruteForceTesting(rectangles []Rectangle, p Point) int64 {
	totalRectangles := int64(0)
	for _, rect := range rectangles {
		if rect.IsInside(p) {
			totalRectangles++
		}
	}
	return totalRectangles
}
