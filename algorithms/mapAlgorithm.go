package algorithms

import (
	"slices"

	. "github.com/karkulevskiy/algLab2/models"
)

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
func (m *MapAlg) Test(p Point) int64 {
	x, y := getLower(m.CordX, p.X), getLower(m.CordY, p.Y)
	if x == -1 || y == -1 {
		return 0
	}
	return m.Map[x][y]
}

// fillCords заполняет нашу map'у
func fillCords(rectangles []Rectangle) *MapAlg {
	cordX := map[int64]int64{}
	cordY := map[int64]int64{}

	// Добавляем точки в map'ы
	for _, rect := range rectangles {
		cordX[rect.LeftPoint.X] = rect.LeftPoint.X
		cordX[rect.RightPoint.X] = rect.RightPoint.X
		cordY[rect.LeftPoint.Y] = rect.LeftPoint.Y
		cordY[rect.RightPoint.Y] = rect.RightPoint.Y
	}

	// Задаем слайсы для отсортированных данных
	sortedX := make([]int64, len(cordX))
	sortedY := make([]int64, len(cordY))

	// Добавляем в слайсы из map'ов
	i := 0
	for k := range cordX {
		sortedX[i] = k
		i++
	}

	i = 0
	for k := range cordX {
		sortedY[i] = k
		i++
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
		downBoundY := getLower(sortedY, rect.LeftPoint.Y)
		upBoundY := getLower(sortedY, rect.RightPoint.Y)
		for y := downBoundY; y < upBoundY; y++ {
			// Находим границы для поиска по X координате
			downBoundX := getLower(sortedX, rect.LeftPoint.X)
			upBoundX := getLower(sortedX, rect.RightPoint.X)
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

// getLower ищет границу следующей координаты
func getLower(arr []int64, target int64) int64 {
	start := int64(0)
	end := int64(len(arr))
	step := int64(0)
	for end > 0 {
		cur := start
		step = end / 2
		cur += step
		if target >= arr[cur] {
			start = cur + 1
			end -= step + 1
		} else {
			end = step
		}
	}
	return start - 1
}
