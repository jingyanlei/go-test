package main

import "fmt"

//27. 自增和自减运算
//很多编程语言都自带前置后置的 ++、-- 运算。但 Go 特立独行，去掉了前置操作，同时 ++、— 只作为运算符而非表达式。

/*
// 错误示例
func main() {
	data := []int{1, 2, 3}
	i := 0
	++i            // syntax error: unexpected ++, expecting }
	fmt.Println(data[i++])    // syntax error: unexpected ++, expecting :
}
*/


//正确示例
func main() {
	data := []int{1, 2, 3}
	i := 0
	i++
	fmt.Println(data[i])    // 2
}