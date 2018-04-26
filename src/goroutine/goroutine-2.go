package main

import (
	"fmt"
	"time"
)

// 协程 无序执行
func main() {
	go func(a, b int) {
		fmt.Println(a + b)
	}(1, 1)
	go func(a, b int) {
		fmt.Println(a + b)
	}(1, 2)
	go func(a, b int) {
		fmt.Println(a + b)
	}(1, 3)
	//main运行结束,goroutine还没有结束，sleep,显示打印内容
	time.Sleep(time.Second)
}
