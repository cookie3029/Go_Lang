package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	doc := `
		{
			"name":"kabigon",
			"age":24
		}
	`
	// 파싱한 내용 저장하기 - 내용을 map으로 생성
	var data map[string]interface{}
	json.Unmarshal([]byte(doc), &data)

	fmt.Println(data["name"], data["age"])
}
