package main

import "fmt"

func main() {
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")

	var x string
	x = "first "
	fmt.Println(x)
	x = x + "second"

	fmt.Println(x)

	fmt.Println(x == "first second")

	y := 10
	//fmt.Println(y);

}
