package main

import "fmt"

func test1(a int, b int) int {
	return a + b
}

//多返回值，匿名返回值
func test2(a int, b int) (int, int) {
	return a + b, a - b
}

//多返回值，命名返回值
func test3(a int, b int) (c int, d int) {
	c = a + b
	d = a - b
	return
}


func main() {
	fmt.Println(test1(1, 2))
	fmt.Println(test2(1, 2))
	fmt.Println(test3(1, 2))
}