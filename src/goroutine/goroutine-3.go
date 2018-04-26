package main

import (
	"fmt"
	"sync"
)

// 协程 等待所有goroutine执行完成 退出main
func main() {
	var wg sync.WaitGroup		//WaitGroup等到所有的goroutine执行完成，并且阻塞主线程的执行，直到所有的goroutine执行完成
	wg.Add(3)				//goroutine数量
	go func(a, b int) {
		defer wg.Done()			//goroutine完成时减1
		fmt.Println(a + b)
	}(1, 1)
	go func(a, b int) {
		defer wg.Done()			//goroutine完成时减1
		fmt.Println(a + b)
	}(1, 2)
	go func(a, b int) {
		defer wg.Done()			//goroutine完成时减1
		fmt.Println(a + b)
	}(1, 3)
	wg.Wait()					//等待所有Add设置的goroutine执行完成
}