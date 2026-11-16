package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		i := 0
		for {
			fmt.Println("goroutine i:", i)
			i++
			time.Sleep(time.Second)
		}
	}()
	i := 0
	for {
		fmt.Println("main goroutine i:", i)
		i++
		time.Sleep(time.Second)
	}
}
