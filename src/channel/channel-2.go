package main

import (
	"fmt"
)

// 非缓冲channel,读取时阻塞，发送操作在接收者准备好之前是阻塞的
func main()  {
	c := make(chan int)						//创建方式一
	var c2 chan int = make(chan int)		//创建方式二
	
	go func() {
		//关闭 关闭chan后不可以写入，已写入chan可以读取到
		defer close(c)
		defer close(c2)
		//读取
		fmt.Println(<-c)
		fmt.Println(<-c2)
	}()

	//写入
	c <- 1
	c2 <- 2
}