package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 0; i < 30; i++ {
		fmt.Println("A:", i)
	}
}
func b() {
	for i := 0; i < 30; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	// runtime.GOMAXPROCS(8) //开启CPU核心数，不然a(), b()两个函数相当于单线程
	// fmt.Println("bar:", runtime.NumGoroutine()) //GOMAXPROCS查看核心数，这里的cpu数量都是指逻辑cpu，即核数，而非物理cpu
	fmt.Println("Start:", runtime.NumGoroutine()) //goroutine程数
	runtime.GOMAXPROCS(6)
	fmt.Println("End:", runtime.NumCPU()) //8
	go a()                                //相当于异步效果
	for i := 0; i < 10; i++ {
		go b()
	}
	fmt.Println("End:", runtime.NumGoroutine())
	time.Sleep(time.Second)
}
