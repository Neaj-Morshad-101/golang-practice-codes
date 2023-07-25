package main

import "math"

func Distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}


func TotalArea(shapes ...Shape) float64 {
  var area float64
  for _, s := range shapes {
    area += s.Area()
  }
  return area
}