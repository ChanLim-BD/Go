package main

import (
	"fmt"
	"math"
)

type Vertax struct {
	X, Y float64
}

func (v *Vertax) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertax{3, 4}
	fmt.Println(v.Abs())
}
