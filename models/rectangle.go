package models

type Rectangle struct {
	LeftPoint  Point
	RightPoint Point
}

func (r *Rectangle) IsInside(p Point) bool {
	if p.X >= r.LeftPoint.X && p.Y >= r.LeftPoint.Y &&
		p.X <= r.RightPoint.X && p.Y <= r.RightPoint.Y {
		return true
	}
	return false
}
