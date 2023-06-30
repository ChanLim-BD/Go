package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)

	sum2 := 1
	for sum2 < 1000 {
		fmt.Println(sum2)
		sum2 += sum2
	}
}
