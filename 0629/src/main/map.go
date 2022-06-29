package main

import "fmt"

func main() {
	dictionary := map[string]string{"name": "kabigon", "1": "Snorlax", "2": "잠만보"}

	// 전체 값 출력
	fmt.Println(dictionary)

	// 하나의 값 출력
	fmt.Println(dictionary["name"])
	fmt.Println(dictionary["age"])

	// value에는 데이터가 존재하면 데이터가, 존재하지 않으면 기본값이 대입되며
	// ok에는 데이터의 존재여부가 bool로 대입됩니다.
	value, ok := dictionary["age"]

	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("데이터가 존재하지 않습니다.")
	}

	// 데이터 추가 및 업데이트
	dictionary["address"] = "서울시"
	fmt.Println(dictionary)

	// 데이터 삭제
	delete(dictionary, "address")
	fmt.Println(dictionary)

	// 데이터 전체 순회
	for key, value := range dictionary {
		fmt.Printf("%s:%s\n", key, value)
	}

	// 그룹별 데이터 저장 - 배열(슬라이스)을 배열(슬라이스)로
	// 배열은 인덱스를 이용해서 저장하고 map은 key를 이용해서 저장합니다.
	// 세부 데이터를 배열을 이용해서 저장하는 것보다는 Map을 이용해서 세부 데이터를 저장하고
	// 이들의 list를 이용하는 것이 MVC 구현에 유리
	// MVC : 데이터를 생성하는 부분과 출력하는 부분을 분리하고 이를 Controller를 이용해서 연결하는 방식

	// 팀별 야구선수 명단을 저장

	// 데이터 생성
	haitai := []string{"선동렬", "이종범", "최향남", "이대진", "김상진"}
	kia := []string{"양현종", "이범호", "김주찬", "최형우"}

	haitai = append(haitai, "김진우")

	kbo := map[string][]string{
		"해태": haitai,
		"기아": kia,
	}

	kbo["빙그레"] = []string{"이강돈", "이정훈"}

	// map은 레퍼런스 타입이므로 참조를 대입하면 참조가 복사됩니다.
	baseball := kbo
	baseball["빙그레"] = []string{"이상군", "한희민"}

	// 데이터 출력
	for key, value := range kbo {
		fmt.Printf("%s\n", key)
		for _, player := range value {
			fmt.Printf("%s\t", player)
		}
		fmt.Println()
	}
}
