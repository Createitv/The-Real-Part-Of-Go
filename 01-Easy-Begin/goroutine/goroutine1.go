package main

import (
	"fmt"
	"sync"
	"time"
)

func CounterUnSafe() {

	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("counter = %d", counter) // 线程无锁共享内存不安全，产生竞争关系

}
func CounterSafe() {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock() // 加锁保证线程安全
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("counter = %d", counter)
}

func WaitGroup() {
	var mut sync.Mutex
	var wg sync.WaitGroup // waitgruop等待
	count := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			count++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("counter = %d", count)
}

func goFirstLance() {
	for i := 0; i <= 10; i++ {
		// 携程
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(100 * time.Millisecond)
}

func main() {
	goFirstLance()
	CounterUnSafe()
	CounterSafe()
	WaitGroup()
}
