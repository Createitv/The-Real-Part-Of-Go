package main

import (
	"fmt"
	"strconv"
)

type String interface {
	String() string
}
type Human struct {
	name  string
	age   int
	phone string
}

func (h *Human) String() string {
	return "?" + h.name + "-" + strconv.Itoa(h.age) + "-" + h.phone
}
func main() {
	Bob := Human{name: "Bob",age: 30,phone: "23-213-213"}
	fmt.Println("This Human is:", Bob)
}
