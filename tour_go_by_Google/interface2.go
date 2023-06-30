package main

import "fmt"

// Speaker 인터페이스 정의
type Speaker interface {
	Speak()
}

// Person 타입 구현
type Person struct {
	Name string
}

// Person에 대한 Speak 메서드 구현
func (p Person) Speak() {
	fmt.Printf("안녕하세요, 제 이름은 %s입니다.\n", p.Name)
}

// Dog 타입 구현
type Dog struct {
	Name string
}

// Dog에 대한 Speak 메서드 구현
func (d Dog) Speak() {
	fmt.Printf("멍멍! 저는 %s입니다.\n", d.Name)
}

func main() {
	// Speaker 인터페이스 타입의 변수 선언
	var speaker Speaker

	// Person 인스턴스 할당
	person := Person{Name: "John"}
	// Person을 Speaker 인터페이스 타입으로 할당
	speaker = person
	speaker.Speak()

	// Dog 인스턴스 할당
	dog := Dog{Name: "Buddy"}
	// Dog를 Speaker 인터페이스 타입으로 할당
	speaker = dog
	speaker.Speak()
}
