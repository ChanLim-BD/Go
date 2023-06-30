package main

import (
	"fmt"
	"sync"
)

func sum(a []int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go sum(a[:len(a)/2], c, &wg)
	go sum(a[len(a)/2:], c, &wg)

	go func() {
		wg.Wait()
		close(c)
	}()

	x, ok := <-c
	if ok {
		y := <-c
		fmt.Println(x, y, x+y)
	}
}
