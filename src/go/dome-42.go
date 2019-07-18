package main

import (
	"bytes"
	"fmt"
)

//42. Slice 中数据的误用
//举个简单例子，重写文件路径（存储在 slice 中）
//
//分割路径来指向每个不同级的目录，修改第一个目录名再重组子目录名，创建新路径：



// 错误使用 slice 的拼接示例
func main() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/') // 4
	fmt.Println(string(path))
	println(sepIndex)

	dir1 := path[:sepIndex]
	dir2 := path[sepIndex+1:]
	println("dir1: ", string(dir1), len(dir1), cap(dir1))        // AAAA
	println("dir2: ", string(dir2), len(dir2), cap(dir2))        // BBBBBBBBB

	dir1 = append(dir1, "suffix"...)
	println("current path: ", string(path))    // AAAAsuffixBBBB

	path = bytes.Join([][]byte{dir1, dir2}, []byte{'/'})
	println("dir1: ", string(dir1))        // AAAAsuffix
	println("dir2: ", string(dir2))        // uffixBBBB

	println("new path: ", string(path))    // AAAAsuffix/uffixBBBB    // 错误结果
}


//拼接的结果不是正确的 AAAAsuffix/BBBBBBBBB，因为 dir1、 dir2 两个 slice 引用的数据都是 path 的底层数组，第 13 行修改 dir1 同时也修改了 path，也导致了 dir2 的修改
//
//解决方法：
//
//重新分配新的 slice 并拷贝你需要的数据
//使用完整的 slice 表达式：input[low:high:max]，容量便调整为 max - low
// 使用 full slice expression

/*
func main() {

	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/') // 4
	dir1 := path[:sepIndex:sepIndex]        // 此时 cap(dir1) 指定为4， 而不是先前的 16
	fmt.Println("dir1:", len(dir1), cap(dir1))
	dir2 := path[sepIndex+1:]
	fmt.Println("dir2:", len(dir2), cap(dir2))
	dir1 = append(dir1, "suffix"...)

	path = bytes.Join([][]byte{dir1, dir2}, []byte{'/'})
	println("dir1: ", string(dir1))        // AAAAsuffix
	println("dir2: ", string(dir2))        // BBBBBBBBB
	println("new path: ", string(path))    // AAAAsuffix/BBBBBBBBB
}
*/
//第 6 行中第三个参数是用来控制 dir1 的新容量，再往 dir1 中 append 超额元素时，将分配新的 buffer 来保存。而不是覆盖原来的 path 底层数组