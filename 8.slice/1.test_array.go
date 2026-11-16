package main

import "fmt"

func main() {
	var arr [10]int
	for i := 0; i < len(arr); i++ {
		fmt.Println("arr[i] =", arr[i])
		arr[i] = i
	}
	fmt.Println("arr =", arr)

	for index, value := range arr {
		fmt.Println("index =", index, "value =", value)
	}
}