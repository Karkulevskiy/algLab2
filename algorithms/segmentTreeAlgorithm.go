package algorithms

import (
	"slices"

	. "github.com/karkulevskiy/algLab2/models"
)

type PST struct {
	Rectangles        []Rectangle
	CompressedCoordsX []int64
	CompressedCoordsY []int64
}

type Node struct {
	Val   int64
	Left  *Node
	Right *Node
}

type Action struct {
	CompressedX int64
	TopY        int64
	BottomY     int64
	IsOpening   bool
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
	for i, rect := range pst.Rectangles {
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

	// Добавляем в слайсы из map'ов
	for i, v := range cordX {
		sortedX[i] = v
	}

	for i, v := range cordY {
		sortedY[i] = v
	}

	// Сортируем слайсы
	slices.Sort(sortedX)
	slices.Sort(sortedY)

	// Сохраняем сжатые координаты
	pst.CompressedCoordsX = sortedX
	pst.CompressedCoordsY = sortedY
}

func (pst *PST) createActions() {
	// Сжимаем координаты
	pst.compressCoords()

	actions := []Action{}

	for _, rect := range pst.Rectangles {
		openAction := Action{
			CompressedX: getPoint(pst.CompressedCoordsX, rect.LeftPoint.X),
			BottomY:     getPoint(pst.CompressedCoordsY, rect.LeftPoint.Y),
			TopY:        getPoint(pst.CompressedCoordsY, rect.RightPoint.Y+1),
			IsOpening:   true,
		}
		closeAction := Action{
			CompressedX: getPoint(pst.CompressedCoordsX, rect.RightPoint.X+1),
			BottomY:     getPoint(pst.CompressedCoordsY, rect.LeftPoint.Y),
			TopY:        getPoint(pst.CompressedCoordsY, rect.RightPoint.Y+1),
			IsOpening:   false,
		}
		actions = append(actions, openAction, closeAction)
	}
	// TODO: мб нужно будет урезать слайс, чтобы не было TL

	slices.SortFunc(actions, func(first Action, second Action) int {
		return int(first.CompressedX - second.CompressedX)
	})

}

func getPoint(arr []int64, target int64) int64 {
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
