package main

import (
	"fmt"
	"sync"
)

// 并发锁，内存共享

var (
	x  int
	wg sync.WaitGroup
	lock sync.Mutex
)

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock() //内存加锁--互斥锁，保证只有一个goroutine访问资源
		x += 1
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
