package main

import (
	"crypto/sha512"
	"fmt"
)

func main() {
	// 해시값을 추출하고자 하는 데이터를 생성
	s := "LimeMojito"
	h1 := sha512.Sum512([]byte(s))
	fmt.Printf("%x\n", h1)

	sha := sha512.New()
	sha.Write([]byte("Lime"))
	sha.Write([]byte("Mojito"))
	h2 := sha.Sum(nil)
	fmt.Printf("%x\n", h2)
}
