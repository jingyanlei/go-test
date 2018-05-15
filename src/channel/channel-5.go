package main

import (
	"fmt"
	"time"
)

// 返回default c1,c2未初始化
func chan5_1()  {
	var c1 chan int
	var c2 chan int
	select {
		case c1 <- 111:
			fmt.Println("111")
		case c2 <- 222:
			fmt.Println("222")
		default:
			fmt.Println("default")
	}
}

// 返回111 c1符合条件
func chan5_2()  {
	var c1 chan int = make(chan int, 1)
	var c2 chan int = make(chan int, 1)
	select {
	case c1 <- 111:
		fmt.Println("111")
	case c2 <- 222:
		fmt.Println("222")
	default:
		fmt.Println("default")
	}
}

// 循环取出c1中的数据，3秒后关闭chan退出循环-break退出select循环，但没有退出for循环，会造成死循环
func chan5_3()  {
	var c1 chan int = make(chan int, 10)
	for i:=0; i<10; i++ {
		c1<-1
	}
	//3秒后关闭chan
	go func() {
		time.Sleep(time.Second*3)
		close(c1)
	}()
	go func() {
		var sum int = 1
		for {
			fmt.Println("sum:", sum)
			sum++
			select {
			case e, ok := <- c1:
				if !ok {
					fmt.Println("chan close")
					break
				}
				fmt.Println(e)
			}
		}
	}()
}

// 循环取出c1中的数据，3秒后关闭chan退出循环-正确退出for循环
func chan5_4()  {
	var c1 chan int = make(chan int, 10)
	for i:=0; i<10; i++ {
		c1<-1
	}
	//3秒后关闭chan
	go func() {
		time.Sleep(time.Second*3)
		close(c1)
	}()
	go func() {
		var sum int = 1
		var e int
		var ok bool
		for {
			fmt.Println("sum:", sum)
			sum++
			select {
			case e, ok = <- c1:
				if !ok {
					fmt.Println("chan close")
					break
				}
				fmt.Println(e)
			}
			if !ok {
				fmt.Println("for end")
				break
			}
		}
	}()
}

// 循环取出c1中的数据，3秒后超时退出for循环
func chan5_5()  {
	var c1 chan int = make(chan int, 10)
	for i:=0; i<10; i++ {
		c1<-1
	}
	//3秒后关闭chan
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(time.Second*5)
		timeout <- false
	}()
	go func() {
		var sum int = 1
		var e int
		var ok bool
		for {
			fmt.Println("sum:", sum)
			sum++
			select {
				case e, ok = <- c1:
					if !ok {
						fmt.Println("chan close")
						break
					}
					fmt.Println(e)
				case ok = <-timeout:
					fmt.Println("timeout")
			}
			if !ok {
				fmt.Println("for end")
				break
			}
		}
	}()
}

// channel select 语句只能用于通道的接收和写入
func main()  {
	chan5_1()
	chan5_2()
	//chan5_3()
	chan5_4()
	chan5_5()
	
	for  {
		
		t1 := * time.NewTimer(time.Second * 3)
		if <-t1.C {
			fmt.Println("定时任务")
			t1.Reset(3 * time.Second)
		}
	}
	
	
	
	for {
	
	}
}