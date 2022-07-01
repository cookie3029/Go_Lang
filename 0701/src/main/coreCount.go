package main

import (
	"fmt"
	"runtime"
	"time"
)

func runSomeLoop(n int) {
	for i := 0; i < n; i++ {
		fmt.Println("Printing:", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// Core 개수를 출력
	fmt.Println(runtime.NumCPU())

	// Go가 Core의 개수만큼 동시에 수행할 수 있는 개수를 설정
	runtime.GOMAXPROCS((runtime.NumCPU()))

	for i := 0; i < 100; i++ {
		go runSomeLoop(3)
	}

	time.Sleep(30 * time.Second)
}
