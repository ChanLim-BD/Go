package main

import "fmt"

// Writer 인터페이스 정의
type Writer interface {
	WriteData(data string)
	Clear()
}

// ConsoleWriter 구조체 정의
type ConsoleWriter struct{}

// ConsoleWriter에 대한 WriteData 메서드 구현
func (cw ConsoleWriter) WriteData(data string) {
	fmt.Println("Writing data to console:", data)
}

// ConsoleWriter에 대한 Clear 메서드 구현
func (cw ConsoleWriter) Clear() {
	fmt.Println("Clearing console")
}

func main() {
	// Writer 인터페이스 타입의 변수 선언
	var writer Writer

	// ConsoleWriter 인스턴스 할당
	consoleWriter := ConsoleWriter{}
	// ConsoleWriter를 Writer 인터페이스 타입으로 할당
	writer = consoleWriter

	// Writer 인터페이스의 메서드 호출
	writer.WriteData("Hello, world!")
	writer.Clear()
}
