package main

```
함수는 int의 타입 매개변수 및 반환활 타입을 지정한다.
동일한 매개변수는 하나만 명시하고 나머지는 생략 가능

함수 클로져
각각의 클로져는 자신만의 sum 변수를 갖는다.
```

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		ret := first
		first, second = second, first+second
		return ret
	}
}

func main() {
	fmt.Println(add(50, 100))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i), // 0, 1, 2, 3, 4, 5, 6 ... 쭉쭉 더한다.
			neg(-2*i),
		)
	}

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
