package main

import "fmt"

//???????????????
//14. slice 和 array 其实是一维数据
//
//
//看起来 Go 支持多维的 array 和 slice，可以创建数组的数组、切片的切片，但其实并不是。
//
//对依赖动态计算多维数组值的应用来说，就性能和复杂度而言，用 Go 实现的效果并不理想。
//
//可以使用原始的一维数组、“独立“ 的切片、“共享底层数组”的切片来创建动态的多维数组。
//
//使用原始的一维数组：要做好索引检查、溢出检测、以及当数组满时再添加值时要重新做内存分配。
//使用“独立”的切片分两步：
//创建外部 slice
//
//对每个内部 slice 进行内存分配
//注意内部的 slice 相互独立，使得任一内部 slice 增缩都不会影响到其他的 slice

// 使用各自独立的 6 个 slice 来创建 [2][3] 的动态多维数组
//func main() {
//	x := 2
//	y := 4
//
//	table := make([][]int, x)
//	for i  := range table {
//		table[i] = make([]int, y)
//	}
//}
//使用“共享底层数组”的切片
//创建一个存放原始数据的容器 slice
//创建其他的 slice
//切割原始 slice 来初始化其他的 slice
func main() {
	h, w := 2, 4
	raw := make([]int, h*w)

	for i := range raw {
		raw[i] = i
	}

	// 初始化原始 slice
	fmt.Println(raw, &raw[4])    // [0 1 2 3 4 5 6 7] 0xc420012120

	table := make([][]int, h)
	for i := range table {

		// 等间距切割原始 slice，创建动态多维数组 table
		// 0: raw[0*4: 0*4 + 4]
		// 1: raw[1*4: 1*4 + 4]
		table[i] = raw[i*w : i*w + w]
	}

	fmt.Println(table, &table[1][0])    // [[0 1 2 3] [4 5 6 7]] 0xc420012120
}

//更多关于多维数组的参考
//
//go-how-is-two-dimensional-arrays-memory-representation
//
//what-is-a-concise-way-to-create-a-2d-slice-in-go