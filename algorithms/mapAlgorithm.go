package algorithms

import (
	"slices"

	. "github.com/karkulevskiy/algLab2/models"
)

// MapAlg описывает 2 алг
type MapAlg struct {
	CordX []int64   // Сжатые координаты по X
	CordY []int64   // Сжатые координаты по Y
	Map   [][]int64 // Построенная карта
}

// Конструктор типа MapAlg
func NewMapAlg(rectanles []Rectangle) *MapAlg {
	return fillCords(rectanles)
}

// Testing считает сколько прямоугольников попадает в точку
// P.S. (ну или наборот, я не знаю как правильно сказать)
func (m *MapAlg) MapTesting(p Point) int64 {
	if p.X < m.CordX[0] || p.Y < m.CordY[0] {
		return 0
	}
	x, y := getPrev(m.CordX, p.X), getPrev(m.CordY, p.Y)
	return m.Map[x][y]
}

// fillCords заполняет нашу map'у
func fillCords(rectangles []Rectangle) *MapAlg {
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

	mapAlg := &MapAlg{
		CordX: sortedX,
		CordY: sortedY,
		Map:   resMap,
	}
	return mapAlg
}

// getPrev ищет границу следующей координаты
func getPrev(arr []int64, target int64) int64 {
	left := int64(0)
	right := int64(len(arr) - 1)
	for left <= right {
		middle := int64((right + left) / 2)
		if arr[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left - 1
}
