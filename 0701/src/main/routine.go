package main

import (
	"fmt"
	"time"
)

func runSomeLoop(n int) {
	for i := 0; i < n; i++ {
		fmt.Println("Printing:", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// 다른 함수와 동시에 실행되도록 앞에 go를 추가
	go runSomeLoop(10)
	go runSomeLoop(10)

	// 동시성으로 실행하는 경우에는 main이 종료되지 않아야 합니다.
	// main 함수가 종료되지 않도록 하기 위해서 15초 대기
	time.Sleep(15 * time.Second)
}
