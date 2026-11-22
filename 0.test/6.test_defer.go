package main

import "fmt"

func test1() {
	fmt.Println("test1")
}

func test2() {
	fmt.Println("test2")
}

func test3() {
	fmt.Println("test3")
}

//defer 延迟执行，在函数结束前执行
func main() {
	defer fmt.Println("defer main")
	fmt.Println("main")
	defer test1()
	defer test2()
	defer test3()
}