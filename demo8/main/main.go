package main

import (
	"fmt"
)

func main() {
    defer fmt.Println("1")
    defer fmt.Println("2")
    defer fmt.Println("3")
    fmt.Println("正常代码")
    return
}
// 输出:
// 正常代码
// 3
// 2
// 1

//1.清理资源
// f, err := os.Open("file.txt")
// if err != nil {
//     return
// }
// defer f.Close()  // 保证文件一定被关闭
// // ... 不管中间 return 还是 panic，都会执行 Close


//2. 配对操作（加锁/解锁）
// mu.Lock()
// defer mu.Unlock()
// // ... 临界区代码

//3. 记录函数执行时间
// start := time.Now()
// ... 函数体
//defer func() {
//     fmt.Println("耗时:", time.Since(start))
// }()