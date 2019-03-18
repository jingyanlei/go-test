package main

import (
	"fmt"
	"unicode/utf8"
)

//20. 字符串的长度
//在 Python 中：
//
//data = u'♥'
//print(len(data)) # 1
//然而在 Go 中：

/*
func main() {
	char := "♥"
	fmt.Println(len(char))    // 3
}
*/

//Go 的内建函数 len() 返回的是字符串的 byte 数量，而不是像 Python 中那样是计算 Unicode 字符数。
//
//如果要得到字符串的字符数，可使用 "unicode/utf8" 包中的 RuneCountInString(str string) (n int)

/*
func main() {
	char := "♥"
	fmt.Println(utf8.RuneCountInString(char))    // 1
}
*/

//注意： RuneCountInString 并不总是返回我们看到的字符数，因为有的字符会占用 2 个 rune：

func main() {
	char := "é"
	fmt.Println(len(char))    // 3
	fmt.Println(utf8.RuneCountInString(char))    // 2
	fmt.Println("cafe\u0301")    // café    // 法文的 cafe，实际上是两个 rune 的组合
}