package main

type Rectangle struct {
	leftPoint  Point
	rightPoint Point
}

func (r *Rectangle) isInside(p Point) bool {
	if p.x >= r.leftPoint.x && p.y >= r.leftPoint.y &&
		p.x <= r.rightPoint.x && p.y <= r.rightPoint.y {
		return true
	}
	return false
}
