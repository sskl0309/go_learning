package main

import (
	"fmt"
)

func Fun(arg interface{}) {
	value, ok := arg.(int)
	if ok {
		fmt.Println("arg is int")
		fmt.Println("value =", value)
	} else {
		fmt.Println("arg is not int")
	}
}

func main() {
	Fun(10)
	Fun("hello")
	Fun(true)
}
