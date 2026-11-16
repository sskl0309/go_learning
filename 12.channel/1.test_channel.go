package main

import (
	"fmt"
)

func main() {
	//创建一个整型channel
	ch := make(chan int)

	//启动一个goroutine，向channel发送数据
	go func() {
		for i := 0; ; i++ {
			ch <- i // 发送数据到channel
		}
	}()

	//在主goroutine中从channel接收数据
	for i := range ch {
		fmt.Println("Received:", i)
		if i >= 5 {
			break // 接收5个数据后退出循环
		}
	}
}

