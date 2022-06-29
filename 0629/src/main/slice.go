package main

import "fmt"

func main() {
	fmt.Println("===========배열과 같은 점=============")
	// 3개의 데이터가 생성됨 - 기본값으로 초기화가 수행됨
	var ar [3]int

	// 참조를 저장할 준비만 됨
	var slice []int

	fmt.Println(ar)
	fmt.Println(slice)

	// 5개의 데이터를 저장할 수 있는 슬라이스를 생성해서 대입
	slice = make([]int, 5)
	fmt.Println(slice)

	a := []string{"Kabigon", "Cookie", "Lime"}
	fmt.Println(a[0])

	for _, value := range a {
		fmt.Println(value)
	}
	fmt.Println("=================================")

	fmt.Println("===========배열과 다른 점=============")
	xr := [3]string{"WaterMelon", "Lemon", "StrawBerry"}
	list := []string{"쿠키", "초코바", "아이스크림"}

	// 배열을 다른 배열에 대입 - 배열의 데이터가 복제가 되기 때문에 yr이 데이터를 수정해도 xr은 변경되지 않음
	// 메모리를 2배로 사용하지만 외부의 영향을 받지 않음
	yr := xr
	yr[2] = "Apple"
	fmt.Println(xr)

	// slice를 다른 slice에 대입 - 복제가 되지 않고 참조를 넘기기 때문에 ar을 변경하면 list에도 변경된 내용이 적용됨
	// 메모리를 적게 사용하지만 외부의 영향을 받음 - 가독성이 떨어질 수 있음
	al := list
	al[2] = "견과류"
	fmt.Println(list)

	// 3개의 데이터를 저장할 수 있는 slice 생성
	var li = make([]string, 3)

	// slice의 데이터를 li에 복제
	copy(li, list)
	li[0] = "캔디"
	fmt.Println(list)

	// slice에 데이터 추가
	list = append(list, "카라멜")
	list = append(list, "껌", "식료품")
	fmt.Println(list)

	st := []string{"GraphicCard", "CPU", "RAM"}
	list = append(list, st...)
	fmt.Println(list)
}
