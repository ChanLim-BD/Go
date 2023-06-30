package main

import "fmt"

func main() {
	p := []int{2, 3, 6, 1, 2, 6, 20}
	fmt.Println(p[1:5]) // 1 ~ 4

	q := make([]int, 5)
	r := make([]int, 0, 5)

	fmt.Println(len(q), cap(q), len(r), cap(r), q, r)

	r = append(r, 100)

	fmt.Println(len(q), cap(q), len(r), cap(r), q, r)

	var z []int
	fmt.Println(z)
	if z == nil {
		fmt.Println(true)
	}
}
