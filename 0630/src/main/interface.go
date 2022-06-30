package main

import "fmt"

type Service interface {
	GetName() string
}

// ServiceImpl 구조체에서 호출할 수 있는 GetName 메서드를 생성
// ServiceImpl 구조체는 Service 인터페이스의 메서드를 만들었으므로, 이를 implements 되었다고 합니다.
func (s ServiceImpl) GetName() string {
	return s.name
}

type ServiceImpl struct {
	name string
}

// 인터페이스 생성 - Attack()이라는 메서드를 소유한 구조체는 Starcraft를 implement한 것이 됩니다.
// Starcraft 변수에 구조체를 대입할 수 있습니다.
type Starcraft interface {
	Attack()
}

type Protoss struct {
	unit string
}

func (p Protoss) Attack() {
	fmt.Println(p.unit, "공격")
}

type Terran struct {
	unit string
}

func (t Terran) Attack() {
	fmt.Println(t.unit, "공격")
}

type Zerg struct {
	unit string
}

func (z Zerg) Attack() {
	fmt.Println(z.unit, "공격")
}

// 패닉을 확인하는 함수
func checkPanic() {
	// 패닉이 발생하면 프로그램을 중단하지 않고 패닉 메시지를 출력
	if r := recover(); r != nil {
		fmt.Println("Panic Captured!!", r)
	}
}

// panic을 발생시키는 함수
func panicTest(p bool) {
	defer checkPanic()
	if p == true {
		panic("문제가 발생해서 종료합니다.")
	}
}

func main() {
	// 인터페이스는 변수만 선언할 수 있고 직접 메모리 할당을 할 수는 없습니다.
	// 즉, 인터페이스는 인스턴스를 생성할 수 없다라는 의미와 같습니다.
	var service Service
	fmt.Println(service)

	// 인터페이스 변수에 인터페이스를 Implements한 구조체의 데이터를 대입하는 것이 가능
	service = ServiceImpl{"Pokemon"}

	// 인터페이스 타입으로 선언한 변수를 이용해서 메서드를 호출하면 구조체에 구현한 메서드가 호출됩니다.
	fmt.Println(service.GetName())

	// 일반적인 구조체 생성을 통한 함수 호출
	// 같은 기능을 호출하기 위해 결과적으로 3개의 변수를 선언
	p := Protoss{"질럿"}
	p.Attack()

	t := Terran{"마린"}
	t.Attack()

	z := Zerg{"저글링"}
	z.Attack()

	// 인터페이스를 활용한 함수 호출
	// 인터페이스 변수만으로 함수 호출이 이루어짐 -> 다형성[Polymorphism]을 활용한 방식
	// 인터페이스 타입의 변수이므로 implement한 구조체를 대입받을 수 있습니다.
	var starcraft Starcraft

	starcraft = Protoss{"질럿"}

	// starcraft.Attack()는 대입된 구조체의 메서드를 호출합니다.
	starcraft.Attack()

	starcraft = Terran{"마린"}
	starcraft.Attack()

	starcraft = Zerg{"저글링"}
	starcraft.Attack()

	panicTest(true)

	// panic이 발생하면 이 문장은 수행되지 않고 프로그램이 중단됩니다.
	fmt.Println("프로그램의 모든 작업을 수행하고 종료합니다.")
}
