package main

import (
	"fmt"
	"unicode/utf8"
)

//18. string 与索引操作符
//对字符串用索引访问返回的不是字符，而是一个 byte 值。
//
//这种处理方式和其他语言一样，比如 PHP 中：
//
//> php -r '$name="中文"; var_dump($name);'    # "中文" 占用 6 个字节
//string(6) "中文"
//
//> php -r '$name="中文"; var_dump($name[0]);' # 把第一个字节当做 Unicode 字符读取，显示 U+FFFD
//string(1) "�"
//
//> php -r '$name="中文"; var_dump($name[0].$name[1].$name[2]);'
//string(3) "中"

func main() {
	x := "景ascii"
	fmt.Println(x[0])        // 97
	fmt.Printf("%T\n", x[0])// uint8
	fmt.Printf("%T，%d, %d, %d\n", x[0], x[0], x[1], x[2])// uint8
	xRune := []rune(x)
	fmt.Println(xRune[0], string(xRune[0]))
	fmt.Println(utf8.ValidString(x))
}
//如果需要使用 for range 迭代访问字符串中的字符（unicode code point / rune），标准库中有 "unicode/utf8" 包来做 UTF8 的相关解码编码。另外 utf8string 也有像 func (s *String) At(i int) rune 等很方便的库函数。

