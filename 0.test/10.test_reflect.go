package main

import (
	"fmt"
	"reflect"
)

type Book struct {
	Title  string
	Author string
}

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

func (b *Book) Read() {
	fmt.Println("Reading", b.Title)
}

func (b *Book) Write() {
	fmt.Println("Writing", b.Title)
}

func Print(arg interface{}) {
	fmt.Println(reflect.TypeOf(arg))
	fmt.Println(reflect.ValueOf(arg))
	t := reflect.TypeOf(arg).Elem()
	fmt.Println(t)
}

func main() {
	b := &Book{
		Title:  "Go in Action",
		Author: "Chen",
	}
	Print(b)
}
