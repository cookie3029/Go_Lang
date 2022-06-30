package main

import "fmt"

type super struct {
	num  int
	name string
}

// 슈퍼 구조체에 연결된 메소드
func (_ *super) method() {
	fmt.Println("super의 메소드")
}

type embedding struct {
	// 구조체 안에 다른 구조체 데이터가 필드로 존재하면 has a 관계
	s     super
	socre int
}

type sub struct {
	// 구조체 타입 이름만 사용하면 is a 관계 - 객체 지향 언어의 상속과 유사
	super
	score int
}

type person struct {
}

// person에 연결된 display
func (p *person) display() {
	fmt.Println("Person's Display Method!")
}

type student struct {
	p person
}

// student에 연결된 display
func (s *student) display() {
	// 포함된 구조체의 메서드 호출
	s.p.display()
	fmt.Println("Student's Display Method!")
}

func main() {
	e := new(embedding)
	// has a 관계일 때는 포함된 구조체 필드를 통해서 내부 구조체의 메서드 호출
	e.s.method()

	child := new(sub)
	child.method()

	s := new(student)
	s.display()
}
