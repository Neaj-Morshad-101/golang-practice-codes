package main

import "fmt"

func swap(x, y int) (int, int) {
	temp := x
	x = y
	y = temp

	return x, y
}

func main() {
	x, y := swap(2, 3)
	fmt.Println(x, y);
}
