package main

import "fmt"

func main() {
	// 정수 데이터가 저장된 공간의 참조를 저장하기 위한 포인터 변수를 생성
	var numPtr *int
	fmt.Println(numPtr)

	// 동적 메모리 할당 - heap의 공간을 확보해서 할당
	numPtr = new(int)
	// 메모리의 참조는 16진수 리턴 - 원래는 2진수 형태로 되어 있지만
	// 2진수로 표현하면 너무 길기 때문에 16진수로 표현
	fmt.Println(numPtr)
}
