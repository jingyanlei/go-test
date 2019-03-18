package main

import "fmt"

//13. range 遍历 slice 和 array 时混淆了返回值
//与其他编程语言中的 for-in 、foreach 遍历语句不同，Go 中的 range 在遍历时会生成 2 个值，第一个是元素索引，第二个是元素的值：

/*
// 错误示例
func main() {
	x := []string{"a", "b", "c"}
	for v := range x {
		fmt.Println(v)    // 1 2 3
	}
}
*/


// 正确示例
func main() {
	x := []string{"a", "b", "c"}
	for _, v := range x {    // 使用 _ 丢弃索引
		fmt.Println(v)
	}
}
