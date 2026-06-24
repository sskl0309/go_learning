package main

import (
	"fmt"
)

var (
	fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	//匿名函数
	res := func(n1 int, n2 int) int {
		return n1 + n2
	}(1, 3)
	fmt.Println(res)

	fun := func(n1 int, n2 int) int {
		return n1 - n2
	}
	fmt.Println(fun(1, 3))

	fmt.Println(fun1(1, 3))
}