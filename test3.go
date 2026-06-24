package main

import (
	"fmt"
)

func PrintHello() {
	fmt.Println("Hello, World!")
}

func Add(a, b int) int {
	return a + b
}

func div(a, b int) (int, string) {
	if b == 0 {
		return -1, "除数不能为0"
	}
	return a / b, ""

}
func main() {
	PrintHello()

	fmt.Printf("%d\n", Add(1, 2))

	result, errMsg := div(1, 0)
	if errMsg != "" {
		fmt.Printf("%s\n", errMsg)
	} else {
		fmt.Printf("%d\n", result)
	}
	
}
