package main

import (
	"fmt"
	"time"
)

// 缓冲channel,chan满时阻塞，打印日志可以看到chan满后，读取一条，写入一条，写入需要等待
func main()  {
	c := make(chan int, 9)
	
	//读取
	go func() {
		time.Sleep(time.Second)
		fmt.Println("chan:",len(c)+1)
		//方式一chan关闭后自动退出for
		for s := range c {
			//fmt.Println("chan:",len(c)+1)
			fmt.Println("读取chan " ,s)
			time.Sleep(time.Second)
		}
		fmt.Println("chan close")
		
		//方式二chan关闭后需要退出for循环，否则会造成死循环
		//for {
		//	if s, ok := <-c; ok {
		//		fmt.Println(s, ok)
		//	} else {
		//		fmt.Println("chan close", ok)
		//		break
		//	}
		//}
	}()
	
	//写入
	for i:=1; i<=30; i++ {
		fmt.Println("写入chan i:",i)
		c <- i
	}
	
	for {

	}
}