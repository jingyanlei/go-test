package main

import "fmt"

func main()  {

	// slice 切片 引用类型 对值的修改 a b 会同时修改
	a := []int{1,2,3,4,5}
	b := a
	b[2] = 22
	testSlice(b)
	fmt.Println(a, b)
	fmt.Println(&a, &b)

	// array 数组 值类型 对值的赋值是copy 对 bb 的修改不会影响aa
	aa := [5]int{1,2,3,4,5}
	bb := aa
	bb[2] = 22
	testArray(bb)
	fmt.Println(aa, bb)
	fmt.Println(&aa, &bb)

	// map 切片 引用类型 对值的修改 aaa bbb 会同时修改
	aaa := map[string]string{}
	aaa["a"] = "a"
	aaa["b"] = "b"
	aaa["c"] = "c"
	bbb := aaa
	bbb["b"] = "bb"
	testMap(bbb)
	fmt.Println(aaa, bbb)
	fmt.Println(&aaa, &bbb)

	// string 字符串 值类型 对值的赋值是copy 对 str2 的修改不会影响str
	str := "string"
	str2 := str
	str2 = "string222"
	testString(str2)
	fmt.Println(str, str2)
	fmt.Println(&str, &str2)
	
}

// map
func testMap(aaa map[string]string) {
	aaa["e"] = "e"
	fmt.Println(&aaa)
}

// slice
func testSlice(a []int) {
	a[4] = 444
}

func testString(str2 string)  {
	str2 += "ssssssssss"
}

// array
func testArray(aa [5]int)  {
	aa[4] = 444
}