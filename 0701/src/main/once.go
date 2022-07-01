package main

import (
	"fmt"
	"sync"
)

func Func() {
	fmt.Println("한번만 실행하는 함수")
}

func main() {
	// 함수를 한번만 실행하도록 해주는 구조체 생성
	once := new(sync.Once)

	for i := 0; i < 3; i++ {
		go func(n int) {
			fmt.Println("goRoutine :", n)
			once.Do(Func)
		}(i)
	}

	// 키보드 입력받는 함수 호출 - 키보드를 누를 때까지 대기
	fmt.Scanln()
}
