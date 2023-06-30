package main

import "fmt"

type Vertax struct {
	X int
	Y int
}

func main() {
	m := make(map[string]Vertax)
	m["Data Universe"] = Vertax{99, -23}

	var n = map[string]Vertax{
		"Hello, World": Vertax{
			20, -10,
		},
		"Data Universe": Vertax{
			99, -23,
		},
	}

	fmt.Println(m)
	fmt.Println(n)

	n["Data Universe"] = Vertax{50, 50}
	fmt.Println(n)
	delete(n, "Data Universe")
	fmt.Println(n)
}
