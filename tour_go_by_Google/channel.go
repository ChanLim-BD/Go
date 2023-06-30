package main

import (
	"fmt"
)

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c) // a[0] a[1] a[2] : 7 2 8
	go sum(a[len(a)/2:], c) // a[3] a[4] a[5] : -9 4 0
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}
