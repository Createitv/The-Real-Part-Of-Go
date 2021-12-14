package main

import (
	"errors"
	"fmt"
)

type Example struct {
	Content string
}

func main() {
	data := Example{Content: "show"}
	fmt.Printf("%v\n", data)
	fmt.Printf("%v\n", errors.New("我错了"))
	fmt.Printf("%+v\n", data)
	fmt.Printf("%#v\n", data)
	fmt.Printf("%T\n", data) //类型
	fmt.Printf("%t\n", true)
	fmt.Printf("%f\n", 10.2)
	fmt.Printf("%.3f\n", 10.2)

	fmt.Printf("%s|%s\n", "talk", []byte("is cheaper")) // 字符[]byte
	fmt.Printf("%p\n", &data) // 指针
	fmt.Printf("%#p\n", &data) // 指针
}