package main

import "fmt"

type Vertax struct {
	X int
	Y int
}

func main() {
	p := Vertax{1, 2}
	q := &p
	q.X = 1e9
	fmt.Println(q.X, q.Y)

	r := Vertax{Y: 99, X: 20}
	fmt.Println(r.X, r.Y)

	x := new(Vertax)
	fmt.Println(x, x.X, x.Y)
}
