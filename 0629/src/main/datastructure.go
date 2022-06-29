package main

import "fmt"

func main() {
	// 배열 선언
	var ar1 [3]int
	fmt.Println(ar1)

	// 선언한 후 초기화
	var ar2 [3]string
	ar2 = [3]string{"카비곤", "메르시", "피카츄"}
	fmt.Println(ar2)

	// 선언과 동시에 초기화
	ar3 := [3]float32{20.5, 30.2, 21.9}
	fmt.Println(ar3)

	// 반복문을 이용한 ar2 배열의 모든 데이터 접근
	// len 함수를 4번 호출
	for idx := 0; idx < len(ar2); idx = idx + 1 {
		temp := ar2[idx]
		fmt.Println(temp)
	}

	// 아래의 코드가 직전 코드보다 효율성이 더 높습니다.
	// len 함수를 1번만 호출
	size := len(ar2)
	for idx := 0; idx < size; idx = idx + 1 {
		temp := ar2[idx]
		fmt.Println(temp)
	}

	// 배열 ar2의 0번째부터 2번째 앞 요소까지 출력
	fmt.Println(ar2[0:2])

	for idx, value := range ar2 {
		fmt.Printf("%d:%s\n", idx, value)
	}

	// 배열 복제
	br := ar2
	br[0] = "라이츄"
	fmt.Println(ar2)
	fmt.Println(br)
}
