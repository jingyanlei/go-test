package main

import (
	"fmt"
	"runtime"
)

func test4(s string)  {
	for i:=0; i<2; i++ {
		runtime.Gosched()
		fmt.Println(s, i)
	}
}

// 协程 出让时间片 出让时间片前每次第个输出都是b0，出让时间片后第一个输出可能会是a0
func main() {
	go test4("a")
	test4("b")
}