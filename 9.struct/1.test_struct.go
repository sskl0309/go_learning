package main

import "fmt"

type Person struct {
	name string
	age  int
	sex  string
}

func (this Person) Show() {
	fmt.Println(this.name)
	fmt.Println(this.age)
	fmt.Println(this.sex)
}

func (this *Person) SetName(name string) {
	this.name = name
}

type Animal interface {
	Call()
}

type Cat struct {
	name string
}

func (this Cat) Call() {
	fmt.Println(this.name + " is calling")
}

type Dog struct {
	name string
}

func (this Dog) Call() {
	fmt.Println(this.name + " is calling")
}

func Call(animal Animal) {
	animal.Call()
}

func main() {
	person := Person{name: "zhansan", age: 10, sex: "nan"}
	person.Show()
	person.SetName("lisi")
	person.Show()

	cat := Cat{name: "kitty"}
	dog := Dog{name: "doggy"}
	Call(cat)
	Call(dog)
}
