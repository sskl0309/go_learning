package main

import (
	"fmt"
)


func main(){
	num := new(int)
	slice := make([]int, 2)
	a := 10
	nums := [...]int{1, 2, 3, 4, 5}
	num = &a
	slice = nums[0:3]
	fmt.Printf("%v\n", num)
	fmt.Printf("%v\n", slice)
}
