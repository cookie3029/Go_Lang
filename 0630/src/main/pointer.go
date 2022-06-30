package main

import "fmt"

// Call By Value 함수 - 매개변수가 참조형이 아닌 경우
func callByValue(n int) {
	n = n + 1
	fmt.Println("n :", n)
}

// Call By Reference 함수 - 매개변수가 참조형(포인터)인 경우
// 참조를 넘겼는데 데이터의 변경이 안되는 경우는 해당 함수가 데이터 변경을 안하는 경우인데
// 이런 경우는 return문이 반드시 존재합니다.
func callByReference(p *int) {
	*p = *p + 1
	fmt.Println("*p :", *p)
}

// 리시버가 메서드 안에서 사용이 안될 때는 _로 이름을 만드는 것이 가능
func (_ *Rect) display() {
	fmt.Println("I Have Width And Height")
}

func main() {
	var numPtr *int
	fmt.Println(numPtr) // 현재 상태는 nil

	// 동적 메모리 공간을 할당한 후 값을 대입하고 값을 출력
	numPtr = new(int)
	*numPtr = 30
	fmt.Println(numPtr, ":", *numPtr) // 30이 대입된 상태

	// 다른 메모리 영역의 주소를 가리키기
	x := 123
	numPtr = &x
	fmt.Println(numPtr, ":", *numPtr) // 123이 대입된 상태

	// 더블 포인터
	doublePointer := &numPtr
	fmt.Println(doublePointer, ":", *doublePointer)

	x = 1
	// x는 value 타입이기에 함수의 매개변수로 대입해도 값이 변경되지 않습니다.
	callByValue(x)
	fmt.Println("x :", x)

	// x의 참조를 넘기기 때문에 데이터가 변경될 수 있습니다.
	callByReference(&x)
	fmt.Println("x :", x)

}
