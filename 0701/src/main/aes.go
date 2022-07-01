package main

import (
	"crypto/aes"
	"fmt"
)

func main() {
	// 해시값을 추출하고자 하는 데이터를 생성
	s := "Hello World! 123"
	key := "Hello Key 123456"

	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		fmt.Println(err)
		return
	}

	// 암호화
	cipherText := make([]byte, len(s))
	block.Encrypt(cipherText, []byte(s))
	fmt.Printf("%x\n", cipherText)

	// 복호화
	plainText := make([]byte, len(s))
	block.Decrypt(plainText, cipherText)
	fmt.Println(string(plainText))
}
