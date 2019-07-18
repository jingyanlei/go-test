package main

import "fmt"

//39. 从 panic 中恢复
//在一个 defer 延迟执行的函数中调用 recover() ，它便能捕捉 / 中断 panic

/*
// 错误的 recover 调用示例
func main() {
	recover()    // 什么都不会捕捉
	panic("not good")    // 发生 panic，主程序退出
	recover()    // 不会被执行
	println("ok")
}
*/

/*
// 正确的 recover 调用示例
func main() {
	defer func() {
		fmt.Println("recovered: ", recover())
	}()
	panic("not good")
}
*/

//从上边可以看出，recover() 仅在 defer 执行的函数中调用才会生效。

/*
// 错误的调用示例
func main() {
	defer func() {
		doRecover()
	}()
	panic("not good")
}

func doRecover() {
	fmt.Println("recobered: ", recover())
}
*/

// 调用示例
func main() {
	defer func() {
		fmt.Println(recover())
	}()
	a()
}

func a() {
	defer func() {
		deferRecover()
	}()
	b()
}

func b()  {
	c()
}

func c()  {
	panic("c error")
}

func deferRecover()  {
	fmt.Println(recover(), "deferRecover")
}
