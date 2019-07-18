package main

import "fmt"

//28. 按位取反
//很多编程语言使用 ~ 作为一元按位取反（NOT）操作符，Go 重用 ^ XOR 操作符来按位取反

/*
// 错误的取反操作
func main() {
	fmt.Println(~2)        // bitwise complement operator is ^
}
*/

/*
// 正确示例
func main() {
	var d uint8 = 2
	fmt.Printf("%08b\n", d)        // 00000010
	fmt.Printf("%08b\n", ^d)    // 11111101

	var e uint8 = 2
	fmt.Printf("%08b\n", e)        // 00000010
	fmt.Printf("%08b\n", ^e)    // 11111101
}
*/

//同时 ^ 也是按位异或（XOR）操作符。
//
//一个操作符能重用两次，是因为一元的 NOT 操作 NOT 0x02，与二元的 XOR 操作 0x22 XOR 0xff 是一致的。
//
//Go 也有特殊的操作符 AND NOT &^ 操作符，不同位才取1。

func main() {
	var a uint8 = 0x82
	var b uint8 = 0x02
	fmt.Printf("%08b [A]\n", a)
	fmt.Printf("%08b [B]\n", b)

	fmt.Printf("%08b (NOT B)\n", ^b)
	fmt.Printf("%08b ^ %08b = %08b [B XOR 0xff]\n", b, 0xff, b^0xff)

	fmt.Printf("%08b ^ %08b = %08b [A XOR B]\n", a, b, a^b)
	fmt.Printf("%08b & %08b = %08b [A AND B]\n", a, b, a&b)
	fmt.Printf("%08b &^%08b = %08b [A 'AND NOT' B]\n", a, b, a&^b)
	fmt.Printf("%08b&(^%08b)= %08b [A AND (NOT B)]\n", a, b, a&(^b))
}