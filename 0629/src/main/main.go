package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"math/rand"
// 	"time"
// )

// func main() {
// 	F1()
// 	Add(10, 30)

// 	// 가변 매개변수를 가진 함수 호출
// 	VarArgs(10, 20)
// 	VarArgs(10, 20, 30)

// 	result := ReturnFunc(100, 200)
// 	fmt.Println("결과 : ", result)

// 	// 2개의 데이터를 리턴하는 함수 호출
// 	result1, result2 := MultipleReturn(10, 20)
// 	fmt.Println(result1, result2)

// 	// 2개의 데이터를 리턴하는 함수 호출
// 	result3, result4 := MultipleReturn(10, 20)
// 	fmt.Println(result3, result4)

// 	result5 := RecursiveSum(7)
// 	fmt.Println(result5)

// 	// result6 := Fibonacci(100)
// 	// fmt.Println(result6)

// 	result7 := FibonacciNoRecursive(100)
// 	fmt.Println(result7)

// 	// 함수를 변수에 저장
// 	f := F1
// 	f()

// 	F2()

// 	// 익명 함수 생성과 호출
// 	r := func(n int) int {
// 		sum := 0
// 		for idx := 0; idx <= n; idx = idx + 1 {
// 			sum += idx
// 		}
// 		return sum
// 	}(5)

// 	fmt.Println(r)

// 	// 파일의 내용을 읽어오기
// 	// 첫번재 리턴되는 내용은 데이터를 읽은 경우 데이터
// 	// 두번째 리턴되는 내용은 에러가 발생하지 않으면 nil이 리턴되고 에러가 발생하면 에러 내용이 전달됩니다.
// 	b, err := ioutil.ReadFile("./kabigon.txt")
// 	if err == nil {
// 		fmt.Printf("content:%s\n", b)
// 	}

// 	if b, err := ioutil.ReadFile("./kabigon.txt"); err == nil {
// 		fmt.Printf("contents%s\n", b)
// 	}

// 	// switch 표현식에서 함수 호출
// 	rand.Seed(time.Now().UnixNano())

// 	switch i := rand.Intn(10); {
// 	case i >= 1 && i < 6:
// 		fmt.Println("작은 숫자")
// 	default:
// 		fmt.Println("큰 숫자")
// 	}

// 	// outer 함수에 만들어진 변수 접근
// 	// localVar = localVar + 1
// 	// fmt.Println(localVar)

// 	// Outer가 리턴한 함수를 대입
// 	closer := Outer()
// 	closer()
// 	closer()
// }

// var F2 = func() {
// 	fmt.Println("함수")
// }

// func F1() {
// 	fmt.Println("함수")
// }

// // 매개변수의 자료형이 모두 일치한다면 자료형은 1번만 기재해도 됩니다.
// func Add(a, b int) {
// 	fmt.Println(a + b)
// }

// // 가변 인자 사용
// func VarArgs(ar ...int) {
// 	total := 0
// 	for _, value := range ar {
// 		total = total + value
// 	}
// 	fmt.Println("합게 :", total)
// }

// // 리턴이 있는 함수
// func ReturnFunc(x, y int) int {
// 	return x + y
// }

// // 여러 개의 데이터를 리턴하는 함수
// // 정수 2개를 리턴하는 함수
// // 2개 정도까지는 이 방식이 유용하지만 서로 다른 종류의 데이터 여러 개를 이런 방식으로 리턴하면 가독성이 떨어집니다.
// func MultipleReturn(a, b int) (int, int) {
// 	return a + b, a - b
// }

// // NamedReturn : 리턴할 데이터의 이름을 리턴하는 방식
// func NamedReturn(a, b int) (add, sub int) {
// 	add = a + b
// 	sub = a - b
// 	return
// }

// // 1부터 n까지의 합을 리턴하는 함수 - n + n-1까지의 함
// func RecursiveSum(n int) int {
// 	if n == 1 {
// 		return 1
// 	}
// 	return n + RecursiveSum(n-1)
// }

// // 피보나치 수열
// // 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89..
// func Fibonacci(n int) int {
// 	if n == 1 || n == 2 {
// 		return 1
// 	}
// 	return Fibonacci(n-1) + Fibonacci(n-2)
// }

// func FibonacciNoRecursive(n int) int {
// 	a := 1
// 	b := 1
// 	result := 1

// 	for i := 3; i <= n; i = i + 1 {
// 		result = a + b
// 		a = b
// 		b = result
// 	}

// 	return result
// }

// func Outer() func() {
// 	localVar := 1

// 	// 리턴하는 함수 - 이 함수는 Outer 안에 존재하기 때문에 localVar을 사용하는 것이 가능
// 	return func() {
// 		localVar = localVar + 1
// 		fmt.Println("localVar:", localVar)
// 	}
// }
