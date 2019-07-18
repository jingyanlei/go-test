package main

import "fmt"

//switch 语句中的 case 代码块会默认带上 break，但可以使用 fallthrough 来强制执行下一个 case 代码块。

/*
func main() {
	isSpace := func(char byte) bool {
		switch char {
		case ' ':    // 空格符会直接 break，返回 false // 和其他语言不一样
			//fallthrough    // 返回 true
		case '\t':
			return true
		}
		return false
	}
	fmt.Println(isSpace('\t'))    // true
	fmt.Println(isSpace(' '))    // false
}
*/

//不过你可以在 case 代码块末尾使用 fallthrough，强制执行下一个 case 代码块。

//也可以改写 case 为多条件判断：

func main() {
	isSpace := func(char byte) bool {
		switch char {
		case ' ', '\t':
			return true
		}
		return false
	}
	fmt.Println(isSpace('\t'))    // true
	fmt.Println(isSpace(' '))    // true
	fmt.Println(isSpace('1'))    // true
}