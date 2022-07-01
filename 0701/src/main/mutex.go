package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // CPU의 모든 코어 개수만큼 Go에서 사용

	// 공유 데이터 생성
	var data = []int{} // 정수형 슬라이스 생성

	// 공유 자원을 동시에 사용할 수 없도록 해주는 Mutex 생성
	var mutex = new(sync.Mutex)

	// 아래의 동시성 함수들이 주석 처리된 부분보다 효용성이 높습니다.
	// 이는 주석 처리된 함수들의 경우, 반복문이 끝날 때까지 data 변수에 접근하지 못하는 데 반해,
	// 아래의 경우, data에 값을 추가하는 작업이 끝날 때마다 데이터에 접근할 수 있기 때문입니다.
	go func() {
		for i := 0; i < 1000; i++ {
			// Unlock()이 호출될 때까지 이 영역의 자원은 다른 곳에서 사용할 수 없음
			mutex.Lock()
			data = append(data, 1)
			mutex.Unlock()
		}
	}() // 익명 함수 바로 실행

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()
			data = append(data, 1)
			mutex.Unlock()
		}
	}() // 익명 함수 바로 실행

	// 	mutex.Lock()
	// 	for i := 0; i < 1000; i++ {
	// 		// Unlock()이 호출될 때까지 이 영역의 자원은 다른 곳에서 사용할 수 없음
	// 		data = append(data, 1)
	// 	}
	// 	mutex.Unlock()
	// }() // 익명 함수 바로 실행

	// go func() {
	// 	mutex.Lock()
	// 	for i := 0; i < 1000; i++ {
	// 		data = append(data, 1)
	// 	}
	// 	mutex.Unlock()
	// }() // 익명 함수 바로 실행

	time.Sleep(2 * time.Second)
	fmt.Println(len(data)) // 슬라이스의 길이
}
