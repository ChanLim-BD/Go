package main

import "fmt"

func main() {
	var pow = []int{1, 2, 3, 4, 5}
	for i, v := range pow {
		fmt.Printf("2 ** %d = %d\n", i, v)
	}

	for _, x := range pow {
		fmt.Printf("%d\n", x)
	}
}
