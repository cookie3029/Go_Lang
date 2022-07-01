package main

import (
	"fmt"
	"os"
)

func main() {
	var num1 int
	var num2 float32
	var str string

	file1, _ := os.Open("Kabigons_Test.txt")
	defer file1.Close()
	fmt.Fscanln(file1, &num1, &num2, &str)

	fmt.Println(num1, num2, str)

	file2, _ := os.Open("Kabigons_Test2.txt")
	defer file2.Close()
	fmt.Fscanf(file2, "%d, %f, %s", &num1, &num2, &str)

	fmt.Println(num1, num2, str)
}
