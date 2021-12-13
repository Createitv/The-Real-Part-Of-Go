package main //必须

import "fmt"

func test(x int) {
	result := 100 / x

	fmt.Println("result = ", result)
}

func main() {

	defer fmt.Println("aaaaaaaaaaaaaa")

	defer fmt.Println("bbbbbbbbbbbbbb")

	//调用一个函数，导致内存出问题
	defer test(10)

	defer fmt.Println("ccccccccccccccc")
}
