package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {
	p := Person{
		Name:  "Tom",
		Age:   20,
		Email: "tom@qq.com",
	}
	b, _ := json.Marshal(p)
	fmt.Printf("b:%v\n", string(b))
	ToString()

}

func ToString() {
	b := []byte(`{"Name":"Tom","Age":20,"Email":"tom@qq.com"}`)
	var m Person
	err := json.Unmarshal(b, &m)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	fmt.Printf("b:%v\n", m)
}
