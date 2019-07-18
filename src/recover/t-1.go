package main

import (
	"fmt"
	"time"
)

//panic和recover

//panic和recover的使用需要遵循以下原则：

//defer 需要放在 panic 之前定义，另外recover只有在 defer 调用的函数中才有效。
//recover处理异常后，逻辑并不会恢复到 panic 那个点去，函数跑到 defer 之后的那个点.
//多个 defer 会形成 defer 栈，后定义的 defer 语句会被最先调用
//使用recover捕获程序中的错误的用法如下


func main() {
	f()
	fmt.Println("end")
}

func f() {
	defer func() {//必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("defer start")
		if err := recover(); err != nil {
			fmt.Println(err) //这里的err其实就是panic传入的内容，"bug"
			}
		fmt.Println("defer end")
		}()
	for {
		fmt.Println("func begin")
		a := []string{"a", "b"}
		fmt.Println(a[3])    // 越界访问，肯定出现异常
		panic("bug") // 上面已经出现异常了,所以肯定走不到这里了。
		fmt.Println("func end") // 不会运行的.
		time.Sleep(1 * time.Second)
	}
}