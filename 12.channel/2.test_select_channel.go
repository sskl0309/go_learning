package main

import (
	"fmt"
)

// func main() {
// 	ch := make(chan int)
// 	quit := make(chan bool)

// 	go func() {
// 		for i := 0; i < 5; i++ {
// 			ch <- i // 发送数据到channel
// 		}
// 		close(ch)
// 		quit <- true
// 	}()

// loop:
// 	for {
// 		select {
// 		case i, ok := <-ch:
// 			if !ok {
// 				// ch 已关闭，跳过
// 				continue
// 			}
// 			fmt.Println("Received:", i)
// 		case <-quit:
// 			fmt.Println("Quit")
// 			break loop
// 		}
// 	}
// 	fmt.Println("Done")
// }

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i // 发送数据到channel
		}
		close(ch) // 必须 close，否则 range ch 会永远阻塞
	}()

	for i := range ch {
		fmt.Println("Received:", i)
	}
	fmt.Println("Done")

}
