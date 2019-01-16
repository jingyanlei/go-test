package main

import (
	"fmt"
	"time"
)
// chan close 关闭chan后接收者，也会接收完chan里的所有数据
func main()  {
	dataChan := make(chan int, 10)
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 3)
	go func() { //用于演示接收操作
		fmt.Println("Received chan sleep 3s")
		time.Sleep(time.Second*3)
		<-syncChan1
		for {
			if elem, ok := <- dataChan; ok {
				fmt.Printf("Received [%d] [recvier]\n", elem)
			} else {
				fmt.Println("dataChan is close")
				break
			}
		}
		fmt.Println("Done.[recvier]")
		syncChan2 <- struct{}{}
	}()

	go func() { //用于演示写入操作
		for i:=0; i<10; i++{
			if dataChan != nil {}
			fmt.Println(len(dataChan))
			dataChan <- i
			fmt.Printf("Sent: %d [sender]\n", i)
		}
		close(dataChan)
		fmt.Println("sent dataChan close")
		syncChan1 <- struct{}{}
		fmt.Println("Done. [sender]")
		syncChan2 <- struct{}{}
	}()
	<-syncChan2
	<-syncChan2
}
