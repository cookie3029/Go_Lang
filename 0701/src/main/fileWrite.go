package main

import (
	"fmt"
	"os"
)

func main() {
	file1, _ := os.Create("Kabigons_Test.txt")
	defer file1.Close()
	fmt.Fprintln(file1, 1, 1.7, "This_Is_Written_By_Kabigon")

	file2, _ := os.Create("Kabigons_Test2.txt")
	defer file1.Close()
	fmt.Fprintf(file2, "%d, %.2f, %s", 3, 1.3, "This_Is_Formatted_Text")
}
