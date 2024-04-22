package algorithms

import (
	"slices"

	. "github.com/karkulevskiy/algLab2/models"
)

// Алгоритм на карте
func TestingMap(rectangles []Rectangle, p Point) {
	resMap := fillCords(rectangles)
	//TODO: Доделать поиск точек на map'е
}

// fillCords заполняет нашу map'у
func fillCords(rectangles []Rectangle) [][]int64 {
	cordX := map[int64]int64{}
	cordY := map[int64]int64{}

	// Добавляем точки в map'ы
	for i, rect := range rectangles {
		cordX[int64(i*3)] = rect.LeftPoint.X
		cordX[int64(i*3)+1] = rect.RightPoint.X
		cordX[int64(i*3)+2] = rect.RightPoint.X + 1
		cordY[int64(i*3)] = rect.LeftPoint.Y
		cordY[int64(i*3)+1] = rect.RightPoint.Y
		cordY[int64(i*3)+2] = rect.RightPoint.Y + 1
	}

	// Задаем слайсы для отсортированных данных
	sortedX := make([]int64, len(cordX))
	sortedY := make([]int64, len(cordY))

	for i, v := range cordX {
		sortedX[i] = v
	}

	for i, v := range cordY {
		sortedY[i] = v
	}

	// Сортируем слайсы
	slices.Sort(sortedX)
	slices.Sort(sortedY)

	// Задаем нашу map'у
	resMap := make([][]int64, len(sortedY))
	for i := range len(sortedX) {
		resMap[i] = make([]int64, len(sortedX))
	}

	// Заполняем нашу map'у
	for _, rect := range rectangles {
		// Находим границы для поиска по Y координате
		downBoundY := getPrev(sortedY, rect.LeftPoint.Y)
		upBoundY := getPrev(sortedY, rect.RightPoint.Y)
		for y := downBoundY; y < upBoundY; y++ {
			// Находим границы для поиска по X координате
			downBoundX := getPrev(sortedX, rect.LeftPoint.X)
			upBoundX := getPrev(sortedX, rect.RightPoint.X)
			for x := downBoundX; x < upBoundX; x++ {
				resMap[y][x]++
			}
		}
	}

	return resMap
}

// getPrev ищет границу следующей координаты
func getPrev(coords []int64, target int64) int64 {
	for i := range int64(len(coords)) {
		if coords[i] > target {
			return i
		}
	}
	return int64(len(coords) - 1)
}
