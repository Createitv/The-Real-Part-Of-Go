package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Read(strChan chan string)  {
    data := <-strChan
    fmt.Println(data)
    wg.Done()
}

func main() {
    wg.Add(1)

    strChan := make(chan string)
    go Read(strChan)
    strChan <- "老苗"

    wg.Wait()
}