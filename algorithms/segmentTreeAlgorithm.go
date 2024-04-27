package algorithms

import (
	"slices"

	. "github.com/karkulevskiy/algLab2/models"
)

type PST struct {
	Rectangles             []Rectangle
	CompressedCoordsX      []int64
	CompressedCoordsY      []int64
	Roots                  []*Node
	CompressedRootsIndexes []int64
}

type Node struct {
	Val   int64
	Left  *Node
	Right *Node
}

type Action struct {
	CompressedIndexesX int64
	TopY               int64
	BottomY            int64
	IsOpening          bool
}

func NewPSTAlg(rectangles []Rectangle) *PST {
	return &PST{
		Rectangles: rectangles,
	}
}

func newCopyNode(node *Node) *Node {
	return &Node{
		Val:   node.Val,
		Left:  node.Left,
		Right: node.Right,
	}
}

func add(root *Node, left, right, start, end, value int64) *Node {
	if right <= start || left >= end {
		return root
	}

	if start <= left && right <= end {
		node := newCopyNode(root)
		node.Val += value
		return node
	}

	var middle int64 = (left + right) / 2
	node := newCopyNode(root)

	if node.Left == nil {
		node.Left = &Node{}
	}
	node.Left = add(node.Left, left, middle, start, end, value)

	if node.Right == nil {
		node.Right = &Node{}
	}
	node.Right = add(node.Right, middle, right, start, end, value)

	return node
}

func getTotalCover(root *Node, left, right, target int64) int64 {
	if right-left == 1 {
		return root.Val
	}

	middle := (left + right) / 2

	if target < middle {
		if root.Left == nil {
			return root.Val
		}
		return root.Val + getTotalCover(root.Left, left, middle, target)
	} else {
		if root.Right == nil {
			return root.Val
		}
		return root.Val + getTotalCover(root.Right, middle, right, target)
	}
}

func (pst *PST) compressCoords() {
	cordX := map[int64]int64{}
	cordY := map[int64]int64{}

	// Добавляем точки в map'ы

	for _, rect := range pst.Rectangles {
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

	// Сохраняем сжатые координаты
	pst.CompressedCoordsX = sortedX
	pst.CompressedCoordsY = sortedY
}

func (pst *PST) CreateActions() {
	// Сжимаем координаты
	pst.compressCoords()

	actions := []Action{}

	for _, rect := range pst.Rectangles {
		openAction := Action{
			CompressedIndexesX: getLower(pst.CompressedCoordsX, rect.LeftPoint.X),
			BottomY:            getLower(pst.CompressedCoordsY, rect.LeftPoint.Y),
			TopY:               getLower(pst.CompressedCoordsY, rect.RightPoint.Y),
			IsOpening:          true,
		}
		closeAction := Action{
			CompressedIndexesX: getLower(pst.CompressedCoordsX, rect.RightPoint.X),
			BottomY:            getLower(pst.CompressedCoordsY, rect.LeftPoint.Y),
			TopY:               getLower(pst.CompressedCoordsY, rect.RightPoint.Y),
			IsOpening:          false,
		}
		actions = append(actions, openAction, closeAction)
	}

	slices.SortFunc(actions, func(first Action, second Action) int {
		return int(first.CompressedIndexesX - second.CompressedIndexesX)
	})

	root := &Node{}
	prevCompressedIndexX := actions[0].CompressedIndexesX

	for _, action := range actions {
		if action.CompressedIndexesX != prevCompressedIndexX {
			pst.Roots = append(pst.Roots, root)
			pst.CompressedRootsIndexes = append(pst.CompressedRootsIndexes, prevCompressedIndexX)
			prevCompressedIndexX = action.CompressedIndexesX
		}

		if action.IsOpening {
			root = add(root, 0, int64(len(pst.CompressedCoordsY)), action.BottomY, action.TopY, 1)
		} else {
			root = add(root, 0, int64(len(pst.CompressedCoordsY)), action.BottomY, action.TopY, -1)
		}

	}

	pst.CompressedRootsIndexes = append(pst.CompressedRootsIndexes, prevCompressedIndexX)
	pst.Roots = append(pst.Roots, root)
}

func (pst *PST) PSTTesting(p Point) int64 {
	x := getLower(pst.CompressedCoordsX, p.X)
	y := getLower(pst.CompressedCoordsY, p.Y)

	if x == -1 || y == -1 {
		return 0
	}

	root := pst.Roots[getLower(pst.CompressedRootsIndexes, x)]

	res := getTotalCover(root, 0, int64(len(pst.CompressedCoordsY)), y)
	return res

}
