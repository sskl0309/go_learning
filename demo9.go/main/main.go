package main

import (
	"fmt"
)

func test() {
	defer func(){
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
		}()
	num1 := 10
	num2 := 0
	fmt.Println("正常代码")
	num3 := num1 / num2
	fmt.Println(num3)
	fmt.Println("正常代码")
	
}

func readConfig(name string)error {
	if(name == "") {
		return fmt.Errorf("name is empty")
	}
	return nil
}

func test2() {
	err := readConfig("")
	if err != nil {
		panic(err)
	}
	fmt.Println("正常代码")
}


func main() {
	// test()
	// fmt.Println("正常代码")
	test2()
}