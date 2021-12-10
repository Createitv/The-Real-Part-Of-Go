package main

import "fmt"

func main() {

	/*
	var ch1 chan int // 引用类型必须make初始化
	// 缓冲区通道：驿站
	// 无缓存区通道： 送货上门
	// len, cap也对通道有效
	ch1 = make(chan int, 1) // 带缓冲区通道类型，给无缓冲区发送数据会造成通道阻塞
	ch1 <- 10 // 发送
	x := <-ch1 // 接受channel数据类型
	fmt.Println(x)
	close(ch1)

	*/

	ch1 := make(chan int, 100)
	ch2 := make(chan int, 200)
	go f1(ch1)
	go f2(ch1, ch2)
	for ret := range ch2 {
		fmt.Println(ret)
	}

}


/*
两个goroutine：
 1.生成1-100
 2.读ch1，计算平方
*/

func f1(ch chan<- int) {
	for i:= 1; i < 101; i++ {
		ch <- i
	}
	close(ch)
}

func f2(ch1 chan int, ch2 chan int) {
	for {
		tmp,ok := <-ch1
		if !ok {
			break
		}
		ch2 <- tmp * tmp
	}
	close(ch2)
}