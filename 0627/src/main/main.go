package main

import "fmt"

func main() {
	var x int = 20
	fmt.Println(x)

	// 선언과 동시에 초기화 하는 경우는 자료형 생략이 가능
	var y = 30
	fmt.Println(y)

	// 서로 다른 자료형의 데이터를 동시에 초기화
	var i, f = 10, 20.7
	fmt.Printf("x=%d y=%f\n", i, f)

	// var str = "Hello World"
	str := "Hello World"
	fmt.Printf("str=%s", str)
}
