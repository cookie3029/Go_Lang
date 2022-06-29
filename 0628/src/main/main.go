package main

import "fmt"

// 함수 선언
func display() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("함수 종료")
}

func main() {
	// 일정한 패턴의 함수 호출
	// function call
	display()
	fmt.Println("동일한 코드를 사용하는 중간에 추가된 코드")
	display()
}
