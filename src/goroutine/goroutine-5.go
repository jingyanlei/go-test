package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}
func main() {
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")
}

// https://blog.csdn.net/think2me/article/details/79610838

// 在上面程序中的第 21 行和第 22 行，启动了两个 Go 协程。现在，这两个协程并发地运行。numbers 协程首先休眠 250 微秒，接着打印 1，然后再次休眠，打印 2，依此类推，一直到打印 5 结束。alphabete 协程同样打印从 a 到 e 的字母，并且每次有 400 微秒的休眠时间。 Go 主协程启动了 numbers 和 alphabete 两个 Go 协程，休眠了 3000 微秒后终止程序。

// 该程序会输出：

// 1 a 2 3 b 4 c 5 d e main terminated
