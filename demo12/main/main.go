package main

import (
    "fmt"
)

// 定义结构体
type Student struct {
    Name  string
    Age   int
    Score float64
}

// 值接收者方法
func (s Student) GetInfo() string {
    return fmt.Sprintf("姓名：%s，年龄：%d，成绩：%.2f", s.Name, s.Age, s.Score)
}

// 指针接收者方法
func (s *Student) UpdateScore(newScore float64) {
    s.Score = newScore
}

func main() {
    // 创建实例
    stu := Student{Name: "小明", Age: 18, Score: 92.5}
    
    // 调用方法
    fmt.Println(stu.GetInfo())  // 姓名：小明，年龄：18，成绩：92.50
    
    // 修改成绩
    stu.UpdateScore(95.0)
    fmt.Println(stu.GetInfo())  // 姓名：小明，年龄：18，成绩：95.00
}