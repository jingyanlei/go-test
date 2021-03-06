package main

import (
	"fmt"
	"os"
	"os/signal"
)

// 监听全部信号
func main()  {
	//创建chan
	c := make(chan os.Signal)
	//监听所有信号
	signal.Notify(c)
	//阻塞直到有信号传入
	fmt.Println("启动")
	//阻塞直至有信号传入
	s := <-c
	fmt.Println("退出信号", s)
}