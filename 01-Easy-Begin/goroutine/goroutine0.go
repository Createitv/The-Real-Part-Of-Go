package main

import (
	"fmt"
	"sync"
)

// func hello(num int) {
// 	fmt.Println("hello", num)
// 	wg.Done() // 通知主函数任务结束，计数器-1，知道计算器为0程序退出
// }

func run() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1) // 技术牌登记+1
		//go hello(i) // 开启协程,造成异步的效果
		go func(i int) {
			fmt.Println("hello", i)
			wg.Done()
		}(i)
	}
}

func main() {
	// sync执行
	// hello()

	run()
	fmt.Println("hello main")
	wg.Wait() // 等待协程结束
	/*
		time.Sleep(time.Second * 1) // 不Sleep主程序结束，协程函数被关闭
	*/

}
