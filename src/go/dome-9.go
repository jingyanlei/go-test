package main

import "fmt"

// 9. 直接使用值为 nil 的 slice、map
// 允许对值为 nil 的 slice 添加元素，但对值为 nil 的 map 添加元素则会造成运行时 panic

/*
// map 错误示例
func main() {
	var m map[string]int
	m["one"] = 1        // error: panic: assignment to entry in nil map
	// m := make(map[string]int)// map 的正确声明，分配了实际的内存
}
*/

// slice 正确示例
func main() {
	var s []int
	s = append(s, 1, 2, 3)
	fmt.Println(cap(s), len(s))
	fmt.Println(s)
	fmt.Println(cap(s), len(s))
	s[0] = 1
	fmt.Println(s)
}