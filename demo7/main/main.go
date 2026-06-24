package main

import (
	"fmt"
	"strings"
)

//闭包
func Fun() func(int) int {
	var n = 10
	return func(x int) int {
		n = n + x			
		return n
	}
} 

func makeSuffix(suffix string) func(string) string {
	return func (name string) string{
		if strings.HasSuffix(name, suffix) {
			return name
			}
	return name + suffix
	}
}
	


func main() {
	f := Fun()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
	fun := makeSuffix(".go")
	fmt.Println(fun("test"))
	fmt.Println(fun("test.go"))
}


