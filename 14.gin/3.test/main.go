package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Age  int    `gorm:"not null"`
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
	u1 := &User{
		Name: "LJY",
		Age:  18,
	}
	db.Create(u1)
	var u2 User
	db.First(&u2, 1)
	fmt.Println(u2)
	db.Model(&u2).Update("name", "nailong")
	fmt.Println(u2)
	db.Delete(&u2)
}
