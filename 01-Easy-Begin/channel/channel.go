package main

import (
	"fmt"
	"time"
)

//CSP：通信顺序进程，是 Go 语言采用的并发同步模型，是一种形式语言，用来描述并发系统间进行交互的模式。

func firstChannel() {
	msg := make(chan string)
	go func() { msg <- "from goroutine" }()
	fmt.Println(<-msg)
}

func main() {
	firstChannel()
	IntChannel()

}

func IntChannel() {
	c := make(chan int) // 通道发送数据
	// goroute接受数据
	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()
	c <- 1
	c <- 2
	time.Sleep(time.Second * 3)
}
