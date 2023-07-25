package main

import "math"

func (c *Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) Area() float64 {
	l := Distance(r.x1, r.y1, r.x1, r.y2)
	w := Distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}
