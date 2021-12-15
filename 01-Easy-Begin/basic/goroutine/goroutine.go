package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func write(data chan<- int) {
	for i := 0; i < 10; i++ {
		data <- i
	}
	wg.Done()
	close(data)
}

func read(data <-chan int) {
	for {
		d, ok := <-data
		if !ok {
			break
		}

		fmt.Println(d)
	}
	wg.Done()
}

func selectChannel() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch1 <- 1
	ch2 <- 2
	select {
	case v, ok := <-ch1:
		if ok {
			fmt.Println("v1", v)
		}
	case v, ok := <-ch2:
		if ok {
			fmt.Println("v2", v)
		}
	default:
		fmt.Println("No channel userd")
	}
}

func main() {
	wg.Add(2)
	dataChan := make(chan int)
	go write(dataChan)
	go read(dataChan)

	wg.Wait()
	selectChannel()
}
