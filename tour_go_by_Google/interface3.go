package main

import "fmt"

// Shape 인터페이스 정의
type Shape interface {
	Area() float64
}

// Circle 구조체 정의
type Circle struct {
	Radius float64
}

// Circle에 대한 Area 메서드 구현
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Rectangle 구조체 정의
type Rectangle struct {
	Width  float64
	Height float64
}

// Rectangle에 대한 Area 메서드 구현
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	// Shape 인터페이스 타입의 변수 선언
	var shape Shape

	// Circle 인스턴스 할당
	circle := Circle{Radius: 5}
	// Circle을 Shape 인터페이스 타입으로 할당
	shape = circle
	fmt.Printf("원의 넓이: %.2f\n", shape.Area())

	// Rectangle 인스턴스 할당
	rectangle := Rectangle{Width: 3, Height: 4}
	// Rectangle을 Shape 인터페이스 타입으로 할당
	shape = rectangle
	fmt.Printf("사각형의 넓이: %.2f\n", shape.Area())
}
