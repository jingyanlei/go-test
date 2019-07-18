package main

import "fmt"

//47. defer 函数的参数值
//对 defer 延迟执行的函数，它的参数会在声明时候就会求出具体值，而不是在执行时才求值：

// 在 defer 函数中参数会提前求值
func main() {
	var i = 1
	defer fmt.Println("result: ", func() int { return i * 2 }())
	i++
}
//result: 2