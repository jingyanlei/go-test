package main

//11. string 类型的变量值不能为 nil
//对那些喜欢用 nil 初始化字符串的人来说，这就是坑：

/*
// 错误示例
func main() {
	var s string = nil    // cannot use nil as type string in assignment
	if s == nil {    // invalid operation: s == nil (mismatched types string and nil)
		s = "default"
	}
}
*/

// 正确示例
func main() {
	var s string    // 字符串类型的零值是空串 ""
	if s == "" {
		s = "default"
	}
}