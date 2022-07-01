package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		// GoRoutine을 추가
		wg.Add(1)
		go func(n int) {
			time.Sleep(1 * time.Second)
			fmt.Println(n)
			wg.Done()
		}(i)
	}

	// 모든 GoRoutine이 종료될 때까지 대기
	wg.Wait()
	fmt.Println("The End")
}
