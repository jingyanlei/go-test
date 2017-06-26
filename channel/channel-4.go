package main

import (
	"fmt"
)

//只读chan
func readChan(c <-chan int)  {
	c <- 111
	for s := range c {
		fmt.Println("read", s)
	}
}

//只写chan
func writeChan(c chan<- int, i int)  {
	fmt.Println("write i:",i)
	c <- i
}

// channel 只读 只写
func main()  {
	c := make(chan int)
	
	//读取
	go func() {
		readChan(c)
	}()
	
	//写入
	for i:=1; i<=30; i++ {
		writeChan(c, i)
	}
	
	for {

	}
}