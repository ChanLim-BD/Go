package main

import "fmt"

// Person 구조체 정의
type Person struct {
	Name string
	Age  int
}

// 포인터 리시버를 사용한 함수
func (p *Person) SayHello() {
	fmt.Printf("안녕하세요, 저는 %s이고 %d살입니다.\n", p.Name, p.Age)
}

func main() {
	// 구조체 인스턴스 생성
	person := Person{
		Name: "John",
		Age:  25,
	}

	// 구조체의 메서드 호출
	person.SayHello()
}
