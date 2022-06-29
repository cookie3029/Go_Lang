package main

import "fmt"

func main() {
	// // 이 함수 호출이 나중에 이루어집니다.
	// // main 함수의 수행이 종료되기 직전에 호출합니다.
	// defer lazy1() // 스택에 놓임
	// defer lazy2() // lazy1을 누르고 그 위에 올라감
	// fmt.Println("안녕하세요")

	// Open한 후 작업을 처리하고 정리
	Open()

	// 작업 처리
	fmt.Println("작업 수행")

	// 정리
	Close()
}

func Open() {
	fmt.Println("열어서 내용을 리턴")
}

func Close() {
	fmt.Println("정리")
}
