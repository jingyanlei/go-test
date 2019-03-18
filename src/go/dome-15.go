package main

import "fmt"

//15. 访问 map 中不存在的 key

//Go 则会返回元素对应数据类型的零值，比如 nil、'' 、false 和 0，取值操作总有值返回，故不能通过取出来的值来判断 key 是不是在 map 中。

//检查 key 是否存在可以用 map 直接访问，检查返回的第二个参数即可：

/*
// 错误的 key 检测方式
func main() {
	x := map[string]string{"one": "2", "two": "", "three": "3"}
	if v := x["two"]; v == "" {
		fmt.Println("key two is no entry")    // 键 two 存不存在都会返回的空字符串
	}
}
*/

// 正确示例
func main() {
	x := map[string]string{"one": "2", "two": "", "three": "3"}
	if _, ok := x["two"]; !ok {
		fmt.Println("key two is no entry")
	}
}