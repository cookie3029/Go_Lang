package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // CPU의 모든 코어 개수만큼 Go에서 사용

	var mutex = new(sync.Mutex)
	var cond = sync.NewCond(mutex)

	c := make(chan bool, 3)

	for i := 0; i < 3; i++ {
		go func(n int) {
			mutex.Lock()
			c <- true

			fmt.Println("wait begin:", n)
			cond.Wait() // 대기
			fmt.Println("wait end:", n)

			mutex.Unlock() // 잠금 상태 해제
		}(i)
	}

	for i := 0; i < 3; i++ {
		<-c
	}

	// 대기 중인 GoRoutine 깨우는 작업
	for i := 0; i < 3; i++ {
		mutex.Lock()

		fmt.Println("signal:", i)
		cond.Signal()

		mutex.Unlock()
	}

	// cond.Broadcast()
}
