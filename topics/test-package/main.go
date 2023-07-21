package main

import (
	"fmt"
	math1 "github.com/Neaj-Morshad-101/golang-book-codes/topics/test-package/math"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math1.Average(xs)
	fmt.Println(avg)
}
