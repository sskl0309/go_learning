package main

import "fmt"

//全局变量
var gA int = 10
var gB = 20
//方法四不支持全局变量声明，只能用于局部变量


func main() {
	//1.第一种
	var a int
	fmt.Println("a =", a)

	//2.第二种
	var b int = 10
	fmt.Println("b =", b)

	//3.第三种
	var c = 10
	fmt.Println("c =", c)

	//4.第四种
	d := 10
	fmt.Println("d =", d)

	var ga, gb = 10, 20
	fmt.Println("ga =", ga, "gb =", gb)
	
	var (
		gc int = 30
		gd string 
	)
	fmt.Println("gc =", gc, "gd =", gd)

}