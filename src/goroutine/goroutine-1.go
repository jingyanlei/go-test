package main

import (
	"fmt"
	"time"
)

func Add(a, b int) {
	fmt.Println(a + b)
}

// 协程
func main() {
	//go关键字启动函数
	go Add(1,1)
	
	//goroutine方式运行
	go func(a, b int) {
		fmt.Println(a + b)
	}(1, 2)
	
	//main运行结束,goroutine还没有结束，需要sleep,显示打印内容
	time.Sleep(time.Second)
}
