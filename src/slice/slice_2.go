package main

import "fmt"

// slice
func main()  {
	var a = make([]int, 10)
	a[0] = 0
	a[1] = 1
	a[2] = 2
	a[3] = 3
	a[4] = 4
	fmt.Printf("a init：%+v\n", a)
	t(a)
	fmt.Printf("a appe：%+v\n", a)
	fmt.Println("---")

	fmt.Println("a copy to b")
	var b = make([]int, 10)
	copy(b, a[:])
	fmt.Printf("b init：%+v\n", b)
	t2(&b)
	fmt.Printf("b appe：%+v\n", b)

}

// slice 值传递append添加，函数外不会更新
func t(s []int)  {
	s[5] = 5
	s[6] = 6
	b := []int{6,7,8,9}
	s = append(s, b...)
	fmt.Printf("a func：%+v\n", s)

}

// slice 指针传递append添加的值，函数外会更新
// 值指针不能使用indexing (invalid operation: s[0] (type *[]int does not support indexing))
func t2(s *[]int) {
	//*s[0] = 1
	b := []int{6,7,8,9}
	*s = append(*s, b...)
	fmt.Printf("b func:%+v\n", s)

}


