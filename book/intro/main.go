package main

// this is a comment

import (
	"fmt"
	"github.com/Neaj-Morshad-101/golang-book-codes/book/intro/mathpkg"
)

func main() {
	//fmt.Println("Hello World")
	//fmt.Println("1 + 1 =", 1+1)
	//fmt.Println("1 + 1 =", 1.0+1.0)
	//fmt.Println(len("Hello World"))
	//fmt.Println("Hello World"[1])               // 101 instead of e
	//fmt.Printf("char : %v\n", "Hello World"[1]) // 101 instead of e
	//fmt.Println("Hello " + "World")

	//var x string
	//x = "first "
	//fmt.Println(x)
	//x = x + "second"
	//fmt.Println(x)

	//fmt.Print("Enter a number: ")
	//var input float64
	//fmt.Scanf("%f", &input)
	//
	//output := input * 2
	//
	//fmt.Println(output)

	//loop :

	//for i := 1; i <= 10; i++ {
	//	if i%2 == 0 {
	//		fmt.Println(i, "even")
	//	} else {
	//		fmt.Println(i, "odd")
	//	}
	//
	//}

	//array and slice:

	//var arr [3]int
	//arr[1] = 100
	//fmt.Println(arr)

	//slice1 := []int{1, 2, 3}
	//slice2 := append(slice1, 4, 5)
	//fmt.Println(slice1, slice2)

	//slice1 := []int{1, 2, 3}
	//slice2 := make([]int, 2)
	//copy(slice2, slice1)
	//fmt.Println(slice1, slice2)

	//x := make(map[string]int)
	//x["key"] = 100
	//x["neaj"] = 101
	//fmt.Println(x)

	//mapFunction()

	//c := Circle{0, 0, 5}
	////fmt.Println(c.Area())
	//
	//r := Rectangle{0, 0, 10, 10}
	////fmt.Println(r.Area())
	//
	//fmt.Println(TotalArea(&c, &r))
	//
	//for i := 0; i < 10; i++ {
	//	go f(i)
	//}
	//var input string
	//fmt.Scanln(&input)

	//var c chan string = make(chan string)
	//
	//go pinger(c)
	//go ponger(c)
	//go printer(c)
	//
	//var input string
	//fmt.Scanln(&input)

	//
	//c1 := make(chan string)
	//c2 := make(chan string)
	//
	//go func() {
	//	for {
	//		c1 <- "from 1"
	//		time.Sleep(time.Second * 2)
	//	}
	//}()
	//
	//go func() {
	//	for {
	//		c2 <- "from 2"
	//		time.Sleep(time.Second * 3)
	//	}
	//}()
	//
	//go func() {
	//	for {
	//		select {
	//		case msg1 := <- c1:
	//			fmt.Println("Message 1", msg1)
	//			time.Sleep(time.Second * 2)
	//		case msg2 := <- c2:
	//			fmt.Println("Message 2", msg2)
	//			time.Sleep(time.Second * 2)
	//		case <- time.After(time.Second):
	//			fmt.Println("timeout")
	//			time.Sleep(time.Second * 2)
	//		default:
	//			fmt.Println("nothing ready")
	//			time.Sleep(time.Second * 2)
	//		}
	//	}
	//}()
	//
	//var input string
	//fmt.Scanln(&input)
	xs := []float64{1, 2, 3, 4}
	avg := mathpkg.Average(xs)
	fmt.Println(avg)

	mapFunction()
}
