package main

import "fmt"

type myStruct struct {
	// 정수 필드
	intField int

	// 문자열 필드
	stringField string

	// 슬라이스 필드
	sliceField []int
}

// 3개의 매개변수를 받아서 myStruct 구조체를 초기화한 후 리턴하는 함수
//
func NewMyStruct(intField int, stringField string, sliceField []int) *myStruct {
	// 스택에 만들고 그 참조를 리턴
	// return &myStruct{intField, stringField, sliceField}

	imsi := new(myStruct)

	imsi.intField = intField
	imsi.stringField = stringField
	imsi.sliceField = sliceField

	return imsi
}

// Rect 구조체의 필드 값을 가져오는 메서드와 설정하는 메서드를 선언
type Rect struct {
	width  int
	height int
}

func (rect *Rect) GetWidth() int {
	return rect.width
}

func (rect *Rect) getHeight() int {
	return rect.height
}

func (rect *Rect) setWidth(width int) {
	rect.width = width
}

func (rect *Rect) setHeight(height int) {
	rect.height = height
}

// 메서드를 만들 때 *을 생략하고 리시버를 만들면 리시버는 호출하는 리시버의 데이터가 새로 만들어져서 대입됩니다.
func (rect Rect) copyHeight(height int) {
	rect.height = height
}

func main() {
	// 구조체 초기화
	var s1 = myStruct{
		intField:    1,
		stringField: "kabigon",
		sliceField:  []int{1, 2, 3},
	}
	fmt.Println(s1) // 전체를 출력

	// 필드 단위로 접근
	fmt.Println(s1.sliceField)

	// 필드 이름을 생략하고 초기화
	var s2 = myStruct{2, "잠만보", []int{1, 3, 5}}

	fmt.Println(s2) // 전체를 출력

	// 구조체 생성을 하고 나중에 초기화
	var s3 = myStruct{}
	s3.intField = 3
	fmt.Println(s3) // 전체를 출력

	// 구조체 포인터의 메모리 동적 할당
	ins := new(myStruct)
	fmt.Println(ins) // 이렇게 출력하면 앞에 &가 붙습니다.

	ins.intField = 20
	fmt.Println(ins)

	m := NewMyStruct(4, "Pokemon", []int{3, 2, 1})
	fmt.Println(m)

	// Rect 구조체와 연결된 메서드 사용
	rect := Rect{}
	rect.setWidth(10)
	rect.setHeight(20)
	fmt.Println(rect.GetWidth(), rect.getHeight())

	// copyHeight 메서드는 height를 매개변수의 값으로 변경하도록 되어 있지만
	// 리시버가 포인터 형태가 아니기 때문에 호출하는 리시버의 데이터를 변경할 수 없습니다.
	rect.copyHeight(30)
	fmt.Println(rect.GetWidth(), rect.getHeight())
}
