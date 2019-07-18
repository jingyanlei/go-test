package main

import "fmt"

//12. Array 类型的值作为函数参数
//在 C/C++ 中，数组（名）是指针。将数组作为参数传进函数时，相当于传递了数组内存地址的引用，在函数内部会改变该数组的值。
//
//在 Go 中，数组是值。作为参数传进函数时，传递的是数组的原始值拷贝，此时在函数内部是无法更新该数组的：

/*
// 数组使用值拷贝传参，函数外部数组不会更新
func main() {
	x := [3]int{1,2,3}

	func(arr [3]int) {
		arr[0] = 7
		fmt.Println(arr)    // [7 2 3]
	}(x)
	fmt.Println(x)            // [1 2 3]    // 并不是你以为的 [7 2 3]
}
*/

//如果想修改参数数组：
//
//直接传递指向这个数组的指针类型：

/*
// 传址会修改原数据
func main() {
	x := [3]int{1,2,3}

	func(arr *[3]int) {
		(*arr)[0] = 7
		fmt.Println(arr)    // &[7 2 3]
	}(&x)
	fmt.Println(x)    // [7 2 3]
}
*/

//直接使用 slice：即使函数内部得到的是 slice 的值拷贝，但依旧会更新 slice 的原始数据（底层 array）
// 会修改 slice 的底层 array，从而修改 slice
// function append 追加内容后，会生成新的 slice 地址，函数外的 slice 不会增加
func main() {
	x := []int{1, 2, 3}
	func(arr []int) {
		arr[0] = 7
		fmt.Printf("%p, %v\n", arr, arr)
		arr = append(arr, 4)
		fmt.Printf("%p, %v\n", arr, arr)
		fmt.Printf("%p, %v\n", x, x)
	}(x)
	fmt.Printf("%p, %v\n", x, x)
}