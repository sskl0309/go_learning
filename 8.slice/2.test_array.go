package main

import "fmt"

func test(arr []int) {
	for i := 0; i < len(arr); i++ {
		arr[i] += 10
	}
}
func printArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println("arr[i] =", arr[i])
	}
}

func main() {
	arr := []int{1, 2, 3} //切片，动态数组，长度不固定
	printArray(arr)
	test(arr)
	printArray(arr)
	fmt.Printf("arr type = %T\n", arr)

    fmt.Println("--------------------------------")
	
	arr2 := make([]int, 3) //make([]type, length)
	var arr3 []int = make([]int, 3)
	copy(arr3, arr2)
	a2 := arr2[0:2]
	a2[0] = 100
	
	printArray(arr2)
	printArray(arr3)
}