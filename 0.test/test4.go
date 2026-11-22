package main

import (
	"fmt"
)

func main(){
	nums := []int{1, 2, 3, 4, 5}
	for _, num := range nums {
		fmt.Printf("%d\n", num)
	}
	for i := 0; i < len(nums); i++ {
		fmt.Printf("%d\n", nums[i])
	}
	
	fmt.Printf("%v\n", nums)
}
