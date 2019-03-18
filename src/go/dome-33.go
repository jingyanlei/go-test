package main

import (
	"fmt"
	"time"
)

//33. 向已关闭的 channel 发送数据会造成 panic
//从已关闭的 channel 接收数据是安全的：
//
//接收状态值 ok 是 false 时表明 channel 中已没有数据可以接收了。类似的，从有缓冲的 channel 中接收数据，缓存的数据获取完再没有数据可取时，状态值也是 false
//
//向已关闭的 channel 中发送数据会造成 panic：

/*
func main() {
	ch := make(chan int)
	for i := 0; i < 3; i++ {
		go func(idx int) {
			ch <- idx
		}(i)
	}

	fmt.Println(<-ch)        // 输出第一个发送的值
	close(ch)            // 不能关闭，还有其他的 sender
	time.Sleep(2 * time.Second)    // 模拟做其他的操作
}
*/

//针对上边有 bug 的这个例子，可使用一个废弃 channel done 来告诉剩余的 goroutine 无需再向 ch 发送数据。此时 <- done 的结果是 {}：


func main() {
	ch := make(chan int)
	done := make(chan struct{})

	for i := 0; i < 3; i++ {
		go func(idx int) {
			select {
			case ch <- (idx + 1) * 2:
				fmt.Println(idx, "Send result")
			case <-done:
				fmt.Println(idx, "Exiting")
			}
		}(i)
	}

	fmt.Println("Result: ", <-ch)
	close(done)
	time.Sleep(3 * time.Second)
}