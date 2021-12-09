package main

import "fmt"

// 在struct上实现方法封装
type MyInt struct {
	Number int
}

func (m MyInt) SayHello() {
	fmt.Println("Hello World!")
}

func (m *MyInt) SetNumber(other int) {
	m.Number = other
}

func (m MyInt) SayNumber() {
	fmt.Println(m.Number)
}

// Struct结构体嵌套实现继承
type ViewName struct {
	Name string
	ViewOther
}

func (v ViewName) SayName() {
	fmt.Println(v.Name)
}

type ViewOther struct {
	Value string
}

func (v ViewOther) SayValue() {
	fmt.Println(v.Value)
}

// 接口实现多态
type Controller interface {
	SayHello()
	SayNumber(int)
	SayHi()
}
type DefaultController struct {
}

func (d DefaultController) SayHello() {
	fmt.Println("Hello World!")
}
func (d DefaultController) SayNumber(number int) {
	fmt.Println(fmt.Sprintf("%d", number))
}

func (d DefaultController) SayHi() {
	fmt.Println("Say Hi")
}

func main() {
	var my MyInt
	my.Number = 1
	my.SayHello()
	my.SayNumber()
	my.SetNumber(3)
	my.SayNumber()

	var viewName ViewName
	viewName.Name = "weiwei"
	viewName.Value = "value"
	viewName.SayName()
	viewName.SayValue()

	var d DefaultController
	var c Controller

	c = d
	c.SayHello()
	c.SayNumber(123)
	c.SayHi()
}
