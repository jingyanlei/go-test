package main

import "fmt"

//21. 在多行 array、slice、map 语句中缺少 , 号
//声明语句中 } 折叠到单行后，尾部的 , 不是必需的。
func main() {
	x := []int {
		1,
		2    // syntax error: unexpected newline, expecting comma or }
	}
	fmt.Println(x)
	y := []int{1,2,}
	fmt.Println(y)
	z := []int{1,2}
	fmt.Println(z)
	// ...
}