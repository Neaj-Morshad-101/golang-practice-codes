package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
		i++
	}

	for ii := 1; ii <= 10; ii++ {
		fmt.Println(ii)
	}

}
