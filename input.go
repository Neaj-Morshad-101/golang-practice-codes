package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var x float64
	fmt.Scanf("%f", &x)

	outputx := x * 2

	fmt.Println(outputx)
}
