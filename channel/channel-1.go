package main

import (
	"fmt"
	"time"
)

// 非缓冲channel,读取时阻塞
// 发送操作在接收者准备好之前是阻塞的
func main()  {
	c := make(chan int)						//创建方式一
	var c2 chan int = make(chan int)		//创建方式二
	
	go func() {
		//五秒后写入chan，main主程序一直阻塞，直到读到到数据为止
		time.Sleep(time.Second*5)
		//关闭 关闭chan后不可以写入，已写入chan可以读取到
		defer close(c)
		defer close(c2)
		//写入
		c <- 1
		c2 <- 2
	}()

	//读取
	fmt.Println(<-c)
	fmt.Println(<-c2)
}